// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: monitor.proto

package watcher

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	common "satweave/messenger/common"
	infos "satweave/sat-node/infos"
	timestamp "satweave/utils/timestamp"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NodeStatusReport_Role int32

const (
	NodeStatusReport_UNKNOWN  NodeStatusReport_Role = 0
	NodeStatusReport_LEADER   NodeStatusReport_Role = 1
	NodeStatusReport_FOLLOWER NodeStatusReport_Role = 2
	NodeStatusReport_LEARNER  NodeStatusReport_Role = 3
)

// Enum value maps for NodeStatusReport_Role.
var (
	NodeStatusReport_Role_name = map[int32]string{
		0: "UNKNOWN",
		1: "LEADER",
		2: "FOLLOWER",
		3: "LEARNER",
	}
	NodeStatusReport_Role_value = map[string]int32{
		"UNKNOWN":  0,
		"LEADER":   1,
		"FOLLOWER": 2,
		"LEARNER":  3,
	}
)

func (x NodeStatusReport_Role) Enum() *NodeStatusReport_Role {
	p := new(NodeStatusReport_Role)
	*p = x
	return p
}

func (x NodeStatusReport_Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeStatusReport_Role) Descriptor() protoreflect.EnumDescriptor {
	return file_monitor_proto_enumTypes[0].Descriptor()
}

func (NodeStatusReport_Role) Type() protoreflect.EnumType {
	return &file_monitor_proto_enumTypes[0]
}

func (x NodeStatusReport_Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeStatusReport_Role.Descriptor instead.
func (NodeStatusReport_Role) EnumDescriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{0, 0}
}

type PipelineReport_State int32

const (
	PipelineReport_UNKNOWN    PipelineReport_State = 0
	PipelineReport_OK         PipelineReport_State = 1
	PipelineReport_DOWN_GRADE PipelineReport_State = 2
	PipelineReport_CHANGING   PipelineReport_State = 3
	PipelineReport_ERROR      PipelineReport_State = 4
)

// Enum value maps for PipelineReport_State.
var (
	PipelineReport_State_name = map[int32]string{
		0: "UNKNOWN",
		1: "OK",
		2: "DOWN_GRADE",
		3: "CHANGING",
		4: "ERROR",
	}
	PipelineReport_State_value = map[string]int32{
		"UNKNOWN":    0,
		"OK":         1,
		"DOWN_GRADE": 2,
		"CHANGING":   3,
		"ERROR":      4,
	}
)

func (x PipelineReport_State) Enum() *PipelineReport_State {
	p := new(PipelineReport_State)
	*p = x
	return p
}

func (x PipelineReport_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PipelineReport_State) Descriptor() protoreflect.EnumDescriptor {
	return file_monitor_proto_enumTypes[1].Descriptor()
}

func (PipelineReport_State) Type() protoreflect.EnumType {
	return &file_monitor_proto_enumTypes[1]
}

func (x PipelineReport_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PipelineReport_State.Descriptor instead.
func (PipelineReport_State) EnumDescriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{2, 0}
}

type ClusterReport_ClusterState int32

const (
	ClusterReport_UNKNOWN     ClusterReport_ClusterState = 0
	ClusterReport_HEALTH_OK   ClusterReport_ClusterState = 1
	ClusterReport_HEALTH_WARN ClusterReport_ClusterState = 2
	ClusterReport_HEALTH_ERR  ClusterReport_ClusterState = 3
)

// Enum value maps for ClusterReport_ClusterState.
var (
	ClusterReport_ClusterState_name = map[int32]string{
		0: "UNKNOWN",
		1: "HEALTH_OK",
		2: "HEALTH_WARN",
		3: "HEALTH_ERR",
	}
	ClusterReport_ClusterState_value = map[string]int32{
		"UNKNOWN":     0,
		"HEALTH_OK":   1,
		"HEALTH_WARN": 2,
		"HEALTH_ERR":  3,
	}
)

func (x ClusterReport_ClusterState) Enum() *ClusterReport_ClusterState {
	p := new(ClusterReport_ClusterState)
	*p = x
	return p
}

func (x ClusterReport_ClusterState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClusterReport_ClusterState) Descriptor() protoreflect.EnumDescriptor {
	return file_monitor_proto_enumTypes[2].Descriptor()
}

func (ClusterReport_ClusterState) Type() protoreflect.EnumType {
	return &file_monitor_proto_enumTypes[2]
}

func (x ClusterReport_ClusterState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClusterReport_ClusterState.Descriptor instead.
func (ClusterReport_ClusterState) EnumDescriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{3, 0}
}

type NodeStatusReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId    uint64                `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	NodeUuid  string                `protobuf:"bytes,2,opt,name=node_uuid,json=nodeUuid,proto3" json:"node_uuid,omitempty"`
	Timestamp *timestamp.Timestamp  `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	State     infos.NodeState       `protobuf:"varint,4,opt,name=state,proto3,enum=messenger.NodeState" json:"state,omitempty"`
	Status    *NodeStatus           `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Role      NodeStatusReport_Role `protobuf:"varint,6,opt,name=role,proto3,enum=messenger.NodeStatusReport_Role" json:"role,omitempty"`
	Pipelines []*PipelineReport     `protobuf:"bytes,7,rep,name=pipelines,proto3" json:"pipelines,omitempty"`
}

func (x *NodeStatusReport) Reset() {
	*x = NodeStatusReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStatusReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStatusReport) ProtoMessage() {}

func (x *NodeStatusReport) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStatusReport.ProtoReflect.Descriptor instead.
func (*NodeStatusReport) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{0}
}

func (x *NodeStatusReport) GetNodeId() uint64 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *NodeStatusReport) GetNodeUuid() string {
	if x != nil {
		return x.NodeUuid
	}
	return ""
}

func (x *NodeStatusReport) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *NodeStatusReport) GetState() infos.NodeState {
	if x != nil {
		return x.State
	}
	return infos.NodeState(0)
}

func (x *NodeStatusReport) GetStatus() *NodeStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *NodeStatusReport) GetRole() NodeStatusReport_Role {
	if x != nil {
		return x.Role
	}
	return NodeStatusReport_UNKNOWN
}

func (x *NodeStatusReport) GetPipelines() []*PipelineReport {
	if x != nil {
		return x.Pipelines
	}
	return nil
}

type NodeStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiskAvailable     uint64  `protobuf:"varint,1,opt,name=disk_available,json=diskAvailable,proto3" json:"disk_available,omitempty"`
	DiskTotal         uint64  `protobuf:"varint,2,opt,name=disk_total,json=diskTotal,proto3" json:"disk_total,omitempty"`
	NetworkSpeedSend  uint64  `protobuf:"varint,3,opt,name=network_speed_send,json=networkSpeedSend,proto3" json:"network_speed_send,omitempty"`
	NetworkSpeedRecv  uint64  `protobuf:"varint,4,opt,name=network_speed_recv,json=networkSpeedRecv,proto3" json:"network_speed_recv,omitempty"`
	CpuPercent        float64 `protobuf:"fixed64,5,opt,name=cpu_percent,json=cpuPercent,proto3" json:"cpu_percent,omitempty"`
	MemoryUsage       uint64  `protobuf:"varint,6,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`
	MemoryTotal       uint64  `protobuf:"varint,7,opt,name=memory_total,json=memoryTotal,proto3" json:"memory_total,omitempty"`
	BlockCount        uint64  `protobuf:"varint,8,opt,name=block_count,json=blockCount,proto3" json:"block_count,omitempty"`
	MetaCount         uint64  `protobuf:"varint,9,opt,name=meta_count,json=metaCount,proto3" json:"meta_count,omitempty"`
	GoroutineCount    uint64  `protobuf:"varint,10,opt,name=goroutine_count,json=goroutineCount,proto3" json:"goroutine_count,omitempty"`
	MetaPipelineCount uint64  `protobuf:"varint,11,opt,name=meta_pipeline_count,json=metaPipelineCount,proto3" json:"meta_pipeline_count,omitempty"`
}

func (x *NodeStatus) Reset() {
	*x = NodeStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStatus) ProtoMessage() {}

func (x *NodeStatus) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStatus.ProtoReflect.Descriptor instead.
func (*NodeStatus) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{1}
}

func (x *NodeStatus) GetDiskAvailable() uint64 {
	if x != nil {
		return x.DiskAvailable
	}
	return 0
}

func (x *NodeStatus) GetDiskTotal() uint64 {
	if x != nil {
		return x.DiskTotal
	}
	return 0
}

func (x *NodeStatus) GetNetworkSpeedSend() uint64 {
	if x != nil {
		return x.NetworkSpeedSend
	}
	return 0
}

func (x *NodeStatus) GetNetworkSpeedRecv() uint64 {
	if x != nil {
		return x.NetworkSpeedRecv
	}
	return 0
}

func (x *NodeStatus) GetCpuPercent() float64 {
	if x != nil {
		return x.CpuPercent
	}
	return 0
}

func (x *NodeStatus) GetMemoryUsage() uint64 {
	if x != nil {
		return x.MemoryUsage
	}
	return 0
}

func (x *NodeStatus) GetMemoryTotal() uint64 {
	if x != nil {
		return x.MemoryTotal
	}
	return 0
}

func (x *NodeStatus) GetBlockCount() uint64 {
	if x != nil {
		return x.BlockCount
	}
	return 0
}

func (x *NodeStatus) GetMetaCount() uint64 {
	if x != nil {
		return x.MetaCount
	}
	return 0
}

func (x *NodeStatus) GetGoroutineCount() uint64 {
	if x != nil {
		return x.GoroutineCount
	}
	return 0
}

func (x *NodeStatus) GetMetaPipelineCount() uint64 {
	if x != nil {
		return x.MetaPipelineCount
	}
	return 0
}

type PipelineReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PgId     uint64               `protobuf:"varint,1,opt,name=pg_id,json=pgId,proto3" json:"pg_id,omitempty"`
	NodeIds  []uint64             `protobuf:"varint,2,rep,packed,name=node_ids,json=nodeIds,proto3" json:"node_ids,omitempty"`
	LeaderId uint64               `protobuf:"varint,3,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	State    PipelineReport_State `protobuf:"varint,4,opt,name=state,proto3,enum=messenger.PipelineReport_State" json:"state,omitempty"`
}

func (x *PipelineReport) Reset() {
	*x = PipelineReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineReport) ProtoMessage() {}

func (x *PipelineReport) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineReport.ProtoReflect.Descriptor instead.
func (*PipelineReport) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{2}
}

func (x *PipelineReport) GetPgId() uint64 {
	if x != nil {
		return x.PgId
	}
	return 0
}

func (x *PipelineReport) GetNodeIds() []uint64 {
	if x != nil {
		return x.NodeIds
	}
	return nil
}

func (x *PipelineReport) GetLeaderId() uint64 {
	if x != nil {
		return x.LeaderId
	}
	return 0
}

func (x *PipelineReport) GetState() PipelineReport_State {
	if x != nil {
		return x.State
	}
	return PipelineReport_UNKNOWN
}

type ClusterReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State       ClusterReport_ClusterState `protobuf:"varint,1,opt,name=state,proto3,enum=messenger.ClusterReport_ClusterState" json:"state,omitempty"`
	ClusterInfo *infos.ClusterInfo         `protobuf:"bytes,2,opt,name=cluster_info,json=clusterInfo,proto3" json:"cluster_info,omitempty"`
	Nodes       []*NodeStatusReport        `protobuf:"bytes,3,rep,name=nodes,proto3" json:"nodes,omitempty"`
	Pipelines   []*PipelineReport          `protobuf:"bytes,4,rep,name=pipelines,proto3" json:"pipelines,omitempty"`
}

func (x *ClusterReport) Reset() {
	*x = ClusterReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterReport) ProtoMessage() {}

func (x *ClusterReport) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterReport.ProtoReflect.Descriptor instead.
func (*ClusterReport) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{3}
}

func (x *ClusterReport) GetState() ClusterReport_ClusterState {
	if x != nil {
		return x.State
	}
	return ClusterReport_UNKNOWN
}

func (x *ClusterReport) GetClusterInfo() *infos.ClusterInfo {
	if x != nil {
		return x.ClusterInfo
	}
	return nil
}

func (x *ClusterReport) GetNodes() []*NodeStatusReport {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *ClusterReport) GetPipelines() []*PipelineReport {
	if x != nil {
		return x.Pipelines
	}
	return nil
}

type StateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StateRequest) Reset() {
	*x = StateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateRequest) ProtoMessage() {}

func (x *StateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateRequest.ProtoReflect.Descriptor instead.
func (*StateRequest) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{4}
}

var File_monitor_proto protoreflect.FileDescriptor

var file_monitor_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x82, 0x03, 0x0a, 0x10, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x55, 0x75, 0x69, 0x64, 0x12, 0x32,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x2a, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2d,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67,
	0x65, 0x72, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x09, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x22, 0x3a, 0x0a, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0c, 0x0a,
	0x08, 0x46, 0x4f, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4c,
	0x45, 0x41, 0x52, 0x4e, 0x45, 0x52, 0x10, 0x03, 0x22, 0xae, 0x03, 0x0a, 0x0a, 0x4e, 0x6f, 0x64,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x6b, 0x5f,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0d, 0x64, 0x69, 0x73, 0x6b, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x64, 0x69, 0x73, 0x6b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2c, 0x0a,
	0x12, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x73,
	0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x53, 0x70, 0x65, 0x65, 0x64, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x63,
	0x76, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65, 0x63, 0x76, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x70, 0x75,
	0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a,
	0x63, 0x70, 0x75, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x27, 0x0a, 0x0f, 0x67, 0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x67, 0x6f, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x65, 0x74,
	0x61, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x6d, 0x65, 0x74, 0x61, 0x50, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xdb, 0x01, 0x0a, 0x0e, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x13, 0x0a, 0x05,
	0x70, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x67, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65,
	0x6e, 0x67, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x22, 0x45, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x44, 0x4f, 0x57, 0x4e, 0x5f, 0x47, 0x52, 0x41, 0x44, 0x45, 0x10, 0x02, 0x12, 0x0c,
	0x0a, 0x08, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x22, 0xc0, 0x02, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x3b, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65,
	0x6e, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x31, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e,
	0x67, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x09, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x22, 0x4b, 0x0a,
	0x0c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x45,
	0x41, 0x4c, 0x54, 0x48, 0x5f, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x48, 0x45, 0x41,
	0x4c, 0x54, 0x48, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x48, 0x45,
	0x41, 0x4c, 0x54, 0x48, 0x5f, 0x45, 0x52, 0x52, 0x10, 0x03, 0x22, 0x0e, 0x0a, 0x0c, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xcb, 0x01, 0x0a, 0x07, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x3a, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x1b, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x11, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x00, 0x12, 0x3c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1b, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x00,
	0x12, 0x46, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x73, 0x61, 0x74, 0x77,
	0x65, 0x61, 0x76, 0x65, 0x2f, 0x73, 0x61, 0x74, 0x2d, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_monitor_proto_rawDescOnce sync.Once
	file_monitor_proto_rawDescData = file_monitor_proto_rawDesc
)

func file_monitor_proto_rawDescGZIP() []byte {
	file_monitor_proto_rawDescOnce.Do(func() {
		file_monitor_proto_rawDescData = protoimpl.X.CompressGZIP(file_monitor_proto_rawDescData)
	})
	return file_monitor_proto_rawDescData
}

var file_monitor_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_monitor_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_monitor_proto_goTypes = []interface{}{
	(NodeStatusReport_Role)(0),      // 0: messenger.NodeStatusReport.Role
	(PipelineReport_State)(0),       // 1: messenger.PipelineReport.State
	(ClusterReport_ClusterState)(0), // 2: messenger.ClusterReport.ClusterState
	(*NodeStatusReport)(nil),        // 3: messenger.NodeStatusReport
	(*NodeStatus)(nil),              // 4: messenger.NodeStatus
	(*PipelineReport)(nil),          // 5: messenger.PipelineReport
	(*ClusterReport)(nil),           // 6: messenger.ClusterReport
	(*StateRequest)(nil),            // 7: messenger.StateRequest
	(*timestamp.Timestamp)(nil),     // 8: messenger.Timestamp
	(infos.NodeState)(0),            // 9: messenger.NodeState
	(*infos.ClusterInfo)(nil),       // 10: messenger.ClusterInfo
	(*emptypb.Empty)(nil),           // 11: google.protobuf.Empty
	(*common.Result)(nil),           // 12: messenger.Result
}
var file_monitor_proto_depIdxs = []int32{
	8,  // 0: messenger.NodeStatusReport.timestamp:type_name -> messenger.Timestamp
	9,  // 1: messenger.NodeStatusReport.state:type_name -> messenger.NodeState
	4,  // 2: messenger.NodeStatusReport.status:type_name -> messenger.NodeStatus
	0,  // 3: messenger.NodeStatusReport.role:type_name -> messenger.NodeStatusReport.Role
	5,  // 4: messenger.NodeStatusReport.pipelines:type_name -> messenger.PipelineReport
	1,  // 5: messenger.PipelineReport.state:type_name -> messenger.PipelineReport.State
	2,  // 6: messenger.ClusterReport.state:type_name -> messenger.ClusterReport.ClusterState
	10, // 7: messenger.ClusterReport.cluster_info:type_name -> messenger.ClusterInfo
	3,  // 8: messenger.ClusterReport.nodes:type_name -> messenger.NodeStatusReport
	5,  // 9: messenger.ClusterReport.pipelines:type_name -> messenger.PipelineReport
	3,  // 10: messenger.Monitor.Report:input_type -> messenger.NodeStatusReport
	11, // 11: messenger.Monitor.Get:input_type -> google.protobuf.Empty
	11, // 12: messenger.Monitor.GetClusterReport:input_type -> google.protobuf.Empty
	12, // 13: messenger.Monitor.Report:output_type -> messenger.Result
	3,  // 14: messenger.Monitor.Get:output_type -> messenger.NodeStatusReport
	6,  // 15: messenger.Monitor.GetClusterReport:output_type -> messenger.ClusterReport
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_monitor_proto_init() }
func file_monitor_proto_init() {
	if File_monitor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_monitor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStatusReport); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineReport); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterReport); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_monitor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_monitor_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_monitor_proto_goTypes,
		DependencyIndexes: file_monitor_proto_depIdxs,
		EnumInfos:         file_monitor_proto_enumTypes,
		MessageInfos:      file_monitor_proto_msgTypes,
	}.Build()
	File_monitor_proto = out.File
	file_monitor_proto_rawDesc = nil
	file_monitor_proto_goTypes = nil
	file_monitor_proto_depIdxs = nil
}
