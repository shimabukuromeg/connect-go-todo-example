// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/todo/v1/todo.proto

package todov1

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

// ToDoServiceClient is the client API for ToDoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ToDoServiceClient interface {
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	UpdateTaskStatus(ctx context.Context, in *UpdateTaskStatusRequest, opts ...grpc.CallOption) (*UpdateTaskStatusResponse, error)
	DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*DeleteTaskResponse, error)
}

type toDoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewToDoServiceClient(cc grpc.ClientConnInterface) ToDoServiceClient {
	return &toDoServiceClient{cc}
}

func (c *toDoServiceClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	out := new(CreateTaskResponse)
	err := c.cc.Invoke(ctx, "/proto.todo.v1.ToDoService/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) UpdateTaskStatus(ctx context.Context, in *UpdateTaskStatusRequest, opts ...grpc.CallOption) (*UpdateTaskStatusResponse, error) {
	out := new(UpdateTaskStatusResponse)
	err := c.cc.Invoke(ctx, "/proto.todo.v1.ToDoService/UpdateTaskStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) DeleteTask(ctx context.Context, in *DeleteTaskRequest, opts ...grpc.CallOption) (*DeleteTaskResponse, error) {
	out := new(DeleteTaskResponse)
	err := c.cc.Invoke(ctx, "/proto.todo.v1.ToDoService/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToDoServiceServer is the server API for ToDoService service.
// All implementations must embed UnimplementedToDoServiceServer
// for forward compatibility
type ToDoServiceServer interface {
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	UpdateTaskStatus(context.Context, *UpdateTaskStatusRequest) (*UpdateTaskStatusResponse, error)
	DeleteTask(context.Context, *DeleteTaskRequest) (*DeleteTaskResponse, error)
	mustEmbedUnimplementedToDoServiceServer()
}

// UnimplementedToDoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedToDoServiceServer struct {
}

func (UnimplementedToDoServiceServer) CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedToDoServiceServer) UpdateTaskStatus(context.Context, *UpdateTaskStatusRequest) (*UpdateTaskStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTaskStatus not implemented")
}
func (UnimplementedToDoServiceServer) DeleteTask(context.Context, *DeleteTaskRequest) (*DeleteTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedToDoServiceServer) mustEmbedUnimplementedToDoServiceServer() {}

// UnsafeToDoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ToDoServiceServer will
// result in compilation errors.
type UnsafeToDoServiceServer interface {
	mustEmbedUnimplementedToDoServiceServer()
}

func RegisterToDoServiceServer(s grpc.ServiceRegistrar, srv ToDoServiceServer) {
	s.RegisterService(&ToDoService_ServiceDesc, srv)
}

func _ToDoService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.todo.v1.ToDoService/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_UpdateTaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).UpdateTaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.todo.v1.ToDoService/UpdateTaskStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).UpdateTaskStatus(ctx, req.(*UpdateTaskStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.todo.v1.ToDoService/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).DeleteTask(ctx, req.(*DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ToDoService_ServiceDesc is the grpc.ServiceDesc for ToDoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ToDoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.todo.v1.ToDoService",
	HandlerType: (*ToDoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _ToDoService_CreateTask_Handler,
		},
		{
			MethodName: "UpdateTaskStatus",
			Handler:    _ToDoService_UpdateTaskStatus_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _ToDoService_DeleteTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/todo/v1/todo.proto",
}