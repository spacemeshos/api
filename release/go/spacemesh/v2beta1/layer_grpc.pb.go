// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: spacemesh/v2beta1/layer.proto

package spacemeshv2beta1

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
	LayerService_List_FullMethodName = "/spacemesh.v2beta1.LayerService/List"
)

// LayerServiceClient is the client API for LayerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LayerServiceClient interface {
	// List of layers
	//
	// List is a method that takes a "{{.RequestType.Name}}" body and returns an "{{.ResponseType.Name}}".
	// This method is used to retrieve a list of layers based on the provided request parameters.
	List(ctx context.Context, in *LayerRequest, opts ...grpc.CallOption) (*LayerList, error)
}

type layerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLayerServiceClient(cc grpc.ClientConnInterface) LayerServiceClient {
	return &layerServiceClient{cc}
}

func (c *layerServiceClient) List(ctx context.Context, in *LayerRequest, opts ...grpc.CallOption) (*LayerList, error) {
	out := new(LayerList)
	err := c.cc.Invoke(ctx, LayerService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LayerServiceServer is the server API for LayerService service.
// All implementations should embed UnimplementedLayerServiceServer
// for forward compatibility
type LayerServiceServer interface {
	// List of layers
	//
	// List is a method that takes a "{{.RequestType.Name}}" body and returns an "{{.ResponseType.Name}}".
	// This method is used to retrieve a list of layers based on the provided request parameters.
	List(context.Context, *LayerRequest) (*LayerList, error)
}

// UnimplementedLayerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLayerServiceServer struct {
}

func (UnimplementedLayerServiceServer) List(context.Context, *LayerRequest) (*LayerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

// UnsafeLayerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LayerServiceServer will
// result in compilation errors.
type UnsafeLayerServiceServer interface {
	mustEmbedUnimplementedLayerServiceServer()
}

func RegisterLayerServiceServer(s grpc.ServiceRegistrar, srv LayerServiceServer) {
	s.RegisterService(&LayerService_ServiceDesc, srv)
}

func _LayerService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LayerServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LayerService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LayerServiceServer).List(ctx, req.(*LayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LayerService_ServiceDesc is the grpc.ServiceDesc for LayerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LayerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spacemesh.v2beta1.LayerService",
	HandlerType: (*LayerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _LayerService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spacemesh/v2beta1/layer.proto",
}

const (
	LayerStreamService_Stream_FullMethodName = "/spacemesh.v2beta1.LayerStreamService/Stream"
)

// LayerStreamServiceClient is the client API for LayerStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LayerStreamServiceClient interface {
	Stream(ctx context.Context, in *LayerStreamRequest, opts ...grpc.CallOption) (LayerStreamService_StreamClient, error)
}

type layerStreamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLayerStreamServiceClient(cc grpc.ClientConnInterface) LayerStreamServiceClient {
	return &layerStreamServiceClient{cc}
}

func (c *layerStreamServiceClient) Stream(ctx context.Context, in *LayerStreamRequest, opts ...grpc.CallOption) (LayerStreamService_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &LayerStreamService_ServiceDesc.Streams[0], LayerStreamService_Stream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &layerStreamServiceStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LayerStreamService_StreamClient interface {
	Recv() (*Layer, error)
	grpc.ClientStream
}

type layerStreamServiceStreamClient struct {
	grpc.ClientStream
}

func (x *layerStreamServiceStreamClient) Recv() (*Layer, error) {
	m := new(Layer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LayerStreamServiceServer is the server API for LayerStreamService service.
// All implementations should embed UnimplementedLayerStreamServiceServer
// for forward compatibility
type LayerStreamServiceServer interface {
	Stream(*LayerStreamRequest, LayerStreamService_StreamServer) error
}

// UnimplementedLayerStreamServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLayerStreamServiceServer struct {
}

func (UnimplementedLayerStreamServiceServer) Stream(*LayerStreamRequest, LayerStreamService_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}

// UnsafeLayerStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LayerStreamServiceServer will
// result in compilation errors.
type UnsafeLayerStreamServiceServer interface {
	mustEmbedUnimplementedLayerStreamServiceServer()
}

func RegisterLayerStreamServiceServer(s grpc.ServiceRegistrar, srv LayerStreamServiceServer) {
	s.RegisterService(&LayerStreamService_ServiceDesc, srv)
}

func _LayerStreamService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LayerStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LayerStreamServiceServer).Stream(m, &layerStreamServiceStreamServer{stream})
}

type LayerStreamService_StreamServer interface {
	Send(*Layer) error
	grpc.ServerStream
}

type layerStreamServiceStreamServer struct {
	grpc.ServerStream
}

func (x *layerStreamServiceStreamServer) Send(m *Layer) error {
	return x.ServerStream.SendMsg(m)
}

// LayerStreamService_ServiceDesc is the grpc.ServiceDesc for LayerStreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LayerStreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spacemesh.v2beta1.LayerStreamService",
	HandlerType: (*LayerStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _LayerStreamService_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "spacemesh/v2beta1/layer.proto",
}
