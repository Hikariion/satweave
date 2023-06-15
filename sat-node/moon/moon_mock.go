package moon

import (
	"context"
	"github.com/golang/mock/gomock"
	"go.etcd.io/etcd/raft/v3/raftpb"
	"reflect"
	"satweave/sat-node/infos"
	"satweave/shared/moon"
)

// MockInfoController is a mock of InfoController interface.
type MockInfoController struct {
	ctrl     *gomock.Controller
	recorder *MockInfoControllerMockRecorder
	moon.UnimplementedMoonServer
}

// MockInfoControllerMockRecorder is the mock recorder for MockInfoController.
type MockInfoControllerMockRecorder struct {
	mock *MockInfoController
}

// NewMockInfoController creates a new mock instance.
func NewMockInfoController(ctrl *gomock.Controller) *MockInfoController {
	mock := &MockInfoController{ctrl: ctrl}
	mock.recorder = &MockInfoControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInfoController) EXPECT() *MockInfoControllerMockRecorder {
	return m.recorder
}

// GetInfo mocks base method.
func (m *MockInfoController) GetInfo(arg0 context.Context, arg1 *moon.GetInfoRequest) (*moon.GetInfoReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo", arg0, arg1)
	ret0, _ := ret[0].(*moon.GetInfoReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfo indicates an expected call of GetInfo.
func (mr *MockInfoControllerMockRecorder) GetInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockInfoController)(nil).GetInfo), arg0, arg1)
}

// GetInfoDirect mocks base method.
func (m *MockInfoController) GetInfoDirect(infoType infos.InfoType, id string) (infos.Information, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfoDirect", infoType, id)
	ret0, _ := ret[0].(infos.Information)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfoDirect indicates an expected call of GetInfoDirect.
func (mr *MockInfoControllerMockRecorder) GetInfoDirect(infoType, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfoDirect", reflect.TypeOf((*MockInfoController)(nil).GetInfoDirect), infoType, id)
}

// GetLeaderID mocks base method.
func (m *MockInfoController) GetLeaderID() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLeaderID")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetLeaderID indicates an expected call of GetLeaderID.
func (mr *MockInfoControllerMockRecorder) GetLeaderID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLeaderID", reflect.TypeOf((*MockInfoController)(nil).GetLeaderID))
}

// IsLeader mocks base method.
func (m *MockInfoController) IsLeader() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLeader")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsLeader indicates an expected call of IsLeader.
func (mr *MockInfoControllerMockRecorder) IsLeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLeader", reflect.TypeOf((*MockInfoController)(nil).IsLeader))
}

// ListInfo mocks base method.
func (m *MockInfoController) ListInfo(arg0 context.Context, arg1 *moon.ListInfoRequest) (*moon.ListInfoReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInfo", arg0, arg1)
	ret0, _ := ret[0].(*moon.ListInfoReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInfo indicates an expected call of ListInfo.
func (mr *MockInfoControllerMockRecorder) ListInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInfo", reflect.TypeOf((*MockInfoController)(nil).ListInfo), arg0, arg1)
}

// NodeInfoChanged mocks base method.
func (m *MockInfoController) NodeInfoChanged(nodeInfo *infos.NodeInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NodeInfoChanged", nodeInfo)
}

// NodeInfoChanged indicates an expected call of NodeInfoChanged.
func (mr *MockInfoControllerMockRecorder) NodeInfoChanged(nodeInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeInfoChanged", reflect.TypeOf((*MockInfoController)(nil).NodeInfoChanged), nodeInfo)
}

// ProposeConfChangeAddNode mocks base method.
func (m *MockInfoController) ProposeConfChangeAddNode(ctx context.Context, nodeInfo *infos.NodeInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProposeConfChangeAddNode", ctx, nodeInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProposeConfChangeAddNode indicates an expected call of ProposeConfChangeAddNode.
func (mr *MockInfoControllerMockRecorder) ProposeConfChangeAddNode(ctx, nodeInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProposeConfChangeAddNode", reflect.TypeOf((*MockInfoController)(nil).ProposeConfChangeAddNode), ctx, nodeInfo)
}

// ProposeInfo mocks base method.
func (m *MockInfoController) ProposeInfo(arg0 context.Context, arg1 *moon.ProposeInfoRequest) (*moon.ProposeInfoReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProposeInfo", arg0, arg1)
	ret0, _ := ret[0].(*moon.ProposeInfoReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProposeInfo indicates an expected call of ProposeInfo.
func (mr *MockInfoControllerMockRecorder) ProposeInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProposeInfo", reflect.TypeOf((*MockInfoController)(nil).ProposeInfo), arg0, arg1)
}

// Run mocks base method.
func (m *MockInfoController) Run() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run")
}

// Run indicates an expected call of Run.
func (mr *MockInfoControllerMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockInfoController)(nil).Run))
}

// SendRaftMessage mocks base method.
func (m *MockInfoController) SendRaftMessage(arg0 context.Context, arg1 *raftpb.Message) (*raftpb.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendRaftMessage", arg0, arg1)
	ret0, _ := ret[0].(*raftpb.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendRaftMessage indicates an expected call of SendRaftMessage.
func (mr *MockInfoControllerMockRecorder) SendRaftMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRaftMessage", reflect.TypeOf((*MockInfoController)(nil).SendRaftMessage), arg0, arg1)
}

// Set mocks base method.
func (m *MockInfoController) Set(selfInfo, leaderInfo *infos.NodeInfo, peersInfo []*infos.NodeInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", selfInfo, leaderInfo, peersInfo)
}

// Set indicates an expected call of Set.
func (mr *MockInfoControllerMockRecorder) Set(selfInfo, leaderInfo, peersInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockInfoController)(nil).Set), selfInfo, leaderInfo, peersInfo)
}

// Stop mocks base method.
func (m *MockInfoController) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockInfoControllerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockInfoController)(nil).Stop))
}

// mustEmbedUnimplementedMoonServer mocks base method.
func (m *MockInfoController) mustEmbedUnimplementedMoonServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedMoonServer")
}

// mustEmbedUnimplementedMoonServer indicates an expected call of mustEmbedUnimplementedMoonServer.
func (mr *MockInfoControllerMockRecorder) mustEmbedUnimplementedMoonServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedMoonServer", reflect.TypeOf((*MockInfoController)(nil).mustEmbedUnimplementedMoonServer))
}
