// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: spacemesh/v2beta1/smesher.proto

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
	SmesherService_Version_FullMethodName = "/spacemesh.v2beta1.SmesherService/Version"
	SmesherService_Build_FullMethodName   = "/spacemesh.v2beta1.SmesherService/Build"
)

// SmesherServiceClient is the client API for SmesherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmesherServiceClient interface {
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	Build(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (*BuildResponse, error)
}

type smesherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmesherServiceClient(cc grpc.ClientConnInterface) SmesherServiceClient {
	return &smesherServiceClient{cc}
}

func (c *smesherServiceClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, SmesherService_Version_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smesherServiceClient) Build(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (*BuildResponse, error) {
	out := new(BuildResponse)
	err := c.cc.Invoke(ctx, SmesherService_Build_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmesherServiceServer is the server API for SmesherService service.
// All implementations should embed UnimplementedSmesherServiceServer
// for forward compatibility
type SmesherServiceServer interface {
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
	Build(context.Context, *BuildRequest) (*BuildResponse, error)
}

// UnimplementedSmesherServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSmesherServiceServer struct {
}

func (UnimplementedSmesherServiceServer) Version(context.Context, *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedSmesherServiceServer) Build(context.Context, *BuildRequest) (*BuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Build not implemented")
}

// UnsafeSmesherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmesherServiceServer will
// result in compilation errors.
type UnsafeSmesherServiceServer interface {
	mustEmbedUnimplementedSmesherServiceServer()
}

func RegisterSmesherServiceServer(s grpc.ServiceRegistrar, srv SmesherServiceServer) {
	s.RegisterService(&SmesherService_ServiceDesc, srv)
}

func _SmesherService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmesherServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmesherService_Version_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmesherServiceServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmesherService_Build_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmesherServiceServer).Build(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmesherService_Build_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmesherServiceServer).Build(ctx, req.(*BuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SmesherService_ServiceDesc is the grpc.ServiceDesc for SmesherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmesherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spacemesh.v2beta1.SmesherService",
	HandlerType: (*SmesherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _SmesherService_Version_Handler,
		},
		{
			MethodName: "Build",
			Handler:    _SmesherService_Build_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spacemesh/v2beta1/smesher.proto",
}
