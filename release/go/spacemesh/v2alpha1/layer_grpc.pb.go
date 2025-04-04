// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: spacemesh/v2alpha1/layer.proto

package spacemeshv2alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LayerService_List_FullMethodName = "/spacemesh.v2alpha1.LayerService/List"
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
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LayerList)
	err := c.cc.Invoke(ctx, LayerService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LayerServiceServer is the server API for LayerService service.
// All implementations should embed UnimplementedLayerServiceServer
// for forward compatibility.
type LayerServiceServer interface {
	// List of layers
	//
	// List is a method that takes a "{{.RequestType.Name}}" body and returns an "{{.ResponseType.Name}}".
	// This method is used to retrieve a list of layers based on the provided request parameters.
	List(context.Context, *LayerRequest) (*LayerList, error)
}

// UnimplementedLayerServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLayerServiceServer struct{}

func (UnimplementedLayerServiceServer) List(context.Context, *LayerRequest) (*LayerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedLayerServiceServer) testEmbeddedByValue() {}

// UnsafeLayerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LayerServiceServer will
// result in compilation errors.
type UnsafeLayerServiceServer interface {
	mustEmbedUnimplementedLayerServiceServer()
}

func RegisterLayerServiceServer(s grpc.ServiceRegistrar, srv LayerServiceServer) {
	// If the following call pancis, it indicates UnimplementedLayerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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
	ServiceName: "spacemesh.v2alpha1.LayerService",
	HandlerType: (*LayerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _LayerService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spacemesh/v2alpha1/layer.proto",
}

const (
	LayerStreamService_Stream_FullMethodName = "/spacemesh.v2alpha1.LayerStreamService/Stream"
)

// LayerStreamServiceClient is the client API for LayerStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LayerStreamServiceClient interface {
	Stream(ctx context.Context, in *LayerStreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Layer], error)
}

type layerStreamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLayerStreamServiceClient(cc grpc.ClientConnInterface) LayerStreamServiceClient {
	return &layerStreamServiceClient{cc}
}

func (c *layerStreamServiceClient) Stream(ctx context.Context, in *LayerStreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Layer], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &LayerStreamService_ServiceDesc.Streams[0], LayerStreamService_Stream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[LayerStreamRequest, Layer]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LayerStreamService_StreamClient = grpc.ServerStreamingClient[Layer]

// LayerStreamServiceServer is the server API for LayerStreamService service.
// All implementations should embed UnimplementedLayerStreamServiceServer
// for forward compatibility.
type LayerStreamServiceServer interface {
	Stream(*LayerStreamRequest, grpc.ServerStreamingServer[Layer]) error
}

// UnimplementedLayerStreamServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLayerStreamServiceServer struct{}

func (UnimplementedLayerStreamServiceServer) Stream(*LayerStreamRequest, grpc.ServerStreamingServer[Layer]) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (UnimplementedLayerStreamServiceServer) testEmbeddedByValue() {}

// UnsafeLayerStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LayerStreamServiceServer will
// result in compilation errors.
type UnsafeLayerStreamServiceServer interface {
	mustEmbedUnimplementedLayerStreamServiceServer()
}

func RegisterLayerStreamServiceServer(s grpc.ServiceRegistrar, srv LayerStreamServiceServer) {
	// If the following call pancis, it indicates UnimplementedLayerStreamServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LayerStreamService_ServiceDesc, srv)
}

func _LayerStreamService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LayerStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LayerStreamServiceServer).Stream(m, &grpc.GenericServerStream[LayerStreamRequest, Layer]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LayerStreamService_StreamServer = grpc.ServerStreamingServer[Layer]

// LayerStreamService_ServiceDesc is the grpc.ServiceDesc for LayerStreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LayerStreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spacemesh.v2alpha1.LayerStreamService",
	HandlerType: (*LayerStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _LayerStreamService_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "spacemesh/v2alpha1/layer.proto",
}
