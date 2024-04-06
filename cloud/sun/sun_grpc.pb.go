// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: sun.proto

package sun

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	common "satweave/messenger/common"
	infos "satweave/sat-node/infos"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Sun_MoonRegister_FullMethodName          = "/messenger.Sun/MoonRegister"
	Sun_GetLeaderInfo_FullMethodName         = "/messenger.Sun/GetLeaderInfo"
	Sun_ReportClusterInfo_FullMethodName     = "/messenger.Sun/ReportClusterInfo"
	Sun_SubmitJob_FullMethodName             = "/messenger.Sun/SubmitJob"
	Sun_RegisterTaskManager_FullMethodName   = "/messenger.Sun/RegisterTaskManager"
	Sun_SaveSnapShot_FullMethodName          = "/messenger.Sun/SaveSnapShot"
	Sun_RestoreFromCheckpoint_FullMethodName = "/messenger.Sun/RestoreFromCheckpoint"
	Sun_ReceiverStreamData_FullMethodName    = "/messenger.Sun/ReceiverStreamData"
)

// SunClient is the client API for Sun service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SunClient interface {
	MoonRegister(ctx context.Context, in *infos.NodeInfo, opts ...grpc.CallOption) (*RegisterResult, error)
	GetLeaderInfo(ctx context.Context, in *infos.NodeInfo, opts ...grpc.CallOption) (*infos.NodeInfo, error)
	ReportClusterInfo(ctx context.Context, in *infos.ClusterInfo, opts ...grpc.CallOption) (*common.Result, error)
	// from user front client
	SubmitJob(ctx context.Context, in *SubmitJobRequest, opts ...grpc.CallOption) (*SubmitJobResponse, error)
	// from task manager
	RegisterTaskManager(ctx context.Context, in *RegisterTaskManagerRequest, opts ...grpc.CallOption) (*RegisterTaskManagerResponse, error)
	SaveSnapShot(ctx context.Context, in *SaveSnapShotRequest, opts ...grpc.CallOption) (*common.Result, error)
	RestoreFromCheckpoint(ctx context.Context, in *RestoreFromCheckpointRequest, opts ...grpc.CallOption) (*RestoreFromCheckpointResponse, error)
	ReceiverStreamData(ctx context.Context, in *ReceiverStreamDataRequest, opts ...grpc.CallOption) (*common.Result, error)
}

type sunClient struct {
	cc grpc.ClientConnInterface
}

func NewSunClient(cc grpc.ClientConnInterface) SunClient {
	return &sunClient{cc}
}

func (c *sunClient) MoonRegister(ctx context.Context, in *infos.NodeInfo, opts ...grpc.CallOption) (*RegisterResult, error) {
	out := new(RegisterResult)
	err := c.cc.Invoke(ctx, Sun_MoonRegister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sunClient) GetLeaderInfo(ctx context.Context, in *infos.NodeInfo, opts ...grpc.CallOption) (*infos.NodeInfo, error) {
	out := new(infos.NodeInfo)
	err := c.cc.Invoke(ctx, Sun_GetLeaderInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sunClient) ReportClusterInfo(ctx context.Context, in *infos.ClusterInfo, opts ...grpc.CallOption) (*common.Result, error) {
	out := new(common.Result)
	err := c.cc.Invoke(ctx, Sun_ReportClusterInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sunClient) SubmitJob(ctx context.Context, in *SubmitJobRequest, opts ...grpc.CallOption) (*SubmitJobResponse, error) {
	out := new(SubmitJobResponse)
	err := c.cc.Invoke(ctx, Sun_SubmitJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sunClient) RegisterTaskManager(ctx context.Context, in *RegisterTaskManagerRequest, opts ...grpc.CallOption) (*RegisterTaskManagerResponse, error) {
	out := new(RegisterTaskManagerResponse)
	err := c.cc.Invoke(ctx, Sun_RegisterTaskManager_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sunClient) SaveSnapShot(ctx context.Context, in *SaveSnapShotRequest, opts ...grpc.CallOption) (*common.Result, error) {
	out := new(common.Result)
	err := c.cc.Invoke(ctx, Sun_SaveSnapShot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sunClient) RestoreFromCheckpoint(ctx context.Context, in *RestoreFromCheckpointRequest, opts ...grpc.CallOption) (*RestoreFromCheckpointResponse, error) {
	out := new(RestoreFromCheckpointResponse)
	err := c.cc.Invoke(ctx, Sun_RestoreFromCheckpoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sunClient) ReceiverStreamData(ctx context.Context, in *ReceiverStreamDataRequest, opts ...grpc.CallOption) (*common.Result, error) {
	out := new(common.Result)
	err := c.cc.Invoke(ctx, Sun_ReceiverStreamData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SunServer is the server API for Sun service.
// All implementations must embed UnimplementedSunServer
// for forward compatibility
type SunServer interface {
	MoonRegister(context.Context, *infos.NodeInfo) (*RegisterResult, error)
	GetLeaderInfo(context.Context, *infos.NodeInfo) (*infos.NodeInfo, error)
	ReportClusterInfo(context.Context, *infos.ClusterInfo) (*common.Result, error)
	// from user front client
	SubmitJob(context.Context, *SubmitJobRequest) (*SubmitJobResponse, error)
	// from task manager
	RegisterTaskManager(context.Context, *RegisterTaskManagerRequest) (*RegisterTaskManagerResponse, error)
	SaveSnapShot(context.Context, *SaveSnapShotRequest) (*common.Result, error)
	RestoreFromCheckpoint(context.Context, *RestoreFromCheckpointRequest) (*RestoreFromCheckpointResponse, error)
	ReceiverStreamData(context.Context, *ReceiverStreamDataRequest) (*common.Result, error)
	mustEmbedUnimplementedSunServer()
}

// UnimplementedSunServer must be embedded to have forward compatible implementations.
type UnimplementedSunServer struct {
}

func (UnimplementedSunServer) MoonRegister(context.Context, *infos.NodeInfo) (*RegisterResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoonRegister not implemented")
}
func (UnimplementedSunServer) GetLeaderInfo(context.Context, *infos.NodeInfo) (*infos.NodeInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeaderInfo not implemented")
}
func (UnimplementedSunServer) ReportClusterInfo(context.Context, *infos.ClusterInfo) (*common.Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportClusterInfo not implemented")
}
func (UnimplementedSunServer) SubmitJob(context.Context, *SubmitJobRequest) (*SubmitJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitJob not implemented")
}
func (UnimplementedSunServer) RegisterTaskManager(context.Context, *RegisterTaskManagerRequest) (*RegisterTaskManagerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterTaskManager not implemented")
}
func (UnimplementedSunServer) SaveSnapShot(context.Context, *SaveSnapShotRequest) (*common.Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveSnapShot not implemented")
}
func (UnimplementedSunServer) RestoreFromCheckpoint(context.Context, *RestoreFromCheckpointRequest) (*RestoreFromCheckpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestoreFromCheckpoint not implemented")
}
func (UnimplementedSunServer) ReceiverStreamData(context.Context, *ReceiverStreamDataRequest) (*common.Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiverStreamData not implemented")
}
func (UnimplementedSunServer) mustEmbedUnimplementedSunServer() {}

// UnsafeSunServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SunServer will
// result in compilation errors.
type UnsafeSunServer interface {
	mustEmbedUnimplementedSunServer()
}

func RegisterSunServer(s grpc.ServiceRegistrar, srv SunServer) {
	s.RegisterService(&Sun_ServiceDesc, srv)
}

func _Sun_MoonRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(infos.NodeInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SunServer).MoonRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sun_MoonRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SunServer).MoonRegister(ctx, req.(*infos.NodeInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sun_GetLeaderInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(infos.NodeInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SunServer).GetLeaderInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sun_GetLeaderInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SunServer).GetLeaderInfo(ctx, req.(*infos.NodeInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sun_ReportClusterInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(infos.ClusterInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SunServer).ReportClusterInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sun_ReportClusterInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SunServer).ReportClusterInfo(ctx, req.(*infos.ClusterInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sun_SubmitJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SunServer).SubmitJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sun_SubmitJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SunServer).SubmitJob(ctx, req.(*SubmitJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sun_RegisterTaskManager_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterTaskManagerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SunServer).RegisterTaskManager(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sun_RegisterTaskManager_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SunServer).RegisterTaskManager(ctx, req.(*RegisterTaskManagerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sun_SaveSnapShot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveSnapShotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SunServer).SaveSnapShot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sun_SaveSnapShot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SunServer).SaveSnapShot(ctx, req.(*SaveSnapShotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sun_RestoreFromCheckpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreFromCheckpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SunServer).RestoreFromCheckpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sun_RestoreFromCheckpoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SunServer).RestoreFromCheckpoint(ctx, req.(*RestoreFromCheckpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sun_ReceiverStreamData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiverStreamDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SunServer).ReceiverStreamData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sun_ReceiverStreamData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SunServer).ReceiverStreamData(ctx, req.(*ReceiverStreamDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sun_ServiceDesc is the grpc.ServiceDesc for Sun service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sun_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messenger.Sun",
	HandlerType: (*SunServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MoonRegister",
			Handler:    _Sun_MoonRegister_Handler,
		},
		{
			MethodName: "GetLeaderInfo",
			Handler:    _Sun_GetLeaderInfo_Handler,
		},
		{
			MethodName: "ReportClusterInfo",
			Handler:    _Sun_ReportClusterInfo_Handler,
		},
		{
			MethodName: "SubmitJob",
			Handler:    _Sun_SubmitJob_Handler,
		},
		{
			MethodName: "RegisterTaskManager",
			Handler:    _Sun_RegisterTaskManager_Handler,
		},
		{
			MethodName: "SaveSnapShot",
			Handler:    _Sun_SaveSnapShot_Handler,
		},
		{
			MethodName: "RestoreFromCheckpoint",
			Handler:    _Sun_RestoreFromCheckpoint_Handler,
		},
		{
			MethodName: "ReceiverStreamData",
			Handler:    _Sun_ReceiverStreamData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sun.proto",
}
