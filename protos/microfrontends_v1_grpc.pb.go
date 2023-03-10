// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: protos/microfrontends_v1.proto

package protos

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

// MicrofrontendsClient is the client API for Microfrontends service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MicrofrontendsClient interface {
	GetMicrofrontends(ctx context.Context, in *MicrofrontendPageRequest, opts ...grpc.CallOption) (*MicrofrontendPageReply, error)
	GetMicrofrontendById(ctx context.Context, in *MicrofrontendIdRequest, opts ...grpc.CallOption) (*MicrofrontendObjectReply, error)
	CreateMicrofrontend(ctx context.Context, in *MicrofrontendObjectRequest, opts ...grpc.CallOption) (*MicrofrontendObjectReply, error)
	UpdateMicrofrontend(ctx context.Context, in *MicrofrontendObjectRequest, opts ...grpc.CallOption) (*MicrofrontendObjectReply, error)
	DeleteMicrofrontendById(ctx context.Context, in *MicrofrontendIdRequest, opts ...grpc.CallOption) (*MicrofrontendObjectReply, error)
}

type microfrontendsClient struct {
	cc grpc.ClientConnInterface
}

func NewMicrofrontendsClient(cc grpc.ClientConnInterface) MicrofrontendsClient {
	return &microfrontendsClient{cc}
}

func (c *microfrontendsClient) GetMicrofrontends(ctx context.Context, in *MicrofrontendPageRequest, opts ...grpc.CallOption) (*MicrofrontendPageReply, error) {
	out := new(MicrofrontendPageReply)
	err := c.cc.Invoke(ctx, "/microfrontends_v1.Microfrontends/get_microfrontends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microfrontendsClient) GetMicrofrontendById(ctx context.Context, in *MicrofrontendIdRequest, opts ...grpc.CallOption) (*MicrofrontendObjectReply, error) {
	out := new(MicrofrontendObjectReply)
	err := c.cc.Invoke(ctx, "/microfrontends_v1.Microfrontends/get_microfrontend_by_id", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microfrontendsClient) CreateMicrofrontend(ctx context.Context, in *MicrofrontendObjectRequest, opts ...grpc.CallOption) (*MicrofrontendObjectReply, error) {
	out := new(MicrofrontendObjectReply)
	err := c.cc.Invoke(ctx, "/microfrontends_v1.Microfrontends/create_microfrontend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microfrontendsClient) UpdateMicrofrontend(ctx context.Context, in *MicrofrontendObjectRequest, opts ...grpc.CallOption) (*MicrofrontendObjectReply, error) {
	out := new(MicrofrontendObjectReply)
	err := c.cc.Invoke(ctx, "/microfrontends_v1.Microfrontends/update_microfrontend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microfrontendsClient) DeleteMicrofrontendById(ctx context.Context, in *MicrofrontendIdRequest, opts ...grpc.CallOption) (*MicrofrontendObjectReply, error) {
	out := new(MicrofrontendObjectReply)
	err := c.cc.Invoke(ctx, "/microfrontends_v1.Microfrontends/delete_microfrontend_by_id", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MicrofrontendsServer is the server API for Microfrontends service.
// All implementations must embed UnimplementedMicrofrontendsServer
// for forward compatibility
type MicrofrontendsServer interface {
	GetMicrofrontends(context.Context, *MicrofrontendPageRequest) (*MicrofrontendPageReply, error)
	GetMicrofrontendById(context.Context, *MicrofrontendIdRequest) (*MicrofrontendObjectReply, error)
	CreateMicrofrontend(context.Context, *MicrofrontendObjectRequest) (*MicrofrontendObjectReply, error)
	UpdateMicrofrontend(context.Context, *MicrofrontendObjectRequest) (*MicrofrontendObjectReply, error)
	DeleteMicrofrontendById(context.Context, *MicrofrontendIdRequest) (*MicrofrontendObjectReply, error)
	mustEmbedUnimplementedMicrofrontendsServer()
}

// UnimplementedMicrofrontendsServer must be embedded to have forward compatible implementations.
type UnimplementedMicrofrontendsServer struct {
}

func (UnimplementedMicrofrontendsServer) GetMicrofrontends(context.Context, *MicrofrontendPageRequest) (*MicrofrontendPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMicrofrontends not implemented")
}
func (UnimplementedMicrofrontendsServer) GetMicrofrontendById(context.Context, *MicrofrontendIdRequest) (*MicrofrontendObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMicrofrontendById not implemented")
}
func (UnimplementedMicrofrontendsServer) CreateMicrofrontend(context.Context, *MicrofrontendObjectRequest) (*MicrofrontendObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMicrofrontend not implemented")
}
func (UnimplementedMicrofrontendsServer) UpdateMicrofrontend(context.Context, *MicrofrontendObjectRequest) (*MicrofrontendObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMicrofrontend not implemented")
}
func (UnimplementedMicrofrontendsServer) DeleteMicrofrontendById(context.Context, *MicrofrontendIdRequest) (*MicrofrontendObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMicrofrontendById not implemented")
}
func (UnimplementedMicrofrontendsServer) mustEmbedUnimplementedMicrofrontendsServer() {}

// UnsafeMicrofrontendsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MicrofrontendsServer will
// result in compilation errors.
type UnsafeMicrofrontendsServer interface {
	mustEmbedUnimplementedMicrofrontendsServer()
}

func RegisterMicrofrontendsServer(s grpc.ServiceRegistrar, srv MicrofrontendsServer) {
	s.RegisterService(&Microfrontends_ServiceDesc, srv)
}

func _Microfrontends_GetMicrofrontends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MicrofrontendPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicrofrontendsServer).GetMicrofrontends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microfrontends_v1.Microfrontends/get_microfrontends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicrofrontendsServer).GetMicrofrontends(ctx, req.(*MicrofrontendPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Microfrontends_GetMicrofrontendById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MicrofrontendIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicrofrontendsServer).GetMicrofrontendById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microfrontends_v1.Microfrontends/get_microfrontend_by_id",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicrofrontendsServer).GetMicrofrontendById(ctx, req.(*MicrofrontendIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Microfrontends_CreateMicrofrontend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MicrofrontendObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicrofrontendsServer).CreateMicrofrontend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microfrontends_v1.Microfrontends/create_microfrontend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicrofrontendsServer).CreateMicrofrontend(ctx, req.(*MicrofrontendObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Microfrontends_UpdateMicrofrontend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MicrofrontendObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicrofrontendsServer).UpdateMicrofrontend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microfrontends_v1.Microfrontends/update_microfrontend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicrofrontendsServer).UpdateMicrofrontend(ctx, req.(*MicrofrontendObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Microfrontends_DeleteMicrofrontendById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MicrofrontendIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicrofrontendsServer).DeleteMicrofrontendById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microfrontends_v1.Microfrontends/delete_microfrontend_by_id",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicrofrontendsServer).DeleteMicrofrontendById(ctx, req.(*MicrofrontendIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Microfrontends_ServiceDesc is the grpc.ServiceDesc for Microfrontends service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Microfrontends_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "microfrontends_v1.Microfrontends",
	HandlerType: (*MicrofrontendsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get_microfrontends",
			Handler:    _Microfrontends_GetMicrofrontends_Handler,
		},
		{
			MethodName: "get_microfrontend_by_id",
			Handler:    _Microfrontends_GetMicrofrontendById_Handler,
		},
		{
			MethodName: "create_microfrontend",
			Handler:    _Microfrontends_CreateMicrofrontend_Handler,
		},
		{
			MethodName: "update_microfrontend",
			Handler:    _Microfrontends_UpdateMicrofrontend_Handler,
		},
		{
			MethodName: "delete_microfrontend_by_id",
			Handler:    _Microfrontends_DeleteMicrofrontendById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/microfrontends_v1.proto",
}
