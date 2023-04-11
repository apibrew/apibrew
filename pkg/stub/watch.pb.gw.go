// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: stub/watch.proto

/*
Package stub is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package stub

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

var (
	filter_Watch_Watch_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_Watch_Watch_0(ctx context.Context, marshaler runtime.Marshaler, client WatchClient, req *http.Request, pathParams map[string]string) (Watch_WatchClient, runtime.ServerMetadata, error) {
	var protoReq WatchRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_Watch_Watch_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.Watch(ctx, &protoReq)
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

// RegisterWatchHandlerServer registers the http handlers for service Watch to "mux".
// UnaryRPC     :call WatchServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterWatchHandlerFromEndpoint instead.
func RegisterWatchHandlerServer(ctx context.Context, mux *runtime.ServeMux, server WatchServer) error {

	mux.Handle("POST", pattern_Watch_Watch_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterWatchHandlerFromEndpoint is same as RegisterWatchHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterWatchHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
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

	return RegisterWatchHandler(ctx, mux, conn)
}

// RegisterWatchHandler registers the http handlers for service Watch to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterWatchHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterWatchHandlerClient(ctx, mux, NewWatchClient(conn))
}

// RegisterWatchHandlerClient registers the http handlers for service Watch
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "WatchClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "WatchClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "WatchClient" to call the correct interceptors.
func RegisterWatchHandlerClient(ctx context.Context, mux *runtime.ServeMux, client WatchClient) error {

	mux.Handle("POST", pattern_Watch_Watch_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/stub.Watch/Watch", runtime.WithHTTPPathPattern("/system/watch"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Watch_Watch_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Watch_Watch_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Watch_Watch_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"system", "watch"}, ""))
)

var (
	forward_Watch_Watch_0 = runtime.ForwardResponseStream
)
