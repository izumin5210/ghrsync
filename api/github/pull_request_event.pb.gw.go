// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: pull_request_event.proto

/*
Package github_pb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package github_pb

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_PullRequestEventService_CreatePullRequestEvent_0(ctx context.Context, marshaler runtime.Marshaler, client PullRequestEventServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreatePullRequestEventRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq.PullRequestEvent); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreatePullRequestEvent(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterPullRequestEventServiceHandlerFromEndpoint is same as RegisterPullRequestEventServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterPullRequestEventServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterPullRequestEventServiceHandler(ctx, mux, conn)
}

// RegisterPullRequestEventServiceHandler registers the http handlers for service PullRequestEventService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterPullRequestEventServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterPullRequestEventServiceHandlerClient(ctx, mux, NewPullRequestEventServiceClient(conn))
}

// RegisterPullRequestEventServiceHandlerClient registers the http handlers for service PullRequestEventService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "PullRequestEventServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "PullRequestEventServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "PullRequestEventServiceClient" to call the correct interceptors.
func RegisterPullRequestEventServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client PullRequestEventServiceClient) error {

	mux.Handle("POST", pattern_PullRequestEventService_CreatePullRequestEvent_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_PullRequestEventService_CreatePullRequestEvent_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PullRequestEventService_CreatePullRequestEvent_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_PullRequestEventService_CreatePullRequestEvent_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"github", "pull_request_events"}, ""))
)

var (
	forward_PullRequestEventService_CreatePullRequestEvent_0 = runtime.ForwardResponseMessage
)