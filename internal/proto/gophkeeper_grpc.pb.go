// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: gophkeeper.proto

package gophkeeper

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

const (
	Gophkeeper_HandleLogin_FullMethodName        = "/api.Gophkeeper/HandleLogin"
	Gophkeeper_HandleRegistration_FullMethodName = "/api.Gophkeeper/HandleRegistration"
	Gophkeeper_HandleCreateText_FullMethodName   = "/api.Gophkeeper/HandleCreateText"
	Gophkeeper_HandleGetListText_FullMethodName  = "/api.Gophkeeper/HandleGetListText"
	Gophkeeper_HandleGetNodeText_FullMethodName  = "/api.Gophkeeper/HandleGetNodeText"
	Gophkeeper_HandlePing_FullMethodName         = "/api.Gophkeeper/HandlePing"
)

// GophkeeperClient is the client API for Gophkeeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GophkeeperClient interface {
	HandleLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	HandleRegistration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error)
	HandleCreateText(ctx context.Context, in *CreateTextRequest, opts ...grpc.CallOption) (*CreateTextResponse, error)
	HandleGetListText(ctx context.Context, in *GetListTextRequest, opts ...grpc.CallOption) (*GetListTextResponse, error)
	HandleGetNodeText(ctx context.Context, in *GetNodeTextRequest, opts ...grpc.CallOption) (*GetNodeTextResponse, error)
	HandlePing(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type gophkeeperClient struct {
	cc grpc.ClientConnInterface
}

func NewGophkeeperClient(cc grpc.ClientConnInterface) GophkeeperClient {
	return &gophkeeperClient{cc}
}

func (c *gophkeeperClient) HandleLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_HandleLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) HandleRegistration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error) {
	out := new(RegistrationResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_HandleRegistration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) HandleCreateText(ctx context.Context, in *CreateTextRequest, opts ...grpc.CallOption) (*CreateTextResponse, error) {
	out := new(CreateTextResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_HandleCreateText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) HandleGetListText(ctx context.Context, in *GetListTextRequest, opts ...grpc.CallOption) (*GetListTextResponse, error) {
	out := new(GetListTextResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_HandleGetListText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) HandleGetNodeText(ctx context.Context, in *GetNodeTextRequest, opts ...grpc.CallOption) (*GetNodeTextResponse, error) {
	out := new(GetNodeTextResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_HandleGetNodeText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) HandlePing(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_HandlePing_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GophkeeperServer is the server API for Gophkeeper service.
// All implementations must embed UnimplementedGophkeeperServer
// for forward compatibility
type GophkeeperServer interface {
	HandleLogin(context.Context, *LoginRequest) (*LoginResponse, error)
	HandleRegistration(context.Context, *RegistrationRequest) (*RegistrationResponse, error)
	HandleCreateText(context.Context, *CreateTextRequest) (*CreateTextResponse, error)
	HandleGetListText(context.Context, *GetListTextRequest) (*GetListTextResponse, error)
	HandleGetNodeText(context.Context, *GetNodeTextRequest) (*GetNodeTextResponse, error)
	HandlePing(context.Context, *PingRequest) (*PingResponse, error)
	mustEmbedUnimplementedGophkeeperServer()
}

// UnimplementedGophkeeperServer must be embedded to have forward compatible implementations.
type UnimplementedGophkeeperServer struct {
}

func (UnimplementedGophkeeperServer) HandleLogin(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleLogin not implemented")
}
func (UnimplementedGophkeeperServer) HandleRegistration(context.Context, *RegistrationRequest) (*RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleRegistration not implemented")
}
func (UnimplementedGophkeeperServer) HandleCreateText(context.Context, *CreateTextRequest) (*CreateTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleCreateText not implemented")
}
func (UnimplementedGophkeeperServer) HandleGetListText(context.Context, *GetListTextRequest) (*GetListTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleGetListText not implemented")
}
func (UnimplementedGophkeeperServer) HandleGetNodeText(context.Context, *GetNodeTextRequest) (*GetNodeTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleGetNodeText not implemented")
}
func (UnimplementedGophkeeperServer) HandlePing(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandlePing not implemented")
}
func (UnimplementedGophkeeperServer) mustEmbedUnimplementedGophkeeperServer() {}

// UnsafeGophkeeperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GophkeeperServer will
// result in compilation errors.
type UnsafeGophkeeperServer interface {
	mustEmbedUnimplementedGophkeeperServer()
}

func RegisterGophkeeperServer(s grpc.ServiceRegistrar, srv GophkeeperServer) {
	s.RegisterService(&Gophkeeper_ServiceDesc, srv)
}

func _Gophkeeper_HandleLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandleLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_HandleLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandleLogin(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_HandleRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandleRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_HandleRegistration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandleRegistration(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_HandleCreateText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandleCreateText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_HandleCreateText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandleCreateText(ctx, req.(*CreateTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_HandleGetListText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandleGetListText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_HandleGetListText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandleGetListText(ctx, req.(*GetListTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_HandleGetNodeText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandleGetNodeText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_HandleGetNodeText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandleGetNodeText(ctx, req.(*GetNodeTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_HandlePing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).HandlePing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_HandlePing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).HandlePing(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gophkeeper_ServiceDesc is the grpc.ServiceDesc for Gophkeeper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gophkeeper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Gophkeeper",
	HandlerType: (*GophkeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleLogin",
			Handler:    _Gophkeeper_HandleLogin_Handler,
		},
		{
			MethodName: "HandleRegistration",
			Handler:    _Gophkeeper_HandleRegistration_Handler,
		},
		{
			MethodName: "HandleCreateText",
			Handler:    _Gophkeeper_HandleCreateText_Handler,
		},
		{
			MethodName: "HandleGetListText",
			Handler:    _Gophkeeper_HandleGetListText_Handler,
		},
		{
			MethodName: "HandleGetNodeText",
			Handler:    _Gophkeeper_HandleGetNodeText_Handler,
		},
		{
			MethodName: "HandlePing",
			Handler:    _Gophkeeper_HandlePing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gophkeeper.proto",
}
