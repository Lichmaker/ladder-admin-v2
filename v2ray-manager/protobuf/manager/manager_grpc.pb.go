// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package manager

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	command "github.com/v2fly/v2ray-core/v5/app/stats/command"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServeServiceClient is the client API for ServeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServeServiceClient interface {
	SetNode(ctx context.Context, in *SetNodeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	AddNode(ctx context.Context, in *AddNodeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteNode(ctx context.Context, in *DeleteNodeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SetUser(ctx context.Context, in *SetUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SysStat(ctx context.Context, in *SysStatRequest, opts ...grpc.CallOption) (*command.SysStatsResponse, error)
	QueryStats(ctx context.Context, in *QueryStatsRequest, opts ...grpc.CallOption) (*command.QueryStatsResponse, error)
}

type serveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServeServiceClient(cc grpc.ClientConnInterface) ServeServiceClient {
	return &serveServiceClient{cc}
}

func (c *serveServiceClient) SetNode(ctx context.Context, in *SetNodeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v2ray.manager.ServeService/SetNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serveServiceClient) AddNode(ctx context.Context, in *AddNodeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v2ray.manager.ServeService/AddNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serveServiceClient) DeleteNode(ctx context.Context, in *DeleteNodeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v2ray.manager.ServeService/DeleteNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serveServiceClient) SetUser(ctx context.Context, in *SetUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v2ray.manager.ServeService/SetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serveServiceClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v2ray.manager.ServeService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serveServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v2ray.manager.ServeService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serveServiceClient) SysStat(ctx context.Context, in *SysStatRequest, opts ...grpc.CallOption) (*command.SysStatsResponse, error) {
	out := new(command.SysStatsResponse)
	err := c.cc.Invoke(ctx, "/v2ray.manager.ServeService/SysStat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serveServiceClient) QueryStats(ctx context.Context, in *QueryStatsRequest, opts ...grpc.CallOption) (*command.QueryStatsResponse, error) {
	out := new(command.QueryStatsResponse)
	err := c.cc.Invoke(ctx, "/v2ray.manager.ServeService/QueryStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServeServiceServer is the server API for ServeService service.
// All implementations must embed UnimplementedServeServiceServer
// for forward compatibility
type ServeServiceServer interface {
	SetNode(context.Context, *SetNodeRequest) (*empty.Empty, error)
	AddNode(context.Context, *AddNodeRequest) (*empty.Empty, error)
	DeleteNode(context.Context, *DeleteNodeRequest) (*empty.Empty, error)
	SetUser(context.Context, *SetUserRequest) (*empty.Empty, error)
	AddUser(context.Context, *AddUserRequest) (*empty.Empty, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*empty.Empty, error)
	SysStat(context.Context, *SysStatRequest) (*command.SysStatsResponse, error)
	QueryStats(context.Context, *QueryStatsRequest) (*command.QueryStatsResponse, error)
	mustEmbedUnimplementedServeServiceServer()
}

// UnimplementedServeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServeServiceServer struct {
}

func (UnimplementedServeServiceServer) SetNode(context.Context, *SetNodeRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetNode not implemented")
}
func (UnimplementedServeServiceServer) AddNode(context.Context, *AddNodeRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNode not implemented")
}
func (UnimplementedServeServiceServer) DeleteNode(context.Context, *DeleteNodeRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNode not implemented")
}
func (UnimplementedServeServiceServer) SetUser(context.Context, *SetUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUser not implemented")
}
func (UnimplementedServeServiceServer) AddUser(context.Context, *AddUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedServeServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedServeServiceServer) SysStat(context.Context, *SysStatRequest) (*command.SysStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysStat not implemented")
}
func (UnimplementedServeServiceServer) QueryStats(context.Context, *QueryStatsRequest) (*command.QueryStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryStats not implemented")
}
func (UnimplementedServeServiceServer) mustEmbedUnimplementedServeServiceServer() {}

// UnsafeServeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServeServiceServer will
// result in compilation errors.
type UnsafeServeServiceServer interface {
	mustEmbedUnimplementedServeServiceServer()
}

func RegisterServeServiceServer(s grpc.ServiceRegistrar, srv ServeServiceServer) {
	s.RegisterService(&ServeService_ServiceDesc, srv)
}

func _ServeService_SetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServeServiceServer).SetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2ray.manager.ServeService/SetNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServeServiceServer).SetNode(ctx, req.(*SetNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServeService_AddNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServeServiceServer).AddNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2ray.manager.ServeService/AddNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServeServiceServer).AddNode(ctx, req.(*AddNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServeService_DeleteNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServeServiceServer).DeleteNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2ray.manager.ServeService/DeleteNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServeServiceServer).DeleteNode(ctx, req.(*DeleteNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServeService_SetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServeServiceServer).SetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2ray.manager.ServeService/SetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServeServiceServer).SetUser(ctx, req.(*SetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServeService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServeServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2ray.manager.ServeService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServeServiceServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServeService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServeServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2ray.manager.ServeService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServeServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServeService_SysStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysStatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServeServiceServer).SysStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2ray.manager.ServeService/SysStat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServeServiceServer).SysStat(ctx, req.(*SysStatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServeService_QueryStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServeServiceServer).QueryStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2ray.manager.ServeService/QueryStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServeServiceServer).QueryStats(ctx, req.(*QueryStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServeService_ServiceDesc is the grpc.ServiceDesc for ServeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v2ray.manager.ServeService",
	HandlerType: (*ServeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetNode",
			Handler:    _ServeService_SetNode_Handler,
		},
		{
			MethodName: "AddNode",
			Handler:    _ServeService_AddNode_Handler,
		},
		{
			MethodName: "DeleteNode",
			Handler:    _ServeService_DeleteNode_Handler,
		},
		{
			MethodName: "SetUser",
			Handler:    _ServeService_SetUser_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _ServeService_AddUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _ServeService_DeleteUser_Handler,
		},
		{
			MethodName: "SysStat",
			Handler:    _ServeService_SysStat_Handler,
		},
		{
			MethodName: "QueryStats",
			Handler:    _ServeService_QueryStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager.proto",
}