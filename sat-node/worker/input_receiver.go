package worker

import (
	"satweave/messenger/common"
	"satweave/utils/logger"
	"sync"
	"time"
)

type InputReceiver struct {
	/*
	   PartitionDispenser   \
	   PartitionDispenser   -    subtask   ->   InputReceiver   ->   channel
	   PartitionDispenser   /            (遇到 event 阻塞，类似 Gate)
	*/
	inputChannel chan *common.Record
	partitions   []*InputPartitionReceiver
	// TODO(qiu): barrier 需要 new 并且初始化，大小 = partition 数量
	barrier *sync.WaitGroup
}

func (i *InputReceiver) RecvData(partitionIdx uint64, record *common.Record) {
	i.partitions[partitionIdx].RecvData(record)
}

func (i *InputReceiver) RunAllPartitionReceiver() {
	for _, partition := range i.partitions {
		partition.Run()
	}
}

type InputPartitionReceiver struct {
	queue   chan *common.Record
	channel chan *common.Record
	barrier *sync.WaitGroup
}

func (ipr *InputPartitionReceiver) RecvData(record *common.Record) {
	ipr.queue <- record
}

func (ipr *InputPartitionReceiver) praseDataAndCarryToChannel(inputQueue chan *common.Record, outputChannel chan *common.Record, barrier *sync.WaitGroup, allowOne chan struct{}) error {
	var needBarrierDataType = common.DataType_CHECKPOINT
	for {
		data := <-inputQueue
		if data.DataType == needBarrierDataType {
			// TODO(qiu): 需要在某个地方初始化 wg
			barrier.Done()
			logger.Infof("Receive Barrier wg --, begin to wait ... ")
			barrier.Wait()
			logger.Infof("Block finished, begin to make checkpoint")
			// 重置 wg，每个协程都只执行一次
			barrier.Add(1)
			// allowOne 用于只让一个协程往 output 发送数据
			select {
			case <-allowOne:
				logger.Infof("Be the allowOne, transfer checkpoint signal")
				outputChannel <- data
				// 休眠 1s
				time.Sleep(1 * time.Second)
				allowOne <- struct{}{}
			default:
				logger.Infof("Do not be the allowOne, do not transfer checkpoint signal")
			}
		} else {
			outputChannel <- data
		}
	}
}

func (ipr *InputPartitionReceiver) Run() {
	go func() {
		err := ipr.praseDataAndCarryToChannel(ipr.queue, ipr.channel, ipr.barrier, make(chan struct{}, 1))
		if err != nil {
			logger.Fatalf("InputPartitionReceiver.praseDataAndCarryToChannel() failed: %v", err)
		}
	}()
}

func NewInputReceiver(inputChannel chan *common.Record, inputEndpoints []*common.InputEndpoints) *InputReceiver {
	inputReceiver := &InputReceiver{
		inputChannel: inputChannel,
		barrier:      &sync.WaitGroup{},
		partitions:   make([]*InputPartitionReceiver, 0),
	}
	for i := 0; i < len(inputEndpoints); i++ {
		inputPartitionReceiver := NewInputPartitionReceiver(inputChannel, inputReceiver.barrier)
		inputReceiver.partitions = append(inputReceiver.partitions, inputPartitionReceiver)
	}
	return inputReceiver
}

func NewInputPartitionReceiver(channel chan *common.Record, eventBarrier *sync.WaitGroup) *InputPartitionReceiver {
	return &InputPartitionReceiver{
		channel: channel,
		queue:   make(chan *common.Record, 1000),
		barrier: eventBarrier,
	}
}
