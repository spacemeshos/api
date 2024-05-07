// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: spacemesh/v2alpha1/tx.proto

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
	TransactionStreamService_Stream_FullMethodName = "/spacemesh.v2alpha1.TransactionStreamService/Stream"
)

// TransactionStreamServiceClient is the client API for TransactionStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionStreamServiceClient interface {
	Stream(ctx context.Context, in *TransactionStreamRequest, opts ...grpc.CallOption) (TransactionStreamService_StreamClient, error)
}

type transactionStreamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionStreamServiceClient(cc grpc.ClientConnInterface) TransactionStreamServiceClient {
	return &transactionStreamServiceClient{cc}
}

func (c *transactionStreamServiceClient) Stream(ctx context.Context, in *TransactionStreamRequest, opts ...grpc.CallOption) (TransactionStreamService_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &TransactionStreamService_ServiceDesc.Streams[0], TransactionStreamService_Stream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &transactionStreamServiceStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TransactionStreamService_StreamClient interface {
	Recv() (*Transaction, error)
	grpc.ClientStream
}

type transactionStreamServiceStreamClient struct {
	grpc.ClientStream
}

func (x *transactionStreamServiceStreamClient) Recv() (*Transaction, error) {
	m := new(Transaction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TransactionStreamServiceServer is the server API for TransactionStreamService service.
// All implementations should embed UnimplementedTransactionStreamServiceServer
// for forward compatibility
type TransactionStreamServiceServer interface {
	Stream(*TransactionStreamRequest, TransactionStreamService_StreamServer) error
}

// UnimplementedTransactionStreamServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTransactionStreamServiceServer struct {
}

func (UnimplementedTransactionStreamServiceServer) Stream(*TransactionStreamRequest, TransactionStreamService_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}

// UnsafeTransactionStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionStreamServiceServer will
// result in compilation errors.
type UnsafeTransactionStreamServiceServer interface {
	mustEmbedUnimplementedTransactionStreamServiceServer()
}

func RegisterTransactionStreamServiceServer(s grpc.ServiceRegistrar, srv TransactionStreamServiceServer) {
	s.RegisterService(&TransactionStreamService_ServiceDesc, srv)
}

func _TransactionStreamService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TransactionStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransactionStreamServiceServer).Stream(m, &transactionStreamServiceStreamServer{stream})
}

type TransactionStreamService_StreamServer interface {
	Send(*Transaction) error
	grpc.ServerStream
}

type transactionStreamServiceStreamServer struct {
	grpc.ServerStream
}

func (x *transactionStreamServiceStreamServer) Send(m *Transaction) error {
	return x.ServerStream.SendMsg(m)
}

// TransactionStreamService_ServiceDesc is the grpc.ServiceDesc for TransactionStreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionStreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spacemesh.v2alpha1.TransactionStreamService",
	HandlerType: (*TransactionStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _TransactionStreamService_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "spacemesh/v2alpha1/tx.proto",
}

const (
	TransactionService_List_FullMethodName              = "/spacemesh.v2alpha1.TransactionService/List"
	TransactionService_ParseTransaction_FullMethodName  = "/spacemesh.v2alpha1.TransactionService/ParseTransaction"
	TransactionService_SubmitTransaction_FullMethodName = "/spacemesh.v2alpha1.TransactionService/SubmitTransaction"
	TransactionService_EstimateGas_FullMethodName       = "/spacemesh.v2alpha1.TransactionService/EstimateGas"
)

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionServiceClient interface {
	List(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionList, error)
	ParseTransaction(ctx context.Context, in *ParseTransactionRequest, opts ...grpc.CallOption) (*ParseTransactionResponse, error)
	SubmitTransaction(ctx context.Context, in *SubmitTransactionRequest, opts ...grpc.CallOption) (*SubmitTransactionResponse, error)
	EstimateGas(ctx context.Context, in *EstimateGasRequest, opts ...grpc.CallOption) (*EstimateGasResponse, error)
}

type transactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) List(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionList, error) {
	out := new(TransactionList)
	err := c.cc.Invoke(ctx, TransactionService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) ParseTransaction(ctx context.Context, in *ParseTransactionRequest, opts ...grpc.CallOption) (*ParseTransactionResponse, error) {
	out := new(ParseTransactionResponse)
	err := c.cc.Invoke(ctx, TransactionService_ParseTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) SubmitTransaction(ctx context.Context, in *SubmitTransactionRequest, opts ...grpc.CallOption) (*SubmitTransactionResponse, error) {
	out := new(SubmitTransactionResponse)
	err := c.cc.Invoke(ctx, TransactionService_SubmitTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) EstimateGas(ctx context.Context, in *EstimateGasRequest, opts ...grpc.CallOption) (*EstimateGasResponse, error) {
	out := new(EstimateGasResponse)
	err := c.cc.Invoke(ctx, TransactionService_EstimateGas_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
// All implementations should embed UnimplementedTransactionServiceServer
// for forward compatibility
type TransactionServiceServer interface {
	List(context.Context, *TransactionRequest) (*TransactionList, error)
	ParseTransaction(context.Context, *ParseTransactionRequest) (*ParseTransactionResponse, error)
	SubmitTransaction(context.Context, *SubmitTransactionRequest) (*SubmitTransactionResponse, error)
	EstimateGas(context.Context, *EstimateGasRequest) (*EstimateGasResponse, error)
}

// UnimplementedTransactionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTransactionServiceServer struct {
}

func (UnimplementedTransactionServiceServer) List(context.Context, *TransactionRequest) (*TransactionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTransactionServiceServer) ParseTransaction(context.Context, *ParseTransactionRequest) (*ParseTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) SubmitTransaction(context.Context, *SubmitTransactionRequest) (*SubmitTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) EstimateGas(context.Context, *EstimateGasRequest) (*EstimateGasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstimateGas not implemented")
}

// UnsafeTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServiceServer will
// result in compilation errors.
type UnsafeTransactionServiceServer interface {
	mustEmbedUnimplementedTransactionServiceServer()
}

func RegisterTransactionServiceServer(s grpc.ServiceRegistrar, srv TransactionServiceServer) {
	s.RegisterService(&TransactionService_ServiceDesc, srv)
}

func _TransactionService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).List(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_ParseTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).ParseTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_ParseTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).ParseTransaction(ctx, req.(*ParseTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_SubmitTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).SubmitTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_SubmitTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).SubmitTransaction(ctx, req.(*SubmitTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_EstimateGas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstimateGasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).EstimateGas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_EstimateGas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).EstimateGas(ctx, req.(*EstimateGasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionService_ServiceDesc is the grpc.ServiceDesc for TransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spacemesh.v2alpha1.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _TransactionService_List_Handler,
		},
		{
			MethodName: "ParseTransaction",
			Handler:    _TransactionService_ParseTransaction_Handler,
		},
		{
			MethodName: "SubmitTransaction",
			Handler:    _TransactionService_SubmitTransaction_Handler,
		},
		{
			MethodName: "EstimateGas",
			Handler:    _TransactionService_EstimateGas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spacemesh/v2alpha1/tx.proto",
}
