// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package vpsproxy

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VPSProxyServiceClient is the client API for VPSProxyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VPSProxyServiceClient interface {
	GetHttpProxy(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetHttpProxyResponse, error)
	UpdateHttpProxy(ctx context.Context, in *UpdateHttpProxyReq, opts ...grpc.CallOption) (*empty.Empty, error)
	Shutdown(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type vPSProxyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVPSProxyServiceClient(cc grpc.ClientConnInterface) VPSProxyServiceClient {
	return &vPSProxyServiceClient{cc}
}

func (c *vPSProxyServiceClient) GetHttpProxy(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetHttpProxyResponse, error) {
	out := new(GetHttpProxyResponse)
	err := c.cc.Invoke(ctx, "/vpsproxy.VPSProxyService/GetHttpProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vPSProxyServiceClient) UpdateHttpProxy(ctx context.Context, in *UpdateHttpProxyReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/vpsproxy.VPSProxyService/UpdateHttpProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vPSProxyServiceClient) Shutdown(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/vpsproxy.VPSProxyService/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VPSProxyServiceServer is the server API for VPSProxyService service.
// All implementations must embed UnimplementedVPSProxyServiceServer
// for forward compatibility
type VPSProxyServiceServer interface {
	GetHttpProxy(context.Context, *empty.Empty) (*GetHttpProxyResponse, error)
	UpdateHttpProxy(context.Context, *UpdateHttpProxyReq) (*empty.Empty, error)
	Shutdown(context.Context, *empty.Empty) (*empty.Empty, error)
	mustEmbedUnimplementedVPSProxyServiceServer()
}

// UnimplementedVPSProxyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVPSProxyServiceServer struct {
}

func (UnimplementedVPSProxyServiceServer) GetHttpProxy(context.Context, *empty.Empty) (*GetHttpProxyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHttpProxy not implemented")
}
func (UnimplementedVPSProxyServiceServer) UpdateHttpProxy(context.Context, *UpdateHttpProxyReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHttpProxy not implemented")
}
func (UnimplementedVPSProxyServiceServer) Shutdown(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}
func (UnimplementedVPSProxyServiceServer) mustEmbedUnimplementedVPSProxyServiceServer() {}

// UnsafeVPSProxyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VPSProxyServiceServer will
// result in compilation errors.
type UnsafeVPSProxyServiceServer interface {
	mustEmbedUnimplementedVPSProxyServiceServer()
}

func RegisterVPSProxyServiceServer(s grpc.ServiceRegistrar, srv VPSProxyServiceServer) {
	s.RegisterService(&VPSProxyService_ServiceDesc, srv)
}

func _VPSProxyService_GetHttpProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VPSProxyServiceServer).GetHttpProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vpsproxy.VPSProxyService/GetHttpProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VPSProxyServiceServer).GetHttpProxy(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VPSProxyService_UpdateHttpProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHttpProxyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VPSProxyServiceServer).UpdateHttpProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vpsproxy.VPSProxyService/UpdateHttpProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VPSProxyServiceServer).UpdateHttpProxy(ctx, req.(*UpdateHttpProxyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VPSProxyService_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VPSProxyServiceServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vpsproxy.VPSProxyService/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VPSProxyServiceServer).Shutdown(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// VPSProxyService_ServiceDesc is the grpc.ServiceDesc for VPSProxyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VPSProxyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vpsproxy.VPSProxyService",
	HandlerType: (*VPSProxyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHttpProxy",
			Handler:    _VPSProxyService_GetHttpProxy_Handler,
		},
		{
			MethodName: "UpdateHttpProxy",
			Handler:    _VPSProxyService_UpdateHttpProxy_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _VPSProxyService_Shutdown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vpsproxy.proto",
}
