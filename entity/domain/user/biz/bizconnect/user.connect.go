// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: domain/user/biz/user.proto

package bizconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	biz "github.com/blackhorseya/godine/entity/domain/user/biz"
	model "github.com/blackhorseya/godine/entity/domain/user/model"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// AccountServiceName is the fully-qualified name of the AccountService service.
	AccountServiceName = "user.AccountService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AccountServiceWhoAmIProcedure is the fully-qualified name of the AccountService's WhoAmI RPC.
	AccountServiceWhoAmIProcedure = "/user.AccountService/WhoAmI"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	accountServiceServiceDescriptor      = biz.File_domain_user_biz_user_proto.Services().ByName("AccountService")
	accountServiceWhoAmIMethodDescriptor = accountServiceServiceDescriptor.Methods().ByName("WhoAmI")
)

// AccountServiceClient is a client for the user.AccountService service.
type AccountServiceClient interface {
	WhoAmI(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[model.Account], error)
}

// NewAccountServiceClient constructs a client for the user.AccountService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAccountServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AccountServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &accountServiceClient{
		whoAmI: connect.NewClient[emptypb.Empty, model.Account](
			httpClient,
			baseURL+AccountServiceWhoAmIProcedure,
			connect.WithSchema(accountServiceWhoAmIMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// accountServiceClient implements AccountServiceClient.
type accountServiceClient struct {
	whoAmI *connect.Client[emptypb.Empty, model.Account]
}

// WhoAmI calls user.AccountService.WhoAmI.
func (c *accountServiceClient) WhoAmI(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[model.Account], error) {
	return c.whoAmI.CallUnary(ctx, req)
}

// AccountServiceHandler is an implementation of the user.AccountService service.
type AccountServiceHandler interface {
	WhoAmI(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[model.Account], error)
}

// NewAccountServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAccountServiceHandler(svc AccountServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	accountServiceWhoAmIHandler := connect.NewUnaryHandler(
		AccountServiceWhoAmIProcedure,
		svc.WhoAmI,
		connect.WithSchema(accountServiceWhoAmIMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/user.AccountService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AccountServiceWhoAmIProcedure:
			accountServiceWhoAmIHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAccountServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAccountServiceHandler struct{}

func (UnimplementedAccountServiceHandler) WhoAmI(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[model.Account], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("user.AccountService.WhoAmI is not implemented"))
}