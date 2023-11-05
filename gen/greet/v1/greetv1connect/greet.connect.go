// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: greet/v1/greet.proto

package greetv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"
	v1 "vote-dapp-backend/proto/gen/greet/v1"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// GreetServiceName is the fully-qualified name of the GreetService service.
	GreetServiceName = "greet.v1.GreetService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GreetServiceGreetProcedure is the fully-qualified name of the GreetService's Greet RPC.
	GreetServiceGreetProcedure = "/greet.v1.GreetService/Greet"
)

// GreetServiceClient is a client for the greet.v1.GreetService service.
type GreetServiceClient interface {
	Greet(context.Context, *connect.Request[v1.GreetRequest]) (*connect.Response[v1.GreetResponse], error)
}

// NewGreetServiceClient constructs a client for the greet.v1.GreetService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGreetServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) GreetServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &greetServiceClient{
		greet: connect.NewClient[v1.GreetRequest, v1.GreetResponse](
			httpClient,
			baseURL+GreetServiceGreetProcedure,
			opts...,
		),
	}
}

// greetServiceClient implements GreetServiceClient.
type greetServiceClient struct {
	greet *connect.Client[v1.GreetRequest, v1.GreetResponse]
}

// Greet calls greet.v1.GreetService.Greet.
func (c *greetServiceClient) Greet(ctx context.Context, req *connect.Request[v1.GreetRequest]) (*connect.Response[v1.GreetResponse], error) {
	return c.greet.CallUnary(ctx, req)
}

// GreetServiceHandler is an implementation of the greet.v1.GreetService service.
type GreetServiceHandler interface {
	Greet(context.Context, *connect.Request[v1.GreetRequest]) (*connect.Response[v1.GreetResponse], error)
}

// NewGreetServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGreetServiceHandler(svc GreetServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	greetServiceGreetHandler := connect.NewUnaryHandler(
		GreetServiceGreetProcedure,
		svc.Greet,
		opts...,
	)
	return "/greet.v1.GreetService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GreetServiceGreetProcedure:
			greetServiceGreetHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGreetServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGreetServiceHandler struct{}

func (UnimplementedGreetServiceHandler) Greet(context.Context, *connect.Request[v1.GreetRequest]) (*connect.Response[v1.GreetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("greet.v1.GreetService.Greet is not implemented"))
}
