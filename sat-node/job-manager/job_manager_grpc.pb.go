// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: job_manager.proto

package job_manager

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	common "satweave/messenger/common"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// JobManagerServiceClient is the client API for JobManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobManagerServiceClient interface {
	// from user front client
	SubmitJob(ctx context.Context, in *SubmitJobRequest, opts ...grpc.CallOption) (*common.NilResponse, error)
}

type jobManagerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobManagerServiceClient(cc grpc.ClientConnInterface) JobManagerServiceClient {
	return &jobManagerServiceClient{cc}
}

func (c *jobManagerServiceClient) SubmitJob(ctx context.Context, in *SubmitJobRequest, opts ...grpc.CallOption) (*common.NilResponse, error) {
	out := new(common.NilResponse)
	err := c.cc.Invoke(ctx, "/messenger.JobManagerService/SubmitJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobManagerServiceServer is the server API for JobManagerService service.
// All implementations must embed UnimplementedJobManagerServiceServer
// for forward compatibility
type JobManagerServiceServer interface {
	// from user front client
	SubmitJob(context.Context, *SubmitJobRequest) (*common.NilResponse, error)
	mustEmbedUnimplementedJobManagerServiceServer()
}

// UnimplementedJobManagerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJobManagerServiceServer struct {
}

func (UnimplementedJobManagerServiceServer) SubmitJob(context.Context, *SubmitJobRequest) (*common.NilResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitJob not implemented")
}
func (UnimplementedJobManagerServiceServer) mustEmbedUnimplementedJobManagerServiceServer() {}

// UnsafeJobManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobManagerServiceServer will
// result in compilation errors.
type UnsafeJobManagerServiceServer interface {
	mustEmbedUnimplementedJobManagerServiceServer()
}

func RegisterJobManagerServiceServer(s grpc.ServiceRegistrar, srv JobManagerServiceServer) {
	s.RegisterService(&JobManagerService_ServiceDesc, srv)
}

func _JobManagerService_SubmitJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobManagerServiceServer).SubmitJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messenger.JobManagerService/SubmitJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobManagerServiceServer).SubmitJob(ctx, req.(*SubmitJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JobManagerService_ServiceDesc is the grpc.ServiceDesc for JobManagerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobManagerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messenger.JobManagerService",
	HandlerType: (*JobManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitJob",
			Handler:    _JobManagerService_SubmitJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job_manager.proto",
}
