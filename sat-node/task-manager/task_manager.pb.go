// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: task_manager.proto

package task_manager

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	_ "satweave/messenger/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ExecuteTask struct {
	ClsName              string   `protobuf:"bytes,1,opt,name=cls_name,json=clsName,proto3" json:"cls_name,omitempty"`
	InputEndpoints       string   `protobuf:"bytes,2,opt,name=input_endpoints,json=inputEndpoints,proto3" json:"input_endpoints,omitempty"`
	OutputEndpoints      string   `protobuf:"bytes,3,opt,name=output_endpoints,json=outputEndpoints,proto3" json:"output_endpoints,omitempty"`
	Resources            *File    `protobuf:"bytes,4,opt,name=resources,proto3" json:"resources,omitempty"`
	TaskFile             *File    `protobuf:"bytes,5,opt,name=task_file,json=taskFile,proto3" json:"task_file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteTask) Reset()         { *m = ExecuteTask{} }
func (m *ExecuteTask) String() string { return proto.CompactTextString(m) }
func (*ExecuteTask) ProtoMessage()    {}
func (*ExecuteTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_370859c40635d1e2, []int{0}
}
func (m *ExecuteTask) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExecuteTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExecuteTask.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExecuteTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteTask.Merge(m, src)
}
func (m *ExecuteTask) XXX_Size() int {
	return m.Size()
}
func (m *ExecuteTask) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteTask.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteTask proto.InternalMessageInfo

func (m *ExecuteTask) GetClsName() string {
	if m != nil {
		return m.ClsName
	}
	return ""
}

func (m *ExecuteTask) GetInputEndpoints() string {
	if m != nil {
		return m.InputEndpoints
	}
	return ""
}

func (m *ExecuteTask) GetOutputEndpoints() string {
	if m != nil {
		return m.OutputEndpoints
	}
	return ""
}

func (m *ExecuteTask) GetResources() *File {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *ExecuteTask) GetTaskFile() *File {
	if m != nil {
		return m.TaskFile
	}
	return nil
}

type Task struct {
	ClsName              string   `protobuf:"bytes,1,opt,name=cls_name,json=clsName,proto3" json:"cls_name,omitempty"`
	Currency             int32    `protobuf:"varint,2,opt,name=currency,proto3" json:"currency,omitempty"`
	InputTasks           []string `protobuf:"bytes,3,rep,name=input_tasks,json=inputTasks,proto3" json:"input_tasks,omitempty"`
	Resources            []*File  `protobuf:"bytes,4,rep,name=resources,proto3" json:"resources,omitempty"`
	TaskFile             *File    `protobuf:"bytes,5,opt,name=task_file,json=taskFile,proto3" json:"task_file,omitempty"`
	Locate               string   `protobuf:"bytes,6,opt,name=locate,proto3" json:"locate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_370859c40635d1e2, []int{1}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Task.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return m.Size()
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetClsName() string {
	if m != nil {
		return m.ClsName
	}
	return ""
}

func (m *Task) GetCurrency() int32 {
	if m != nil {
		return m.Currency
	}
	return 0
}

func (m *Task) GetInputTasks() []string {
	if m != nil {
		return m.InputTasks
	}
	return nil
}

func (m *Task) GetResources() []*File {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Task) GetTaskFile() *File {
	if m != nil {
		return m.TaskFile
	}
	return nil
}

func (m *Task) GetLocate() string {
	if m != nil {
		return m.Locate
	}
	return ""
}

type File struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_370859c40635d1e2, []int{2}
}
func (m *File) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_File.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return m.Size()
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *File) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type DeployTaskRequest struct {
	ExecTask             *ExecuteTask `protobuf:"bytes,1,opt,name=exec_task,json=execTask,proto3" json:"exec_task,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DeployTaskRequest) Reset()         { *m = DeployTaskRequest{} }
func (m *DeployTaskRequest) String() string { return proto.CompactTextString(m) }
func (*DeployTaskRequest) ProtoMessage()    {}
func (*DeployTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_370859c40635d1e2, []int{3}
}
func (m *DeployTaskRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeployTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeployTaskRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeployTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployTaskRequest.Merge(m, src)
}
func (m *DeployTaskRequest) XXX_Size() int {
	return m.Size()
}
func (m *DeployTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeployTaskRequest proto.InternalMessageInfo

func (m *DeployTaskRequest) GetExecTask() *ExecuteTask {
	if m != nil {
		return m.ExecTask
	}
	return nil
}

func init() {
	proto.RegisterType((*ExecuteTask)(nil), "messenger.ExecuteTask")
	proto.RegisterType((*Task)(nil), "messenger.Task")
	proto.RegisterType((*File)(nil), "messenger.File")
	proto.RegisterType((*DeployTaskRequest)(nil), "messenger.DeployTaskRequest")
}

func init() { proto.RegisterFile("task_manager.proto", fileDescriptor_370859c40635d1e2) }

var fileDescriptor_370859c40635d1e2 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0xd7, 0x34, 0xdb, 0x4d, 0x26, 0x88, 0x82, 0x0f, 0xab, 0x50, 0xa1, 0xb0, 0xca, 0x85,
	0x22, 0xd1, 0x2e, 0x6a, 0x79, 0x02, 0xd4, 0x22, 0x2e, 0xf4, 0x10, 0x38, 0xf5, 0x52, 0x19, 0x77,
	0xa8, 0xa2, 0x26, 0x76, 0xb0, 0x9d, 0xd2, 0xbe, 0x09, 0x8f, 0xc4, 0x11, 0x89, 0x13, 0x37, 0x54,
	0x5e, 0x04, 0xd9, 0xe9, 0x9f, 0xa0, 0x22, 0x71, 0xd9, 0x53, 0x66, 0xc6, 0x3f, 0x7d, 0xf3, 0x7d,
	0xca, 0x00, 0x35, 0x4c, 0xaf, 0xe6, 0x05, 0x13, 0x6c, 0x89, 0x6a, 0x50, 0x2a, 0x69, 0x24, 0x0d,
	0x0a, 0xd4, 0x1a, 0xc5, 0x12, 0x55, 0xf7, 0x3e, 0x97, 0x45, 0x21, 0x45, 0xfd, 0x90, 0xfc, 0x24,
	0x10, 0x4e, 0x36, 0xc8, 0x2b, 0x83, 0x1f, 0x98, 0x5e, 0xd1, 0xc7, 0xe0, 0xf3, 0x5c, 0xcf, 0x05,
	0x2b, 0x30, 0x22, 0x37, 0xa4, 0x17, 0xa4, 0x57, 0x3c, 0xd7, 0x53, 0x56, 0x20, 0x7d, 0x06, 0x9d,
	0x4c, 0x94, 0x95, 0x99, 0xa3, 0x58, 0x94, 0x32, 0x13, 0x46, 0x47, 0xf7, 0x1c, 0xf1, 0xc0, 0x8d,
	0x27, 0x87, 0x29, 0x7d, 0x0e, 0x0f, 0x65, 0x65, 0xfe, 0x26, 0x5b, 0x8e, 0xec, 0xd4, 0xf3, 0x13,
	0xda, 0x87, 0x40, 0xa1, 0x96, 0x95, 0xe2, 0xa8, 0x23, 0xef, 0x86, 0xf4, 0xc2, 0x61, 0x67, 0x70,
	0xf4, 0x3a, 0x78, 0x93, 0xe5, 0x98, 0x9e, 0x08, 0xfa, 0x02, 0x02, 0x17, 0xee, 0x53, 0x96, 0x63,
	0x74, 0xf9, 0x6f, 0xdc, 0xb7, 0x84, 0xad, 0x92, 0x1f, 0x04, 0xbc, 0xff, 0x85, 0xea, 0x82, 0xcf,
	0x2b, 0xa5, 0x50, 0xf0, 0xad, 0x4b, 0x73, 0x99, 0x1e, 0x7b, 0xfa, 0x14, 0xc2, 0x3a, 0xb0, 0x55,
	0xb4, 0x11, 0x5a, 0xbd, 0x20, 0x05, 0x37, 0xb2, 0xb2, 0x67, 0xee, 0x5b, 0x77, 0xe9, 0x9e, 0x5e,
	0x43, 0x3b, 0x97, 0x9c, 0x19, 0x8c, 0xda, 0xce, 0xf2, 0xbe, 0x4b, 0x5e, 0x81, 0xe7, 0xde, 0x29,
	0x78, 0x8d, 0x40, 0xae, 0xa6, 0x11, 0x5c, 0x71, 0x29, 0x0c, 0x0a, 0xb3, 0xff, 0x35, 0x87, 0x36,
	0x79, 0x0b, 0x8f, 0xc6, 0x58, 0xe6, 0x72, 0x6b, 0x9d, 0xa7, 0xf8, 0xb9, 0x42, 0x6d, 0xe8, 0x08,
	0x02, 0xdc, 0x20, 0x77, 0xf9, 0x9c, 0x4e, 0x38, 0xbc, 0x6e, 0x18, 0x6a, 0xdc, 0x45, 0xea, 0x5b,
	0xd0, 0x56, 0xc3, 0x19, 0x50, 0xfb, 0x7d, 0x57, 0xdf, 0xd7, 0x7b, 0x54, 0xeb, 0x8c, 0x23, 0x1d,
	0x03, 0x9c, 0xf4, 0xe9, 0x93, 0x86, 0xca, 0xd9, 0xda, 0x6e, 0x73, 0xc7, 0x34, 0xcb, 0x53, 0xd4,
	0xa5, 0x14, 0x1a, 0x93, 0x8b, 0xd7, 0x2f, 0xbf, 0xed, 0x62, 0xf2, 0x7d, 0x17, 0x93, 0x5f, 0xbb,
	0x98, 0x7c, 0xfd, 0x1d, 0x5f, 0xcc, 0x62, 0xcd, 0xcc, 0x17, 0x64, 0x6b, 0xbc, 0xd5, 0xcc, 0xf4,
	0x85, 0x5c, 0xe0, 0xad, 0x75, 0xdb, 0xdf, 0x9f, 0xf7, 0xc7, 0xb6, 0x3b, 0xe3, 0xd1, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xcf, 0xc2, 0xb6, 0x3a, 0xf5, 0x02, 0x00, 0x00,
}

func (m *ExecuteTask) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExecuteTask) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExecuteTask) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.TaskFile != nil {
		{
			size, err := m.TaskFile.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTaskManager(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Resources != nil {
		{
			size, err := m.Resources.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTaskManager(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.OutputEndpoints) > 0 {
		i -= len(m.OutputEndpoints)
		copy(dAtA[i:], m.OutputEndpoints)
		i = encodeVarintTaskManager(dAtA, i, uint64(len(m.OutputEndpoints)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.InputEndpoints) > 0 {
		i -= len(m.InputEndpoints)
		copy(dAtA[i:], m.InputEndpoints)
		i = encodeVarintTaskManager(dAtA, i, uint64(len(m.InputEndpoints)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClsName) > 0 {
		i -= len(m.ClsName)
		copy(dAtA[i:], m.ClsName)
		i = encodeVarintTaskManager(dAtA, i, uint64(len(m.ClsName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Task) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Task) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Task) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Locate) > 0 {
		i -= len(m.Locate)
		copy(dAtA[i:], m.Locate)
		i = encodeVarintTaskManager(dAtA, i, uint64(len(m.Locate)))
		i--
		dAtA[i] = 0x32
	}
	if m.TaskFile != nil {
		{
			size, err := m.TaskFile.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTaskManager(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Resources) > 0 {
		for iNdEx := len(m.Resources) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Resources[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTaskManager(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.InputTasks) > 0 {
		for iNdEx := len(m.InputTasks) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.InputTasks[iNdEx])
			copy(dAtA[i:], m.InputTasks[iNdEx])
			i = encodeVarintTaskManager(dAtA, i, uint64(len(m.InputTasks[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Currency != 0 {
		i = encodeVarintTaskManager(dAtA, i, uint64(m.Currency))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ClsName) > 0 {
		i -= len(m.ClsName)
		copy(dAtA[i:], m.ClsName)
		i = encodeVarintTaskManager(dAtA, i, uint64(len(m.ClsName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *File) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *File) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *File) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintTaskManager(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTaskManager(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DeployTaskRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeployTaskRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeployTaskRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ExecTask != nil {
		{
			size, err := m.ExecTask.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTaskManager(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTaskManager(dAtA []byte, offset int, v uint64) int {
	offset -= sovTaskManager(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExecuteTask) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClsName)
	if l > 0 {
		n += 1 + l + sovTaskManager(uint64(l))
	}
	l = len(m.InputEndpoints)
	if l > 0 {
		n += 1 + l + sovTaskManager(uint64(l))
	}
	l = len(m.OutputEndpoints)
	if l > 0 {
		n += 1 + l + sovTaskManager(uint64(l))
	}
	if m.Resources != nil {
		l = m.Resources.Size()
		n += 1 + l + sovTaskManager(uint64(l))
	}
	if m.TaskFile != nil {
		l = m.TaskFile.Size()
		n += 1 + l + sovTaskManager(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Task) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClsName)
	if l > 0 {
		n += 1 + l + sovTaskManager(uint64(l))
	}
	if m.Currency != 0 {
		n += 1 + sovTaskManager(uint64(m.Currency))
	}
	if len(m.InputTasks) > 0 {
		for _, s := range m.InputTasks {
			l = len(s)
			n += 1 + l + sovTaskManager(uint64(l))
		}
	}
	if len(m.Resources) > 0 {
		for _, e := range m.Resources {
			l = e.Size()
			n += 1 + l + sovTaskManager(uint64(l))
		}
	}
	if m.TaskFile != nil {
		l = m.TaskFile.Size()
		n += 1 + l + sovTaskManager(uint64(l))
	}
	l = len(m.Locate)
	if l > 0 {
		n += 1 + l + sovTaskManager(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *File) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTaskManager(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovTaskManager(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DeployTaskRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExecTask != nil {
		l = m.ExecTask.Size()
		n += 1 + l + sovTaskManager(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTaskManager(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTaskManager(x uint64) (n int) {
	return sovTaskManager(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExecuteTask) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTaskManager
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExecuteTask: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExecuteTask: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClsName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaskManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTaskManager
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTaskManager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClsName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InputEndpoints", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaskManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTaskManager
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTaskManager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InputEndpoints = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutputEndpoints", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaskManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTaskManager
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTaskManager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutputEndpoints = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaskManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTaskManager
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTaskManager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Resources == nil {
				m.Resources = &File{}
			}
			if err := m.Resources.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskFile", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaskManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTaskManager
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTaskManager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TaskFile == nil {
				m.TaskFile = &File{}
			}
			if err := m.TaskFile.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTaskManager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTaskManager
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Task) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTaskManager
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Task: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Task: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClsName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaskManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTaskManager
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTaskManager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClsName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Currency", wireType)
			}
			m.Currency = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaskManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Currency |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InputTasks", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaskManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTaskManager
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTaskManager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InputTasks = append(m.InputTasks, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaskManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTaskManager
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTaskManager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resources = append(m.Resources, &File{})
			if err := m.Resources[len(m.Resources)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskFile", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaskManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTaskManager
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTaskManager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TaskFile == nil {
				m.TaskFile = &File{}
			}
			if err := m.TaskFile.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Locate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaskManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTaskManager
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTaskManager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Locate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTaskManager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTaskManager
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *File) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTaskManager
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: File: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: File: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaskManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTaskManager
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTaskManager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaskManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTaskManager
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTaskManager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTaskManager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTaskManager
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeployTaskRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTaskManager
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeployTaskRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeployTaskRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecTask", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaskManager
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTaskManager
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTaskManager
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExecTask == nil {
				m.ExecTask = &ExecuteTask{}
			}
			if err := m.ExecTask.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTaskManager(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTaskManager
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTaskManager(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTaskManager
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTaskManager
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTaskManager
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTaskManager
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTaskManager
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTaskManager
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTaskManager        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTaskManager          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTaskManager = fmt.Errorf("proto: unexpected end of group")
)
