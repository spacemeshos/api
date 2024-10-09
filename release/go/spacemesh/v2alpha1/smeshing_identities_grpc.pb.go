// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: spacemesh/v2alpha1/smeshing_identities.proto

package spacemeshv2alpha1

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
	SmeshingIdentitiesService_PoetServices_FullMethodName = "/spacemesh.v2alpha1.SmeshingIdentitiesService/PoetServices"
)

// SmeshingIdentitiesServiceClient is the client API for SmeshingIdentitiesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmeshingIdentitiesServiceClient interface {
	// Returns set of configured poet addresses and poets addresses from registrations, if given,
	// and warning in case, if there are registrations with poets, which are not in configured poets set.
	PoetServices(ctx context.Context, in *PoetServicesRequest, opts ...grpc.CallOption) (*PoetServicesResponse, error)
}

type smeshingIdentitiesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmeshingIdentitiesServiceClient(cc grpc.ClientConnInterface) SmeshingIdentitiesServiceClient {
	return &smeshingIdentitiesServiceClient{cc}
}

func (c *smeshingIdentitiesServiceClient) PoetServices(ctx context.Context, in *PoetServicesRequest, opts ...grpc.CallOption) (*PoetServicesResponse, error) {
	out := new(PoetServicesResponse)
	err := c.cc.Invoke(ctx, SmeshingIdentitiesService_PoetServices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmeshingIdentitiesServiceServer is the server API for SmeshingIdentitiesService service.
// All implementations should embed UnimplementedSmeshingIdentitiesServiceServer
// for forward compatibility
type SmeshingIdentitiesServiceServer interface {
	// Returns set of configured poet addresses and poets addresses from registrations, if given,
	// and warning in case, if there are registrations with poets, which are not in configured poets set.
	PoetServices(context.Context, *PoetServicesRequest) (*PoetServicesResponse, error)
}

// UnimplementedSmeshingIdentitiesServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSmeshingIdentitiesServiceServer struct {
}

func (UnimplementedSmeshingIdentitiesServiceServer) PoetServices(context.Context, *PoetServicesRequest) (*PoetServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PoetServices not implemented")
}

// UnsafeSmeshingIdentitiesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmeshingIdentitiesServiceServer will
// result in compilation errors.
type UnsafeSmeshingIdentitiesServiceServer interface {
	mustEmbedUnimplementedSmeshingIdentitiesServiceServer()
}

func RegisterSmeshingIdentitiesServiceServer(s grpc.ServiceRegistrar, srv SmeshingIdentitiesServiceServer) {
	s.RegisterService(&SmeshingIdentitiesService_ServiceDesc, srv)
}

func _SmeshingIdentitiesService_PoetServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PoetServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmeshingIdentitiesServiceServer).PoetServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmeshingIdentitiesService_PoetServices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmeshingIdentitiesServiceServer).PoetServices(ctx, req.(*PoetServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SmeshingIdentitiesService_ServiceDesc is the grpc.ServiceDesc for SmeshingIdentitiesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmeshingIdentitiesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spacemesh.v2alpha1.SmeshingIdentitiesService",
	HandlerType: (*SmeshingIdentitiesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PoetServices",
			Handler:    _SmeshingIdentitiesService_PoetServices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spacemesh/v2alpha1/smeshing_identities.proto",
}