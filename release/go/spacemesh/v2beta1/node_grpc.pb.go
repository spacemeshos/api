// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: spacemesh/v2beta1/node.proto

package spacemeshv2beta1

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
	NodeService_Status_FullMethodName = "/spacemesh.v2beta1.NodeService/Status"
)

// NodeServiceClient is the client API for NodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeServiceClient interface {
	// Node status
	//
	// Status is a method that returns an "{{.ResponseType.Name}}".
	// This method is used to retrieve node status information.
	Status(ctx context.Context, in *NodeStatusRequest, opts ...grpc.CallOption) (*NodeStatusResponse, error)
}

type nodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeServiceClient(cc grpc.ClientConnInterface) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) Status(ctx context.Context, in *NodeStatusRequest, opts ...grpc.CallOption) (*NodeStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NodeStatusResponse)
	err := c.cc.Invoke(ctx, NodeService_Status_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServiceServer is the server API for NodeService service.
// All implementations should embed UnimplementedNodeServiceServer
// for forward compatibility.
type NodeServiceServer interface {
	// Node status
	//
	// Status is a method that returns an "{{.ResponseType.Name}}".
	// This method is used to retrieve node status information.
	Status(context.Context, *NodeStatusRequest) (*NodeStatusResponse, error)
}

// UnimplementedNodeServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNodeServiceServer struct{}

func (UnimplementedNodeServiceServer) Status(context.Context, *NodeStatusRequest) (*NodeStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedNodeServiceServer) testEmbeddedByValue() {}

// UnsafeNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServiceServer will
// result in compilation errors.
type UnsafeNodeServiceServer interface {
	mustEmbedUnimplementedNodeServiceServer()
}

func RegisterNodeServiceServer(s grpc.ServiceRegistrar, srv NodeServiceServer) {
	// If the following call pancis, it indicates UnimplementedNodeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NodeService_ServiceDesc, srv)
}

func _NodeService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_Status_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Status(ctx, req.(*NodeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeService_ServiceDesc is the grpc.ServiceDesc for NodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spacemesh.v2beta1.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _NodeService_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spacemesh/v2beta1/node.proto",
}
