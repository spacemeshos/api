// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: spacemesh/v2/activation.proto

package spacemeshv2

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
	ActivationStreamService_Stream_FullMethodName        = "/spacemesh.v2.ActivationStreamService/Stream"
	ActivationStreamService_StreamHeaders_FullMethodName = "/spacemesh.v2.ActivationStreamService/StreamHeaders"
)

// ActivationStreamServiceClient is the client API for ActivationStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivationStreamServiceClient interface {
	Stream(ctx context.Context, in *ActivationStreamRequest, opts ...grpc.CallOption) (ActivationStreamService_StreamClient, error)
	StreamHeaders(ctx context.Context, in *ActivationStreamRequest, opts ...grpc.CallOption) (ActivationStreamService_StreamHeadersClient, error)
}

type activationStreamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActivationStreamServiceClient(cc grpc.ClientConnInterface) ActivationStreamServiceClient {
	return &activationStreamServiceClient{cc}
}

func (c *activationStreamServiceClient) Stream(ctx context.Context, in *ActivationStreamRequest, opts ...grpc.CallOption) (ActivationStreamService_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ActivationStreamService_ServiceDesc.Streams[0], ActivationStreamService_Stream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &activationStreamServiceStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ActivationStreamService_StreamClient interface {
	Recv() (*Activation, error)
	grpc.ClientStream
}

type activationStreamServiceStreamClient struct {
	grpc.ClientStream
}

func (x *activationStreamServiceStreamClient) Recv() (*Activation, error) {
	m := new(Activation)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *activationStreamServiceClient) StreamHeaders(ctx context.Context, in *ActivationStreamRequest, opts ...grpc.CallOption) (ActivationStreamService_StreamHeadersClient, error) {
	stream, err := c.cc.NewStream(ctx, &ActivationStreamService_ServiceDesc.Streams[1], ActivationStreamService_StreamHeaders_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &activationStreamServiceStreamHeadersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ActivationStreamService_StreamHeadersClient interface {
	Recv() (*ActivationHeader, error)
	grpc.ClientStream
}

type activationStreamServiceStreamHeadersClient struct {
	grpc.ClientStream
}

func (x *activationStreamServiceStreamHeadersClient) Recv() (*ActivationHeader, error) {
	m := new(ActivationHeader)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ActivationStreamServiceServer is the server API for ActivationStreamService service.
// All implementations should embed UnimplementedActivationStreamServiceServer
// for forward compatibility
type ActivationStreamServiceServer interface {
	Stream(*ActivationStreamRequest, ActivationStreamService_StreamServer) error
	StreamHeaders(*ActivationStreamRequest, ActivationStreamService_StreamHeadersServer) error
}

// UnimplementedActivationStreamServiceServer should be embedded to have forward compatible implementations.
type UnimplementedActivationStreamServiceServer struct {
}

func (UnimplementedActivationStreamServiceServer) Stream(*ActivationStreamRequest, ActivationStreamService_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (UnimplementedActivationStreamServiceServer) StreamHeaders(*ActivationStreamRequest, ActivationStreamService_StreamHeadersServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamHeaders not implemented")
}

// UnsafeActivationStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivationStreamServiceServer will
// result in compilation errors.
type UnsafeActivationStreamServiceServer interface {
	mustEmbedUnimplementedActivationStreamServiceServer()
}

func RegisterActivationStreamServiceServer(s grpc.ServiceRegistrar, srv ActivationStreamServiceServer) {
	s.RegisterService(&ActivationStreamService_ServiceDesc, srv)
}

func _ActivationStreamService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ActivationStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ActivationStreamServiceServer).Stream(m, &activationStreamServiceStreamServer{stream})
}

type ActivationStreamService_StreamServer interface {
	Send(*Activation) error
	grpc.ServerStream
}

type activationStreamServiceStreamServer struct {
	grpc.ServerStream
}

func (x *activationStreamServiceStreamServer) Send(m *Activation) error {
	return x.ServerStream.SendMsg(m)
}

func _ActivationStreamService_StreamHeaders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ActivationStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ActivationStreamServiceServer).StreamHeaders(m, &activationStreamServiceStreamHeadersServer{stream})
}

type ActivationStreamService_StreamHeadersServer interface {
	Send(*ActivationHeader) error
	grpc.ServerStream
}

type activationStreamServiceStreamHeadersServer struct {
	grpc.ServerStream
}

func (x *activationStreamServiceStreamHeadersServer) Send(m *ActivationHeader) error {
	return x.ServerStream.SendMsg(m)
}

// ActivationStreamService_ServiceDesc is the grpc.ServiceDesc for ActivationStreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActivationStreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spacemesh.v2.ActivationStreamService",
	HandlerType: (*ActivationStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _ActivationStreamService_Stream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamHeaders",
			Handler:       _ActivationStreamService_StreamHeaders_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "spacemesh/v2/activation.proto",
}

const (
	ActivationService_List_FullMethodName        = "/spacemesh.v2.ActivationService/List"
	ActivationService_ListHeaders_FullMethodName = "/spacemesh.v2.ActivationService/ListHeaders"
)

// ActivationServiceClient is the client API for ActivationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivationServiceClient interface {
	List(ctx context.Context, in *ActivationRequest, opts ...grpc.CallOption) (*ActivationList, error)
	ListHeaders(ctx context.Context, in *ActivationRequest, opts ...grpc.CallOption) (*ActivationHeaderList, error)
}

type activationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActivationServiceClient(cc grpc.ClientConnInterface) ActivationServiceClient {
	return &activationServiceClient{cc}
}

func (c *activationServiceClient) List(ctx context.Context, in *ActivationRequest, opts ...grpc.CallOption) (*ActivationList, error) {
	out := new(ActivationList)
	err := c.cc.Invoke(ctx, ActivationService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activationServiceClient) ListHeaders(ctx context.Context, in *ActivationRequest, opts ...grpc.CallOption) (*ActivationHeaderList, error) {
	out := new(ActivationHeaderList)
	err := c.cc.Invoke(ctx, ActivationService_ListHeaders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivationServiceServer is the server API for ActivationService service.
// All implementations should embed UnimplementedActivationServiceServer
// for forward compatibility
type ActivationServiceServer interface {
	List(context.Context, *ActivationRequest) (*ActivationList, error)
	ListHeaders(context.Context, *ActivationRequest) (*ActivationHeaderList, error)
}

// UnimplementedActivationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedActivationServiceServer struct {
}

func (UnimplementedActivationServiceServer) List(context.Context, *ActivationRequest) (*ActivationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedActivationServiceServer) ListHeaders(context.Context, *ActivationRequest) (*ActivationHeaderList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHeaders not implemented")
}

// UnsafeActivationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivationServiceServer will
// result in compilation errors.
type UnsafeActivationServiceServer interface {
	mustEmbedUnimplementedActivationServiceServer()
}

func RegisterActivationServiceServer(s grpc.ServiceRegistrar, srv ActivationServiceServer) {
	s.RegisterService(&ActivationService_ServiceDesc, srv)
}

func _ActivationService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivationServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActivationService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivationServiceServer).List(ctx, req.(*ActivationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivationService_ListHeaders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivationServiceServer).ListHeaders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActivationService_ListHeaders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivationServiceServer).ListHeaders(ctx, req.(*ActivationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ActivationService_ServiceDesc is the grpc.ServiceDesc for ActivationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActivationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spacemesh.v2.ActivationService",
	HandlerType: (*ActivationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ActivationService_List_Handler,
		},
		{
			MethodName: "ListHeaders",
			Handler:    _ActivationService_ListHeaders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spacemesh/v2/activation.proto",
}