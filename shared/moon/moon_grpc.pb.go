// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package moon

import (
	context "context"
	raftpb "go.etcd.io/etcd/raft/v3/raftpb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MoonClient is the client API for Moon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MoonClient interface {
	SendRaftMessage(ctx context.Context, in *raftpb.Message, opts ...grpc.CallOption) (*raftpb.Message, error)
	// ProposeInfo 用于向集群中提交一个 Info 共识，该 Info 将会在集群所有节点上同步
	// 目前可以提交的 Info: NodeInfo, BucketInfo
	ProposeInfo(ctx context.Context, in *ProposeInfoRequest, opts ...grpc.CallOption) (*ProposeInfoReply, error)
	// GetInfo 从对应 InfoStorage 中取得 Info
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoReply, error)
	// ListInfo 从对应 InfoStorage 中取得指定前缀的 Info 列表
	ListInfo(ctx context.Context, in *ListInfoRequest, opts ...grpc.CallOption) (*ListInfoReply, error)
}

type moonClient struct {
	cc grpc.ClientConnInterface
}

func NewMoonClient(cc grpc.ClientConnInterface) MoonClient {
	return &moonClient{cc}
}

func (c *moonClient) SendRaftMessage(ctx context.Context, in *raftpb.Message, opts ...grpc.CallOption) (*raftpb.Message, error) {
	out := new(raftpb.Message)
	err := c.cc.Invoke(ctx, "/messenger.Moon/SendRaftMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moonClient) ProposeInfo(ctx context.Context, in *ProposeInfoRequest, opts ...grpc.CallOption) (*ProposeInfoReply, error) {
	out := new(ProposeInfoReply)
	err := c.cc.Invoke(ctx, "/messenger.Moon/ProposeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moonClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoReply, error) {
	out := new(GetInfoReply)
	err := c.cc.Invoke(ctx, "/messenger.Moon/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moonClient) ListInfo(ctx context.Context, in *ListInfoRequest, opts ...grpc.CallOption) (*ListInfoReply, error) {
	out := new(ListInfoReply)
	err := c.cc.Invoke(ctx, "/messenger.Moon/ListInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MoonServer is the server API for Moon service.
// All implementations must embed UnimplementedMoonServer
// for forward compatibility
type MoonServer interface {
	SendRaftMessage(context.Context, *raftpb.Message) (*raftpb.Message, error)
	// ProposeInfo 用于向集群中提交一个 Info 共识，该 Info 将会在集群所有节点上同步
	// 目前可以提交的 Info: NodeInfo, BucketInfo
	ProposeInfo(context.Context, *ProposeInfoRequest) (*ProposeInfoReply, error)
	// GetInfo 从对应 InfoStorage 中取得 Info
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoReply, error)
	// ListInfo 从对应 InfoStorage 中取得指定前缀的 Info 列表
	ListInfo(context.Context, *ListInfoRequest) (*ListInfoReply, error)
	mustEmbedUnimplementedMoonServer()
}

// UnimplementedMoonServer must be embedded to have forward compatible implementations.
type UnimplementedMoonServer struct {
}

func (UnimplementedMoonServer) SendRaftMessage(context.Context, *raftpb.Message) (*raftpb.Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRaftMessage not implemented")
}
func (UnimplementedMoonServer) ProposeInfo(context.Context, *ProposeInfoRequest) (*ProposeInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProposeInfo not implemented")
}
func (UnimplementedMoonServer) GetInfo(context.Context, *GetInfoRequest) (*GetInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (UnimplementedMoonServer) ListInfo(context.Context, *ListInfoRequest) (*ListInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInfo not implemented")
}
func (UnimplementedMoonServer) mustEmbedUnimplementedMoonServer() {}

// UnsafeMoonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MoonServer will
// result in compilation errors.
type UnsafeMoonServer interface {
	mustEmbedUnimplementedMoonServer()
}

func RegisterMoonServer(s grpc.ServiceRegistrar, srv MoonServer) {
	s.RegisterService(&Moon_ServiceDesc, srv)
}

func _Moon_SendRaftMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raftpb.Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoonServer).SendRaftMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messenger.Moon/SendRaftMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoonServer).SendRaftMessage(ctx, req.(*raftpb.Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moon_ProposeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProposeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoonServer).ProposeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messenger.Moon/ProposeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoonServer).ProposeInfo(ctx, req.(*ProposeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moon_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoonServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messenger.Moon/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoonServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Moon_ListInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoonServer).ListInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messenger.Moon/ListInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoonServer).ListInfo(ctx, req.(*ListInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Moon_ServiceDesc is the grpc.ServiceDesc for Moon service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Moon_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messenger.Moon",
	HandlerType: (*MoonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendRaftMessage",
			Handler:    _Moon_SendRaftMessage_Handler,
		},
		{
			MethodName: "ProposeInfo",
			Handler:    _Moon_ProposeInfo_Handler,
		},
		{
			MethodName: "GetInfo",
			Handler:    _Moon_GetInfo_Handler,
		},
		{
			MethodName: "ListInfo",
			Handler:    _Moon_ListInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moon.proto",
}
