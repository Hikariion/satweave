package sun

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"path"
	"satweave/messenger"
	"satweave/messenger/common"
	task_manager "satweave/shared/task-manager"
	common2 "satweave/utils/common"
	"satweave/utils/logger"
)

type StreamHelper struct {
	taskRegisteredTaskManagerTable *RegisteredTaskManagerTable
	Scheduler                      *DefaultScheduler
	snapshotDir                    string
	streamJobDataDir               string
}

type TaskTuple struct {
	SatelliteName string
	ExecuteTask   *common.ExecuteTask
}

func (s *StreamHelper) DeployExecuteTasks(ctx context.Context, jobId string, executeMap map[string][]*common.ExecuteTask,
	pathNodes []string, yamlBytes []byte) error {
	for taskManagerId, executeTasks := range executeMap {
		HostPort := s.taskRegisteredTaskManagerTable.table[taskManagerId].HostPort
		host := HostPort.GetHost()
		port := HostPort.GetPort()
		conn, err := messenger.GetRpcConn(host, port)
		if err != nil {
			logger.Errorf("get rpc conn failed: %v", err)
			return err
		}
		client := task_manager.NewTaskManagerServiceClient(conn)
		// 每个 Execute task 都需要 deploy
		// deploy的过程其实是创建一个worker的过程
		for _, executeTask := range executeTasks {
			_, err := client.DeployTask(ctx, &task_manager.DeployTaskRequest{
				ExecTask:  executeTask,
				JobId:     jobId,
				PathNodes: pathNodes,
				YamlBytes: yamlBytes,
			})
			if err != nil {
				logger.Errorf("deploy task on task manager id: %v, failed: %v", taskManagerId, err)
				return err
			}
		}
		logger.Infof("Deploy all execute task on task manager id: %v, success", taskManagerId)
	}

	return nil
}

func (s *StreamHelper) StartExecuteTasks(jobId string, logicalMap map[string][]*common.Task, executeMap map[string][]*common.ExecuteTask) error {
	// clsName -> Task
	taskNameInvertedIndex := make(map[string]*common.Task)
	for _, tasks := range logicalMap {
		for _, task := range tasks {
			if _, ok := taskNameInvertedIndex[task.ClsName]; !ok {
				taskNameInvertedIndex[task.ClsName] = task
			}
		}
	}

	// clsName -> output task: List[str]
	nextLogicalTasks := make(map[string][]string)
	for _, tasks := range logicalMap {
		for _, task := range tasks {
			for _, preTask := range task.InputTasks {
				if _, ok := nextLogicalTasks[preTask]; !ok {
					nextLogicalTasks[preTask] = make([]string, 0)
				}
				nextLogicalTasks[preTask] = append(nextLogicalTasks[preTask], task.ClsName)
			}
		}
	}

	// subtaskName -> (taskManagerId, ExecuteTask)
	subTaskNameInvertedIndex := make(map[string]*TaskTuple)
	for satelliteName, executeTasks := range executeMap {
		for _, executeTask := range executeTasks {
			subTaskNameInvertedIndex[executeTask.SubtaskName] = &TaskTuple{
				SatelliteName: satelliteName,
				ExecuteTask:   executeTask,
			}
		}
	}

	startedTasks := make(map[string]bool)
	for clsName, _ := range taskNameInvertedIndex {
		if _, ok := startedTasks[clsName]; !ok {
			s.dfsToStartExecuteTask(jobId, clsName, nextLogicalTasks, taskNameInvertedIndex, subTaskNameInvertedIndex, startedTasks)
		}
	}

	return nil
}

func (s *StreamHelper) dfsToStartExecuteTask(jobId string, clsName string, nextLogicalTasks map[string][]string, logicalTaskNameInvertedIndex map[string]*common.Task,
	subtaskNameInvertedIndex map[string]*TaskTuple, startedTasks map[string]bool) {
	if _, ok := startedTasks[clsName]; ok {
		return
	}
	startedTasks[clsName] = true
	list, ok := nextLogicalTasks[clsName]
	if ok {
		for _, nextLogicalTaskName := range list {
			if _, exists := startedTasks[nextLogicalTaskName]; !exists {
				s.dfsToStartExecuteTask(jobId, nextLogicalTaskName, nextLogicalTasks, logicalTaskNameInvertedIndex, subtaskNameInvertedIndex, startedTasks)
			}
		}
	}

	logicalTask := logicalTaskNameInvertedIndex[clsName]
	for i := 0; i < int(logicalTask.Currency); i++ {
		subtaskName := s.getSubTaskName(jobId, clsName, i, int(logicalTask.Currency))
		SatelliteName := subtaskNameInvertedIndex[subtaskName].SatelliteName
		executeTask := subtaskNameInvertedIndex[subtaskName].ExecuteTask
		s.innerDfsToStartExecuteTask(SatelliteName, executeTask)
	}

}

func (s *StreamHelper) innerDfsToStartExecuteTask(satelliteName string, executeTask *common.ExecuteTask) {
	HostPort := s.taskRegisteredTaskManagerTable.table[satelliteName].HostPort
	host := HostPort.GetHost()
	port := HostPort.GetPort()
	conn, err := messenger.GetRpcConn(host, port)
	if err != nil {
		logger.Errorf("Fail to get rpc conn on TaskManager %v", satelliteName)
	}
	client := task_manager.NewTaskManagerServiceClient(conn)
	_, err = client.StartTask(context.Background(), &task_manager.StartTaskRequest{
		SubtaskName: executeTask.SubtaskName,
	})
	if err != nil {
		logger.Errorf("Fail to start subtask: %v on task manager id: ", executeTask.SubtaskName, satelliteName)
	}
	return
}

func (s *StreamHelper) RegisterTaskManager(_ context.Context, request *RegisterTaskManagerRequest) (*RegisterTaskManagerResponse, error) {
	err := s.taskRegisteredTaskManagerTable.register(request.TaskManagerDesc)
	if err != nil {
		return &RegisterTaskManagerResponse{
			Success: false,
		}, status.Errorf(codes.Internal, "register task manager failed: %v", err)
	}
	logger.Infof("Register task manager success: %v", request.TaskManagerDesc)
	return &RegisterTaskManagerResponse{
		Success: true,
	}, nil
}

//func (s *StreamHelper) GetRegisterTaskManagerTable(context.Context, *common.NilRequest) (*TaskManagerResult, error) {
//	return &TaskManagerResult{
//		TaskManagerTable: s.taskRegisteredTaskManagerTable.table,
//	}, nil
//}

func (s *StreamHelper) PrintTaskManagerTable() {
	logger.Infof("Sun printing task manager table...")
	logger.Infof("%v", s.taskRegisteredTaskManagerTable.table)
	logger.Infof("%v", s.Scheduler.RegisteredTaskManagerTable)
}

func (s *StreamHelper) getSubTaskName(jobId string, clsName string, idx, currency int) string {
	return fmt.Sprintf("%s#%s#-%d-%d", jobId, clsName, idx+1, currency)
}

func (s *StreamHelper) RestoreFromCheckpoint(SubtaskName string) ([]byte, error) {
	fullPath := path.Join(s.snapshotDir, SubtaskName)

	// 使用os.Stat检查文件是否存在
	if _, err := os.Stat(fullPath); err != nil {
		if os.IsNotExist(err) {
			// 文件不存在，返回nil
			return nil, nil
		}
		// 其他错误，这里简单处理为返回nil
		// 实际应用中你可能需要更复杂的错误处理
		return nil, err
	}

	data, err := os.ReadFile(fullPath)
	if err != nil {
		// 读取错误，返回nil
		// 同样，你可能需要根据实际情况处理这类错误
		return nil, err
	}

	return data, err
}

func (s *StreamHelper) SaveSnapShot(SubtaskName string, state []byte) error {
	file, err := os.Create(path.Join(s.snapshotDir, SubtaskName))
	defer file.Close()
	if err != nil {
		return err
	}

	_, err = file.Write(state)

	file.Sync()

	return nil
}

func (s *StreamHelper) SaveStreamJobData(jobId, dataId, data string) error {
	jobDataFilePath := path.Join(s.streamJobDataDir, jobId)

	file, err := os.OpenFile(jobDataFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		// 如果打开文件失败，则返回错误
		return err
	}
	defer file.Close()

	// 构造要追加的数据，格式为 "dataId: data\n"
	formattedData := fmt.Sprintf("%s: %s\n", dataId, data)

	// 向文件追加数据
	if _, err := file.WriteString(formattedData); err != nil {
		return err // 如果写入文件失败，则返回错误
	}

	return nil
}

func NewStreamHelper() *StreamHelper {
	streamHelper := &StreamHelper{
		taskRegisteredTaskManagerTable: newRegisteredTaskManagerTable(),
		snapshotDir:                    "./snapshot",
		streamJobDataDir:               "./stream-job-data",
	}

	// 创建 snapshot 目录
	err := common2.InitPath(streamHelper.snapshotDir)
	if err != nil {
		logger.Errorf("create snapshot dir failed: %v", err)
		return nil
	}

	err = common2.InitPath(streamHelper.streamJobDataDir)
	if err != nil {
		logger.Errorf("create stream job data dir failed: %v", err)
		return nil
	}

	streamHelper.Scheduler = newUserDefinedScheduler(streamHelper.taskRegisteredTaskManagerTable)
	return streamHelper
}
