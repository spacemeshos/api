// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: spacemesh/v2beta1/activation.proto

/*
Package spacemeshv2beta1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package spacemeshv2beta1

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

func request_ActivationStreamService_Stream_0(ctx context.Context, marshaler runtime.Marshaler, client ActivationStreamServiceClient, req *http.Request, pathParams map[string]string) (ActivationStreamService_StreamClient, runtime.ServerMetadata, error) {
	var protoReq ActivationStreamRequest
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

var (
	filter_ActivationService_List_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_ActivationService_List_0(ctx context.Context, marshaler runtime.Marshaler, client ActivationServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ActivationRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_ActivationService_List_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.List(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ActivationService_List_0(ctx context.Context, marshaler runtime.Marshaler, server ActivationServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ActivationRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_ActivationService_List_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.List(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_ActivationService_ActivationsCount_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_ActivationService_ActivationsCount_0(ctx context.Context, marshaler runtime.Marshaler, client ActivationServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ActivationsCountRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_ActivationService_ActivationsCount_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ActivationsCount(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ActivationService_ActivationsCount_0(ctx context.Context, marshaler runtime.Marshaler, server ActivationServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ActivationsCountRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_ActivationService_ActivationsCount_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.ActivationsCount(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterActivationStreamServiceHandlerServer registers the http handlers for service ActivationStreamService to "mux".
// UnaryRPC     :call ActivationStreamServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterActivationStreamServiceHandlerFromEndpoint instead.
func RegisterActivationStreamServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server ActivationStreamServiceServer) error {

	mux.Handle("POST", pattern_ActivationStreamService_Stream_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterActivationServiceHandlerServer registers the http handlers for service ActivationService to "mux".
// UnaryRPC     :call ActivationServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterActivationServiceHandlerFromEndpoint instead.
func RegisterActivationServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server ActivationServiceServer) error {

	mux.Handle("GET", pattern_ActivationService_List_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/spacemesh.v2beta1.ActivationService/List", runtime.WithHTTPPathPattern("/spacemesh.v2beta1.ActivationService/List"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ActivationService_List_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ActivationService_List_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ActivationService_ActivationsCount_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/spacemesh.v2beta1.ActivationService/ActivationsCount", runtime.WithHTTPPathPattern("/spacemesh.v2beta1.ActivationService/ActivationsCount"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ActivationService_ActivationsCount_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ActivationService_ActivationsCount_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterActivationStreamServiceHandlerFromEndpoint is same as RegisterActivationStreamServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterActivationStreamServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterActivationStreamServiceHandler(ctx, mux, conn)
}

// RegisterActivationStreamServiceHandler registers the http handlers for service ActivationStreamService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterActivationStreamServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterActivationStreamServiceHandlerClient(ctx, mux, NewActivationStreamServiceClient(conn))
}

// RegisterActivationStreamServiceHandlerClient registers the http handlers for service ActivationStreamService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "ActivationStreamServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "ActivationStreamServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "ActivationStreamServiceClient" to call the correct interceptors.
func RegisterActivationStreamServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ActivationStreamServiceClient) error {

	mux.Handle("POST", pattern_ActivationStreamService_Stream_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/spacemesh.v2beta1.ActivationStreamService/Stream", runtime.WithHTTPPathPattern("/spacemesh.v2beta1.ActivationStreamService/Stream"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ActivationStreamService_Stream_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ActivationStreamService_Stream_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_ActivationStreamService_Stream_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"spacemesh.v2beta1.ActivationStreamService", "Stream"}, ""))
)

var (
	forward_ActivationStreamService_Stream_0 = runtime.ForwardResponseStream
)

// RegisterActivationServiceHandlerFromEndpoint is same as RegisterActivationServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterActivationServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterActivationServiceHandler(ctx, mux, conn)
}

// RegisterActivationServiceHandler registers the http handlers for service ActivationService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterActivationServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterActivationServiceHandlerClient(ctx, mux, NewActivationServiceClient(conn))
}

// RegisterActivationServiceHandlerClient registers the http handlers for service ActivationService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "ActivationServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "ActivationServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "ActivationServiceClient" to call the correct interceptors.
func RegisterActivationServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ActivationServiceClient) error {

	mux.Handle("GET", pattern_ActivationService_List_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/spacemesh.v2beta1.ActivationService/List", runtime.WithHTTPPathPattern("/spacemesh.v2beta1.ActivationService/List"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ActivationService_List_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ActivationService_List_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ActivationService_ActivationsCount_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/spacemesh.v2beta1.ActivationService/ActivationsCount", runtime.WithHTTPPathPattern("/spacemesh.v2beta1.ActivationService/ActivationsCount"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ActivationService_ActivationsCount_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ActivationService_ActivationsCount_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_ActivationService_List_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"spacemesh.v2beta1.ActivationService", "List"}, ""))

	pattern_ActivationService_ActivationsCount_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"spacemesh.v2beta1.ActivationService", "ActivationsCount"}, ""))
)

var (
	forward_ActivationService_List_0 = runtime.ForwardResponseMessage

	forward_ActivationService_ActivationsCount_0 = runtime.ForwardResponseMessage
)
