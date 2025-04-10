// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: spacemesh/v1/tx.proto

package spacemeshv1

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
	TransactionService_SubmitTransaction_FullMethodName       = "/spacemesh.v1.TransactionService/SubmitTransaction"
	TransactionService_ParseTransaction_FullMethodName        = "/spacemesh.v1.TransactionService/ParseTransaction"
	TransactionService_TransactionsState_FullMethodName       = "/spacemesh.v1.TransactionService/TransactionsState"
	TransactionService_TransactionsStateStream_FullMethodName = "/spacemesh.v1.TransactionService/TransactionsStateStream"
	TransactionService_StreamResults_FullMethodName           = "/spacemesh.v1.TransactionService/StreamResults"
)

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Provides clients a way to submit a tx to the network for processing, and to
// check or follow the "journey" of a tx from mempool to block inclusion to
// mesh to STF processing. This service is separate from the Mesh and
// GlobalState services because txs move across both.
type TransactionServiceClient interface {
	// Submit a new tx to the node for processing. The response
	// TransactionState message contains both the txid of the new tx, as well
	// as whether or not it was admitted into the mempool.
	SubmitTransaction(ctx context.Context, in *SubmitTransactionRequest, opts ...grpc.CallOption) (*SubmitTransactionResponse, error)
	// ParseTransaction without submitting it to the mempool.
	// It has a limitation that it will work only for an already spawned accounts,
	// and for initial spawn transactions. Client is expected to wait until spawn transaction
	// executes before it will submit other transactions.
	ParseTransaction(ctx context.Context, in *ParseTransactionRequest, opts ...grpc.CallOption) (*ParseTransactionResponse, error)
	// Returns current tx state for one or more txs which indicates if a tx is
	// on the mesh, on its way to the mesh or was rejected and will never get
	// to the mesh
	TransactionsState(ctx context.Context, in *TransactionsStateRequest, opts ...grpc.CallOption) (*TransactionsStateResponse, error)
	// Returns tx state for one or more txs every time the tx state changes for
	// one of these txs
	TransactionsStateStream(ctx context.Context, in *TransactionsStateStreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TransactionsStateStreamResponse], error)
	// StreamResults streams historical data and watch live events with transaction results.
	StreamResults(ctx context.Context, in *TransactionResultsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TransactionResult], error)
}

type transactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) SubmitTransaction(ctx context.Context, in *SubmitTransactionRequest, opts ...grpc.CallOption) (*SubmitTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubmitTransactionResponse)
	err := c.cc.Invoke(ctx, TransactionService_SubmitTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) ParseTransaction(ctx context.Context, in *ParseTransactionRequest, opts ...grpc.CallOption) (*ParseTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ParseTransactionResponse)
	err := c.cc.Invoke(ctx, TransactionService_ParseTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) TransactionsState(ctx context.Context, in *TransactionsStateRequest, opts ...grpc.CallOption) (*TransactionsStateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransactionsStateResponse)
	err := c.cc.Invoke(ctx, TransactionService_TransactionsState_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) TransactionsStateStream(ctx context.Context, in *TransactionsStateStreamRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TransactionsStateStreamResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TransactionService_ServiceDesc.Streams[0], TransactionService_TransactionsStateStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[TransactionsStateStreamRequest, TransactionsStateStreamResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TransactionService_TransactionsStateStreamClient = grpc.ServerStreamingClient[TransactionsStateStreamResponse]

func (c *transactionServiceClient) StreamResults(ctx context.Context, in *TransactionResultsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TransactionResult], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TransactionService_ServiceDesc.Streams[1], TransactionService_StreamResults_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[TransactionResultsRequest, TransactionResult]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TransactionService_StreamResultsClient = grpc.ServerStreamingClient[TransactionResult]

// TransactionServiceServer is the server API for TransactionService service.
// All implementations should embed UnimplementedTransactionServiceServer
// for forward compatibility.
//
// Provides clients a way to submit a tx to the network for processing, and to
// check or follow the "journey" of a tx from mempool to block inclusion to
// mesh to STF processing. This service is separate from the Mesh and
// GlobalState services because txs move across both.
type TransactionServiceServer interface {
	// Submit a new tx to the node for processing. The response
	// TransactionState message contains both the txid of the new tx, as well
	// as whether or not it was admitted into the mempool.
	SubmitTransaction(context.Context, *SubmitTransactionRequest) (*SubmitTransactionResponse, error)
	// ParseTransaction without submitting it to the mempool.
	// It has a limitation that it will work only for an already spawned accounts,
	// and for initial spawn transactions. Client is expected to wait until spawn transaction
	// executes before it will submit other transactions.
	ParseTransaction(context.Context, *ParseTransactionRequest) (*ParseTransactionResponse, error)
	// Returns current tx state for one or more txs which indicates if a tx is
	// on the mesh, on its way to the mesh or was rejected and will never get
	// to the mesh
	TransactionsState(context.Context, *TransactionsStateRequest) (*TransactionsStateResponse, error)
	// Returns tx state for one or more txs every time the tx state changes for
	// one of these txs
	TransactionsStateStream(*TransactionsStateStreamRequest, grpc.ServerStreamingServer[TransactionsStateStreamResponse]) error
	// StreamResults streams historical data and watch live events with transaction results.
	StreamResults(*TransactionResultsRequest, grpc.ServerStreamingServer[TransactionResult]) error
}

// UnimplementedTransactionServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTransactionServiceServer struct{}

func (UnimplementedTransactionServiceServer) SubmitTransaction(context.Context, *SubmitTransactionRequest) (*SubmitTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) ParseTransaction(context.Context, *ParseTransactionRequest) (*ParseTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) TransactionsState(context.Context, *TransactionsStateRequest) (*TransactionsStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransactionsState not implemented")
}
func (UnimplementedTransactionServiceServer) TransactionsStateStream(*TransactionsStateStreamRequest, grpc.ServerStreamingServer[TransactionsStateStreamResponse]) error {
	return status.Errorf(codes.Unimplemented, "method TransactionsStateStream not implemented")
}
func (UnimplementedTransactionServiceServer) StreamResults(*TransactionResultsRequest, grpc.ServerStreamingServer[TransactionResult]) error {
	return status.Errorf(codes.Unimplemented, "method StreamResults not implemented")
}
func (UnimplementedTransactionServiceServer) testEmbeddedByValue() {}

// UnsafeTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServiceServer will
// result in compilation errors.
type UnsafeTransactionServiceServer interface {
	mustEmbedUnimplementedTransactionServiceServer()
}

func RegisterTransactionServiceServer(s grpc.ServiceRegistrar, srv TransactionServiceServer) {
	// If the following call pancis, it indicates UnimplementedTransactionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TransactionService_ServiceDesc, srv)
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

func _TransactionService_TransactionsState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionsStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).TransactionsState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_TransactionsState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).TransactionsState(ctx, req.(*TransactionsStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_TransactionsStateStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TransactionsStateStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransactionServiceServer).TransactionsStateStream(m, &grpc.GenericServerStream[TransactionsStateStreamRequest, TransactionsStateStreamResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TransactionService_TransactionsStateStreamServer = grpc.ServerStreamingServer[TransactionsStateStreamResponse]

func _TransactionService_StreamResults_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TransactionResultsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransactionServiceServer).StreamResults(m, &grpc.GenericServerStream[TransactionResultsRequest, TransactionResult]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TransactionService_StreamResultsServer = grpc.ServerStreamingServer[TransactionResult]

// TransactionService_ServiceDesc is the grpc.ServiceDesc for TransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spacemesh.v1.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitTransaction",
			Handler:    _TransactionService_SubmitTransaction_Handler,
		},
		{
			MethodName: "ParseTransaction",
			Handler:    _TransactionService_ParseTransaction_Handler,
		},
		{
			MethodName: "TransactionsState",
			Handler:    _TransactionService_TransactionsState_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TransactionsStateStream",
			Handler:       _TransactionService_TransactionsStateStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamResults",
			Handler:       _TransactionService_StreamResults_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "spacemesh/v1/tx.proto",
}
