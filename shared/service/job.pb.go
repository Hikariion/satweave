// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: job.proto

package service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// Job 的状态
type JobStatus int32

const (
	// 任务待处理
	JobStatus_PENDING JobStatus = 0
	// 任务已经完成
	JobStatus_FINISHED JobStatus = 1
	// 任务已经转交给其他卫星
	JobStatus_ASSIGNED JobStatus = 2
)

var JobStatus_name = map[int32]string{
	0: "PENDING",
	1: "FINISHED",
	2: "ASSIGNED",
}

var JobStatus_value = map[string]int32{
	"PENDING":  0,
	"FINISHED": 1,
	"ASSIGNED": 2,
}

func (x JobStatus) String() string {
	return proto.EnumName(JobStatus_name, int32(x))
}

func (JobStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{0}
}

// 定义一个离线任务
type Job struct {
	// 发送任务的client ip
	ClientIp string `protobuf:"bytes,1,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	// clint Rpc 的端口
	ClientPort uint64 `protobuf:"varint,2,opt,name=client_port,json=clientPort,proto3" json:"client_port,omitempty"`
	// 任务的id, 以每个客户端编号，然后加上时间戳
	JobId string `protobuf:"bytes,3,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// 任务的优先级
	Priority int32 `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
	// 任务的类型, 目前只有 yolov5
	ImageName string `protobuf:"bytes,5,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	// 任务的附件，这是一个文件
	Attachment string `protobuf:"bytes,6,opt,name=attachment,proto3" json:"attachment,omitempty"`
	// 附件大小，单位是字节
	AttachmentSize int32 `protobuf:"varint,7,opt,name=attachment_size,json=attachmentSize,proto3" json:"attachment_size,omitempty"`
	// 任务状态，枚举类型
	Status JobStatus `protobuf:"varint,8,opt,name=status,proto3,enum=messenger.JobStatus" json:"status,omitempty"`
	// TODO(qiu): 任务的参数
	// 执行命令的参数
	Command string `protobuf:"bytes,9,opt,name=command,proto3" json:"command,omitempty"`
	// 返回的产物名称
	ResultName           string   `protobuf:"bytes,10,opt,name=result_name,json=resultName,proto3" json:"result_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{0}
}
func (m *Job) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Job.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return m.Size()
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetClientIp() string {
	if m != nil {
		return m.ClientIp
	}
	return ""
}

func (m *Job) GetClientPort() uint64 {
	if m != nil {
		return m.ClientPort
	}
	return 0
}

func (m *Job) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *Job) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Job) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func (m *Job) GetAttachment() string {
	if m != nil {
		return m.Attachment
	}
	return ""
}

func (m *Job) GetAttachmentSize() int32 {
	if m != nil {
		return m.AttachmentSize
	}
	return 0
}

func (m *Job) GetStatus() JobStatus {
	if m != nil {
		return m.Status
	}
	return JobStatus_PENDING
}

func (m *Job) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Job) GetResultName() string {
	if m != nil {
		return m.ResultName
	}
	return ""
}

type UploadAttachmentReply struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadAttachmentReply) Reset()         { *m = UploadAttachmentReply{} }
func (m *UploadAttachmentReply) String() string { return proto.CompactTextString(m) }
func (*UploadAttachmentReply) ProtoMessage()    {}
func (*UploadAttachmentReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{1}
}
func (m *UploadAttachmentReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UploadAttachmentReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UploadAttachmentReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UploadAttachmentReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadAttachmentReply.Merge(m, src)
}
func (m *UploadAttachmentReply) XXX_Size() int {
	return m.Size()
}
func (m *UploadAttachmentReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadAttachmentReply.DiscardUnknown(m)
}

var xxx_messageInfo_UploadAttachmentReply proto.InternalMessageInfo

func (m *UploadAttachmentReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterEnum("messenger.JobStatus", JobStatus_name, JobStatus_value)
	proto.RegisterType((*Job)(nil), "messenger.Job")
	proto.RegisterType((*UploadAttachmentReply)(nil), "messenger.UploadAttachmentReply")
}

func init() { proto.RegisterFile("job.proto", fileDescriptor_f32c477d91a04ead) }

var fileDescriptor_f32c477d91a04ead = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0x87, 0xe7, 0x6e, 0x4d, 0x93, 0x33, 0x34, 0x2a, 0x8b, 0x09, 0x0b, 0x44, 0x88, 0x76, 0x43,
	0x40, 0xa8, 0x13, 0x7f, 0x5e, 0x60, 0xa8, 0x65, 0x64, 0x17, 0xd1, 0x94, 0x88, 0x1b, 0x6e, 0x22,
	0x27, 0x39, 0xda, 0x5c, 0x25, 0x71, 0x64, 0xbb, 0x43, 0xed, 0x93, 0xf0, 0x46, 0x70, 0xc9, 0x23,
	0xa0, 0xf2, 0x22, 0x28, 0x4e, 0x9b, 0x72, 0xf9, 0xfb, 0x8e, 0x8f, 0xcf, 0xe7, 0x63, 0xf0, 0x96,
	0x32, 0x9f, 0xb5, 0x4a, 0x1a, 0x49, 0xbd, 0x1a, 0xb5, 0xc6, 0xe6, 0x0e, 0xd5, 0xc5, 0xcf, 0x11,
	0x1c, 0xdf, 0xc8, 0x9c, 0x3e, 0x07, 0xaf, 0xa8, 0x04, 0x36, 0x26, 0x13, 0x2d, 0x23, 0x01, 0x09,
	0xbd, 0xc4, 0xed, 0x41, 0xd4, 0xd2, 0x97, 0x70, 0xba, 0x2b, 0xb6, 0x52, 0x19, 0x36, 0x0a, 0x48,
	0x78, 0x92, 0x40, 0x8f, 0x6e, 0xa5, 0x32, 0xf4, 0x1c, 0x9c, 0xa5, 0xcc, 0x33, 0x51, 0xb2, 0x63,
	0xdb, 0x3a, 0x5e, 0xca, 0x3c, 0x2a, 0xe9, 0x33, 0x70, 0x5b, 0x25, 0xa4, 0x12, 0x66, 0xcd, 0x4e,
	0x02, 0x12, 0x8e, 0x93, 0x21, 0xd3, 0x17, 0x00, 0xa2, 0xe6, 0x77, 0x98, 0x35, 0xbc, 0x46, 0x36,
	0xb6, 0x6d, 0x9e, 0x25, 0x31, 0xaf, 0x91, 0xfa, 0x00, 0xdc, 0x18, 0x5e, 0xdc, 0xd7, 0xd8, 0x18,
	0xe6, 0xd8, 0xf2, 0x7f, 0x84, 0xbe, 0x82, 0xc7, 0x87, 0x94, 0x69, 0xb1, 0x41, 0x36, 0xb1, 0x13,
	0xce, 0x0e, 0x38, 0x15, 0x1b, 0xa4, 0x6f, 0xc1, 0xd1, 0x86, 0x9b, 0x95, 0x66, 0x6e, 0x40, 0xc2,
	0xb3, 0xf7, 0x4f, 0x66, 0xc3, 0xe3, 0x67, 0x37, 0x32, 0x4f, 0x6d, 0x2d, 0xd9, 0x9d, 0xa1, 0x0c,
	0x26, 0x85, 0xac, 0x6b, 0xde, 0x94, 0xcc, 0xb3, 0x33, 0xf7, 0xb1, 0xdb, 0x81, 0x42, 0xbd, 0xaa,
	0x4c, 0x2f, 0x0c, 0xbd, 0x51, 0x8f, 0x3a, 0xe3, 0x8b, 0x77, 0x70, 0xfe, 0xb5, 0xad, 0x24, 0x2f,
	0xaf, 0x06, 0x81, 0x04, 0xdb, 0x6a, 0xdd, 0xdd, 0xa9, 0x57, 0x45, 0x81, 0x5a, 0xdb, 0xc5, 0xba,
	0xc9, 0x3e, 0xbe, 0xf9, 0x08, 0xde, 0xa0, 0x40, 0x4f, 0x61, 0x72, 0xbb, 0x88, 0xe7, 0x51, 0x7c,
	0x3d, 0x3d, 0xa2, 0x8f, 0xc0, 0xfd, 0x1c, 0xc5, 0x51, 0xfa, 0x65, 0x31, 0x9f, 0x92, 0x2e, 0x5d,
	0xa5, 0x69, 0x74, 0x1d, 0x2f, 0xe6, 0xd3, 0xd1, 0xa7, 0xd7, 0xbf, 0xb6, 0x3e, 0xf9, 0xbd, 0xf5,
	0xc9, 0x9f, 0xad, 0x4f, 0x7e, 0xfc, 0xf5, 0x8f, 0xbe, 0x3d, 0xd5, 0xdc, 0x7c, 0x47, 0xfe, 0x80,
	0x97, 0xfa, 0x9e, 0x2b, 0x2c, 0x2f, 0x35, 0xaa, 0x07, 0x51, 0x60, 0xee, 0xd8, 0xff, 0xfe, 0xf0,
	0x2f, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x0d, 0x69, 0x55, 0xfc, 0x01, 0x00, 0x00,
}

func (m *Job) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Job) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Job) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ResultName) > 0 {
		i -= len(m.ResultName)
		copy(dAtA[i:], m.ResultName)
		i = encodeVarintJob(dAtA, i, uint64(len(m.ResultName)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Command) > 0 {
		i -= len(m.Command)
		copy(dAtA[i:], m.Command)
		i = encodeVarintJob(dAtA, i, uint64(len(m.Command)))
		i--
		dAtA[i] = 0x4a
	}
	if m.Status != 0 {
		i = encodeVarintJob(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x40
	}
	if m.AttachmentSize != 0 {
		i = encodeVarintJob(dAtA, i, uint64(m.AttachmentSize))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Attachment) > 0 {
		i -= len(m.Attachment)
		copy(dAtA[i:], m.Attachment)
		i = encodeVarintJob(dAtA, i, uint64(len(m.Attachment)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ImageName) > 0 {
		i -= len(m.ImageName)
		copy(dAtA[i:], m.ImageName)
		i = encodeVarintJob(dAtA, i, uint64(len(m.ImageName)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Priority != 0 {
		i = encodeVarintJob(dAtA, i, uint64(m.Priority))
		i--
		dAtA[i] = 0x20
	}
	if len(m.JobId) > 0 {
		i -= len(m.JobId)
		copy(dAtA[i:], m.JobId)
		i = encodeVarintJob(dAtA, i, uint64(len(m.JobId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ClientPort != 0 {
		i = encodeVarintJob(dAtA, i, uint64(m.ClientPort))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ClientIp) > 0 {
		i -= len(m.ClientIp)
		copy(dAtA[i:], m.ClientIp)
		i = encodeVarintJob(dAtA, i, uint64(len(m.ClientIp)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UploadAttachmentReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UploadAttachmentReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UploadAttachmentReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintJob(dAtA []byte, offset int, v uint64) int {
	offset -= sovJob(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Job) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientIp)
	if l > 0 {
		n += 1 + l + sovJob(uint64(l))
	}
	if m.ClientPort != 0 {
		n += 1 + sovJob(uint64(m.ClientPort))
	}
	l = len(m.JobId)
	if l > 0 {
		n += 1 + l + sovJob(uint64(l))
	}
	if m.Priority != 0 {
		n += 1 + sovJob(uint64(m.Priority))
	}
	l = len(m.ImageName)
	if l > 0 {
		n += 1 + l + sovJob(uint64(l))
	}
	l = len(m.Attachment)
	if l > 0 {
		n += 1 + l + sovJob(uint64(l))
	}
	if m.AttachmentSize != 0 {
		n += 1 + sovJob(uint64(m.AttachmentSize))
	}
	if m.Status != 0 {
		n += 1 + sovJob(uint64(m.Status))
	}
	l = len(m.Command)
	if l > 0 {
		n += 1 + l + sovJob(uint64(l))
	}
	l = len(m.ResultName)
	if l > 0 {
		n += 1 + l + sovJob(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UploadAttachmentReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovJob(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozJob(x uint64) (n int) {
	return sovJob(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Job) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJob
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
			return fmt.Errorf("proto: Job: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Job: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientIp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJob
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
				return ErrInvalidLengthJob
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientIp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientPort", wireType)
			}
			m.ClientPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClientPort |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJob
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
				return ErrInvalidLengthJob
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priority", wireType)
			}
			m.Priority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Priority |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImageName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJob
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
				return ErrInvalidLengthJob
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImageName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attachment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJob
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
				return ErrInvalidLengthJob
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attachment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttachmentSize", wireType)
			}
			m.AttachmentSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AttachmentSize |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= JobStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Command", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJob
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
				return ErrInvalidLengthJob
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Command = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResultName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJob
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
				return ErrInvalidLengthJob
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResultName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJob(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthJob
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
func (m *UploadAttachmentReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJob
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
			return fmt.Errorf("proto: UploadAttachmentReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UploadAttachmentReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipJob(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthJob
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
func skipJob(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowJob
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
					return 0, ErrIntOverflowJob
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
					return 0, ErrIntOverflowJob
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
				return 0, ErrInvalidLengthJob
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupJob
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthJob
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthJob        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowJob          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupJob = fmt.Errorf("proto: unexpected end of group")
)
