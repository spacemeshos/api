// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: spacemesh/v2alpha1/malfeasance.proto

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
	MalfeasanceService_List_FullMethodName = "/spacemesh.v2alpha1.MalfeasanceService/List"
)

// MalfeasanceServiceClient is the client API for MalfeasanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MalfeasanceServiceClient interface {
	// List of malfeasance proofs
	//
	// List is a method that takes a "{{.RequestType.Name}}" body and returns an "{{.ResponseType.Name}}".
	// This method is used to retrieve a list of malfeasance proofs based on the provided request parameters.
	List(ctx context.Context, in *MalfeasanceRequest, opts ...grpc.CallOption) (*MalfeasanceList, error)
}

type malfeasanceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMalfeasanceServiceClient(cc grpc.ClientConnInterface) MalfeasanceServiceClient {
	return &malfeasanceServiceClient{cc}
}

func (c *malfeasanceServiceClient) List(ctx context.Context, in *MalfeasanceRequest, opts ...grpc.CallOption) (*MalfeasanceList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MalfeasanceList)
	err := c.cc.Invoke(ctx, MalfeasanceService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MalfeasanceServiceServer is the server API for MalfeasanceService service.
// All implementations should embed UnimplementedMalfeasanceServiceServer
// for forward compatibility.
type MalfeasanceServiceServer interface {
	// List of malfeasance proofs
	//
	// List is a method that takes a "{{.RequestType.Name}}" body and returns an "{{.ResponseType.Name}}".
	// This method is used to retrieve a list of malfeasance proofs based on the provided request parameters.
	List(context.Context, *MalfeasanceRequest) (*MalfeasanceList, error)
}

// UnimplementedMalfeasanceServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMalfeasanceServiceServer struct{}

func (UnimplementedMalfeasanceServiceServer) List(context.Context, *MalfeasanceRequest) (*MalfeasanceList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedMalfeasanceServiceServer) testEmbeddedByValue() {}

// UnsafeMalfeasanceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MalfeasanceServiceServer will
// result in compilation errors.
type UnsafeMalfeasanceServiceServer interface {
	mustEmbedUnimplementedMalfeasanceServiceServer()
}

func RegisterMalfeasanceServiceServer(s grpc.ServiceRegistrar, srv MalfeasanceServiceServer) {
	// If the following call pancis, it indicates UnimplementedMalfeasanceServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MalfeasanceService_ServiceDesc, srv)
}

func _MalfeasanceService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MalfeasanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MalfeasanceServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MalfeasanceService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MalfeasanceServiceServer).List(ctx, req.(*MalfeasanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MalfeasanceService_ServiceDesc is the grpc.ServiceDesc for MalfeasanceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MalfeasanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spacemesh.v2alpha1.MalfeasanceService",
	HandlerType: (*MalfeasanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _MalfeasanceService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spacemesh/v2alpha1/malfeasance.proto",
}

const (
	MalfeasanceStreamService_Stream_FullMethodName = "/spacemesh.v2alpha1.MalfeasanceStreamService/Stream"
)

// MalfeasanceStreamServiceClient is the client API for MalfeasanceStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MalfeasanceStreamServiceClient interface {
	Stream(ctx context.Context, in *MalfeasanceStreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[MalfeasanceProof], error)
}

type malfeasanceStreamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMalfeasanceStreamServiceClient(cc grpc.ClientConnInterface) MalfeasanceStreamServiceClient {
	return &malfeasanceStreamServiceClient{cc}
}

func (c *malfeasanceStreamServiceClient) Stream(ctx context.Context, in *MalfeasanceStreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[MalfeasanceProof], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MalfeasanceStreamService_ServiceDesc.Streams[0], MalfeasanceStreamService_Stream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[MalfeasanceStreamRequest, MalfeasanceProof]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MalfeasanceStreamService_StreamClient = grpc.ServerStreamingClient[MalfeasanceProof]

// MalfeasanceStreamServiceServer is the server API for MalfeasanceStreamService service.
// All implementations should embed UnimplementedMalfeasanceStreamServiceServer
// for forward compatibility.
type MalfeasanceStreamServiceServer interface {
	Stream(*MalfeasanceStreamRequest, grpc.ServerStreamingServer[MalfeasanceProof]) error
}

// UnimplementedMalfeasanceStreamServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMalfeasanceStreamServiceServer struct{}

func (UnimplementedMalfeasanceStreamServiceServer) Stream(*MalfeasanceStreamRequest, grpc.ServerStreamingServer[MalfeasanceProof]) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (UnimplementedMalfeasanceStreamServiceServer) testEmbeddedByValue() {}

// UnsafeMalfeasanceStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MalfeasanceStreamServiceServer will
// result in compilation errors.
type UnsafeMalfeasanceStreamServiceServer interface {
	mustEmbedUnimplementedMalfeasanceStreamServiceServer()
}

func RegisterMalfeasanceStreamServiceServer(s grpc.ServiceRegistrar, srv MalfeasanceStreamServiceServer) {
	// If the following call pancis, it indicates UnimplementedMalfeasanceStreamServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MalfeasanceStreamService_ServiceDesc, srv)
}

func _MalfeasanceStreamService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MalfeasanceStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MalfeasanceStreamServiceServer).Stream(m, &grpc.GenericServerStream[MalfeasanceStreamRequest, MalfeasanceProof]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MalfeasanceStreamService_StreamServer = grpc.ServerStreamingServer[MalfeasanceProof]

// MalfeasanceStreamService_ServiceDesc is the grpc.ServiceDesc for MalfeasanceStreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MalfeasanceStreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spacemesh.v2alpha1.MalfeasanceStreamService",
	HandlerType: (*MalfeasanceStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _MalfeasanceStreamService_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "spacemesh/v2alpha1/malfeasance.proto",
}
