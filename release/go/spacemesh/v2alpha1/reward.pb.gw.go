// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: spacemesh/v2alpha1/reward.proto

/*
Package spacemeshv2alpha1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package spacemeshv2alpha1

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_RewardService_List_0(ctx context.Context, marshaler runtime.Marshaler, client RewardServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq RewardRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.List(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_RewardService_List_0(ctx context.Context, marshaler runtime.Marshaler, server RewardServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq RewardRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.List(ctx, &protoReq)
	return msg, metadata, err

}

func request_RewardStreamService_Stream_0(ctx context.Context, marshaler runtime.Marshaler, client RewardStreamServiceClient, req *http.Request, pathParams map[string]string) (RewardStreamService_StreamClient, runtime.ServerMetadata, error) {
	var protoReq RewardStreamRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.Stream(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

// RegisterRewardServiceHandlerServer registers the http handlers for service RewardService to "mux".
// UnaryRPC     :call RewardServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterRewardServiceHandlerFromEndpoint instead.
func RegisterRewardServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server RewardServiceServer) error {

	mux.Handle("POST", pattern_RewardService_List_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/spacemesh.v2alpha1.RewardService/List", runtime.WithHTTPPathPattern("/spacemesh.v2alpha1.RewardService/List"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_RewardService_List_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_RewardService_List_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterRewardStreamServiceHandlerServer registers the http handlers for service RewardStreamService to "mux".
// UnaryRPC     :call RewardStreamServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterRewardStreamServiceHandlerFromEndpoint instead.
func RegisterRewardStreamServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server RewardStreamServiceServer) error {

	mux.Handle("POST", pattern_RewardStreamService_Stream_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterRewardServiceHandlerFromEndpoint is same as RegisterRewardServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterRewardServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterRewardServiceHandler(ctx, mux, conn)
}

// RegisterRewardServiceHandler registers the http handlers for service RewardService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterRewardServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterRewardServiceHandlerClient(ctx, mux, NewRewardServiceClient(conn))
}

// RegisterRewardServiceHandlerClient registers the http handlers for service RewardService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "RewardServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "RewardServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "RewardServiceClient" to call the correct interceptors.
func RegisterRewardServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client RewardServiceClient) error {

	mux.Handle("POST", pattern_RewardService_List_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/spacemesh.v2alpha1.RewardService/List", runtime.WithHTTPPathPattern("/spacemesh.v2alpha1.RewardService/List"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_RewardService_List_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_RewardService_List_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_RewardService_List_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"spacemesh.v2alpha1.RewardService", "List"}, ""))
)

var (
	forward_RewardService_List_0 = runtime.ForwardResponseMessage
)

// RegisterRewardStreamServiceHandlerFromEndpoint is same as RegisterRewardStreamServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterRewardStreamServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterRewardStreamServiceHandler(ctx, mux, conn)
}

// RegisterRewardStreamServiceHandler registers the http handlers for service RewardStreamService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterRewardStreamServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterRewardStreamServiceHandlerClient(ctx, mux, NewRewardStreamServiceClient(conn))
}

// RegisterRewardStreamServiceHandlerClient registers the http handlers for service RewardStreamService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "RewardStreamServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "RewardStreamServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "RewardStreamServiceClient" to call the correct interceptors.
func RegisterRewardStreamServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client RewardStreamServiceClient) error {

	mux.Handle("POST", pattern_RewardStreamService_Stream_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/spacemesh.v2alpha1.RewardStreamService/Stream", runtime.WithHTTPPathPattern("/spacemesh.v2alpha1.RewardStreamService/Stream"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_RewardStreamService_Stream_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_RewardStreamService_Stream_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_RewardStreamService_Stream_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"spacemesh.v2alpha1.RewardStreamService", "Stream"}, ""))
)

var (
	forward_RewardStreamService_Stream_0 = runtime.ForwardResponseStream
)
