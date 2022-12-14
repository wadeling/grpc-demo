// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: task.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TaskProxyClient is the client API for TaskProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskProxyClient interface {
	GetTask(ctx context.Context, in *Req, opts ...grpc.CallOption) (TaskProxy_GetTaskClient, error)
}

type taskProxyClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskProxyClient(cc grpc.ClientConnInterface) TaskProxyClient {
	return &taskProxyClient{cc}
}

func (c *taskProxyClient) GetTask(ctx context.Context, in *Req, opts ...grpc.CallOption) (TaskProxy_GetTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &TaskProxy_ServiceDesc.Streams[0], "/pb.TaskProxy/GetTask", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskProxyGetTaskClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TaskProxy_GetTaskClient interface {
	Recv() (*Task, error)
	grpc.ClientStream
}

type taskProxyGetTaskClient struct {
	grpc.ClientStream
}

func (x *taskProxyGetTaskClient) Recv() (*Task, error) {
	m := new(Task)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TaskProxyServer is the server API for TaskProxy service.
// All implementations must embed UnimplementedTaskProxyServer
// for forward compatibility
type TaskProxyServer interface {
	GetTask(*Req, TaskProxy_GetTaskServer) error
	mustEmbedUnimplementedTaskProxyServer()
}

// UnimplementedTaskProxyServer must be embedded to have forward compatible implementations.
type UnimplementedTaskProxyServer struct {
}

func (UnimplementedTaskProxyServer) GetTask(*Req, TaskProxy_GetTaskServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedTaskProxyServer) mustEmbedUnimplementedTaskProxyServer() {}

// UnsafeTaskProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskProxyServer will
// result in compilation errors.
type UnsafeTaskProxyServer interface {
	mustEmbedUnimplementedTaskProxyServer()
}

func RegisterTaskProxyServer(s grpc.ServiceRegistrar, srv TaskProxyServer) {
	s.RegisterService(&TaskProxy_ServiceDesc, srv)
}

func _TaskProxy_GetTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Req)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskProxyServer).GetTask(m, &taskProxyGetTaskServer{stream})
}

type TaskProxy_GetTaskServer interface {
	Send(*Task) error
	grpc.ServerStream
}

type taskProxyGetTaskServer struct {
	grpc.ServerStream
}

func (x *taskProxyGetTaskServer) Send(m *Task) error {
	return x.ServerStream.SendMsg(m)
}

// TaskProxy_ServiceDesc is the grpc.ServiceDesc for TaskProxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskProxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TaskProxy",
	HandlerType: (*TaskProxyServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTask",
			Handler:       _TaskProxy_GetTask_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "task.proto",
}
