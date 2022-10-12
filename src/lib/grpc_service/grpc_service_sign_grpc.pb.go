// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc_service

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Sign(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error) {
	out := new(SignResponse)
	err := c.cc.Invoke(ctx, "/cryptopro_grpc.Service/Sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	Sign(context.Context, *SignRequest) (*SignResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) Sign(context.Context, *SignRequest) (*SignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cryptopro_grpc.Service/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Sign(ctx, req.(*SignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cryptopro_grpc.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sign",
			Handler:    _Service_Sign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_service_sign.proto",
}

// ServiceInternalClient is the client API for ServiceInternal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceInternalClient interface {
	Sign(ctx context.Context, opts ...grpc.CallOption) (ServiceInternal_SignClient, error)
}

type serviceInternalClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceInternalClient(cc grpc.ClientConnInterface) ServiceInternalClient {
	return &serviceInternalClient{cc}
}

func (c *serviceInternalClient) Sign(ctx context.Context, opts ...grpc.CallOption) (ServiceInternal_SignClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServiceInternal_ServiceDesc.Streams[0], "/cryptopro_grpc.ServiceInternal/Sign", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceInternalSignClient{stream}
	return x, nil
}

type ServiceInternal_SignClient interface {
	Send(*SignRequest) error
	Recv() (*SignResponse, error)
	grpc.ClientStream
}

type serviceInternalSignClient struct {
	grpc.ClientStream
}

func (x *serviceInternalSignClient) Send(m *SignRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceInternalSignClient) Recv() (*SignResponse, error) {
	m := new(SignResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceInternalServer is the server API for ServiceInternal service.
// All implementations must embed UnimplementedServiceInternalServer
// for forward compatibility
type ServiceInternalServer interface {
	Sign(ServiceInternal_SignServer) error
	mustEmbedUnimplementedServiceInternalServer()
}

// UnimplementedServiceInternalServer must be embedded to have forward compatible implementations.
type UnimplementedServiceInternalServer struct {
}

func (UnimplementedServiceInternalServer) Sign(ServiceInternal_SignServer) error {
	return status.Errorf(codes.Unimplemented, "method Sign not implemented")
}
func (UnimplementedServiceInternalServer) mustEmbedUnimplementedServiceInternalServer() {}

// UnsafeServiceInternalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceInternalServer will
// result in compilation errors.
type UnsafeServiceInternalServer interface {
	mustEmbedUnimplementedServiceInternalServer()
}

func RegisterServiceInternalServer(s grpc.ServiceRegistrar, srv ServiceInternalServer) {
	s.RegisterService(&ServiceInternal_ServiceDesc, srv)
}

func _ServiceInternal_Sign_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceInternalServer).Sign(&serviceInternalSignServer{stream})
}

type ServiceInternal_SignServer interface {
	Send(*SignResponse) error
	Recv() (*SignRequest, error)
	grpc.ServerStream
}

type serviceInternalSignServer struct {
	grpc.ServerStream
}

func (x *serviceInternalSignServer) Send(m *SignResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceInternalSignServer) Recv() (*SignRequest, error) {
	m := new(SignRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceInternal_ServiceDesc is the grpc.ServiceDesc for ServiceInternal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceInternal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cryptopro_grpc.ServiceInternal",
	HandlerType: (*ServiceInternalServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sign",
			Handler:       _ServiceInternal_Sign_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc_service_sign.proto",
}
