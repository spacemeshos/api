// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: spacemesh/v2beta1/smeshing.proto

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
	SmeshingService_Version_FullMethodName = "/spacemesh.v2beta1.SmeshingService/Version"
	SmeshingService_Build_FullMethodName   = "/spacemesh.v2beta1.SmeshingService/Build"
)

// SmeshingServiceClient is the client API for SmeshingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmeshingServiceClient interface {
	Version(ctx context.Context, in *SmeshingVersionRequest, opts ...grpc.CallOption) (*SmeshingVersionResponse, error)
	Build(ctx context.Context, in *SmeshingBuildRequest, opts ...grpc.CallOption) (*SmeshingBuildResponse, error)
}

type smeshingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmeshingServiceClient(cc grpc.ClientConnInterface) SmeshingServiceClient {
	return &smeshingServiceClient{cc}
}

func (c *smeshingServiceClient) Version(ctx context.Context, in *SmeshingVersionRequest, opts ...grpc.CallOption) (*SmeshingVersionResponse, error) {
	out := new(SmeshingVersionResponse)
	err := c.cc.Invoke(ctx, SmeshingService_Version_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smeshingServiceClient) Build(ctx context.Context, in *SmeshingBuildRequest, opts ...grpc.CallOption) (*SmeshingBuildResponse, error) {
	out := new(SmeshingBuildResponse)
	err := c.cc.Invoke(ctx, SmeshingService_Build_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmeshingServiceServer is the server API for SmeshingService service.
// All implementations should embed UnimplementedSmeshingServiceServer
// for forward compatibility
type SmeshingServiceServer interface {
	Version(context.Context, *SmeshingVersionRequest) (*SmeshingVersionResponse, error)
	Build(context.Context, *SmeshingBuildRequest) (*SmeshingBuildResponse, error)
}

// UnimplementedSmeshingServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSmeshingServiceServer struct {
}

func (UnimplementedSmeshingServiceServer) Version(context.Context, *SmeshingVersionRequest) (*SmeshingVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedSmeshingServiceServer) Build(context.Context, *SmeshingBuildRequest) (*SmeshingBuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Build not implemented")
}

// UnsafeSmeshingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmeshingServiceServer will
// result in compilation errors.
type UnsafeSmeshingServiceServer interface {
	mustEmbedUnimplementedSmeshingServiceServer()
}

func RegisterSmeshingServiceServer(s grpc.ServiceRegistrar, srv SmeshingServiceServer) {
	s.RegisterService(&SmeshingService_ServiceDesc, srv)
}

func _SmeshingService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmeshingVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmeshingServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmeshingService_Version_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmeshingServiceServer).Version(ctx, req.(*SmeshingVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmeshingService_Build_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmeshingBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmeshingServiceServer).Build(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmeshingService_Build_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmeshingServiceServer).Build(ctx, req.(*SmeshingBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SmeshingService_ServiceDesc is the grpc.ServiceDesc for SmeshingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmeshingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spacemesh.v2beta1.SmeshingService",
	HandlerType: (*SmeshingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _SmeshingService_Version_Handler,
		},
		{
			MethodName: "Build",
			Handler:    _SmeshingService_Build_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spacemesh/v2beta1/smeshing.proto",
}
