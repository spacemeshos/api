// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: spacemesh/v2/reward.proto

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
	RewardService_List_FullMethodName = "/spacemesh.v2.RewardService/List"
)

// RewardServiceClient is the client API for RewardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RewardServiceClient interface {
	List(ctx context.Context, in *RewardRequest, opts ...grpc.CallOption) (*RewardList, error)
}

type rewardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRewardServiceClient(cc grpc.ClientConnInterface) RewardServiceClient {
	return &rewardServiceClient{cc}
}

func (c *rewardServiceClient) List(ctx context.Context, in *RewardRequest, opts ...grpc.CallOption) (*RewardList, error) {
	out := new(RewardList)
	err := c.cc.Invoke(ctx, RewardService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RewardServiceServer is the server API for RewardService service.
// All implementations should embed UnimplementedRewardServiceServer
// for forward compatibility
type RewardServiceServer interface {
	List(context.Context, *RewardRequest) (*RewardList, error)
}

// UnimplementedRewardServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRewardServiceServer struct {
}

func (UnimplementedRewardServiceServer) List(context.Context, *RewardRequest) (*RewardList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

// UnsafeRewardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RewardServiceServer will
// result in compilation errors.
type UnsafeRewardServiceServer interface {
	mustEmbedUnimplementedRewardServiceServer()
}

func RegisterRewardServiceServer(s grpc.ServiceRegistrar, srv RewardServiceServer) {
	s.RegisterService(&RewardService_ServiceDesc, srv)
}

func _RewardService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RewardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RewardServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RewardService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RewardServiceServer).List(ctx, req.(*RewardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RewardService_ServiceDesc is the grpc.ServiceDesc for RewardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RewardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spacemesh.v2.RewardService",
	HandlerType: (*RewardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _RewardService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spacemesh/v2/reward.proto",
}