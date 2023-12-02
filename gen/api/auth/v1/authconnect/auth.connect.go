// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/auth/v1/auth.proto

// # Auth service (v1)

package authconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/bxxf/go-template/gen/api/auth/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// AuthServiceName is the fully-qualified name of the AuthService service.
	AuthServiceName = "auth.v1.AuthService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AuthServiceRegisterUserProcedure is the fully-qualified name of the AuthService's RegisterUser
	// RPC.
	AuthServiceRegisterUserProcedure = "/auth.v1.AuthService/RegisterUser"
)

// AuthServiceClient is a client for the auth.v1.AuthService service.
type AuthServiceClient interface {
	RegisterUser(context.Context, *connect.Request[v1.RegisterUserRequest]) (*connect.Response[v1.RegisterUserResponse], error)
}

// NewAuthServiceClient constructs a client for the auth.v1.AuthService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuthServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AuthServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &authServiceClient{
		registerUser: connect.NewClient[v1.RegisterUserRequest, v1.RegisterUserResponse](
			httpClient,
			baseURL+AuthServiceRegisterUserProcedure,
			opts...,
		),
	}
}

// authServiceClient implements AuthServiceClient.
type authServiceClient struct {
	registerUser *connect.Client[v1.RegisterUserRequest, v1.RegisterUserResponse]
}

// RegisterUser calls auth.v1.AuthService.RegisterUser.
func (c *authServiceClient) RegisterUser(ctx context.Context, req *connect.Request[v1.RegisterUserRequest]) (*connect.Response[v1.RegisterUserResponse], error) {
	return c.registerUser.CallUnary(ctx, req)
}

// AuthServiceHandler is an implementation of the auth.v1.AuthService service.
type AuthServiceHandler interface {
	RegisterUser(context.Context, *connect.Request[v1.RegisterUserRequest]) (*connect.Response[v1.RegisterUserResponse], error)
}

// NewAuthServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuthServiceHandler(svc AuthServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	authServiceRegisterUserHandler := connect.NewUnaryHandler(
		AuthServiceRegisterUserProcedure,
		svc.RegisterUser,
		opts...,
	)
	return "/auth.v1.AuthService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AuthServiceRegisterUserProcedure:
			authServiceRegisterUserHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAuthServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAuthServiceHandler struct{}

func (UnimplementedAuthServiceHandler) RegisterUser(context.Context, *connect.Request[v1.RegisterUserRequest]) (*connect.Response[v1.RegisterUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("auth.v1.AuthService.RegisterUser is not implemented"))
}
