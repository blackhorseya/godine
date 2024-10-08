// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: domain/payment/biz/payment.proto

package bizconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	biz "github.com/blackhorseya/godine/entity/domain/payment/biz"
	model "github.com/blackhorseya/godine/entity/domain/payment/model"
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
	// PaymentServiceName is the fully-qualified name of the PaymentService service.
	PaymentServiceName = "payment.PaymentService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PaymentServiceCreatePaymentProcedure is the fully-qualified name of the PaymentService's
	// CreatePayment RPC.
	PaymentServiceCreatePaymentProcedure = "/payment.PaymentService/CreatePayment"
	// PaymentServiceGetPaymentProcedure is the fully-qualified name of the PaymentService's GetPayment
	// RPC.
	PaymentServiceGetPaymentProcedure = "/payment.PaymentService/GetPayment"
	// PaymentServiceListPaymentsProcedure is the fully-qualified name of the PaymentService's
	// ListPayments RPC.
	PaymentServiceListPaymentsProcedure = "/payment.PaymentService/ListPayments"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	paymentServiceServiceDescriptor             = biz.File_domain_payment_biz_payment_proto.Services().ByName("PaymentService")
	paymentServiceCreatePaymentMethodDescriptor = paymentServiceServiceDescriptor.Methods().ByName("CreatePayment")
	paymentServiceGetPaymentMethodDescriptor    = paymentServiceServiceDescriptor.Methods().ByName("GetPayment")
	paymentServiceListPaymentsMethodDescriptor  = paymentServiceServiceDescriptor.Methods().ByName("ListPayments")
)

// PaymentServiceClient is a client for the payment.PaymentService service.
type PaymentServiceClient interface {
	CreatePayment(context.Context, *connect.Request[biz.CreatePaymentRequest]) (*connect.Response[model.Payment], error)
	GetPayment(context.Context, *connect.Request[biz.GetPaymentRequest]) (*connect.Response[model.Payment], error)
	ListPayments(context.Context, *connect.Request[biz.ListPaymentsRequest]) (*connect.Response[biz.ListPaymentsResponse], error)
}

// NewPaymentServiceClient constructs a client for the payment.PaymentService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPaymentServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) PaymentServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &paymentServiceClient{
		createPayment: connect.NewClient[biz.CreatePaymentRequest, model.Payment](
			httpClient,
			baseURL+PaymentServiceCreatePaymentProcedure,
			connect.WithSchema(paymentServiceCreatePaymentMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getPayment: connect.NewClient[biz.GetPaymentRequest, model.Payment](
			httpClient,
			baseURL+PaymentServiceGetPaymentProcedure,
			connect.WithSchema(paymentServiceGetPaymentMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listPayments: connect.NewClient[biz.ListPaymentsRequest, biz.ListPaymentsResponse](
			httpClient,
			baseURL+PaymentServiceListPaymentsProcedure,
			connect.WithSchema(paymentServiceListPaymentsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// paymentServiceClient implements PaymentServiceClient.
type paymentServiceClient struct {
	createPayment *connect.Client[biz.CreatePaymentRequest, model.Payment]
	getPayment    *connect.Client[biz.GetPaymentRequest, model.Payment]
	listPayments  *connect.Client[biz.ListPaymentsRequest, biz.ListPaymentsResponse]
}

// CreatePayment calls payment.PaymentService.CreatePayment.
func (c *paymentServiceClient) CreatePayment(ctx context.Context, req *connect.Request[biz.CreatePaymentRequest]) (*connect.Response[model.Payment], error) {
	return c.createPayment.CallUnary(ctx, req)
}

// GetPayment calls payment.PaymentService.GetPayment.
func (c *paymentServiceClient) GetPayment(ctx context.Context, req *connect.Request[biz.GetPaymentRequest]) (*connect.Response[model.Payment], error) {
	return c.getPayment.CallUnary(ctx, req)
}

// ListPayments calls payment.PaymentService.ListPayments.
func (c *paymentServiceClient) ListPayments(ctx context.Context, req *connect.Request[biz.ListPaymentsRequest]) (*connect.Response[biz.ListPaymentsResponse], error) {
	return c.listPayments.CallUnary(ctx, req)
}

// PaymentServiceHandler is an implementation of the payment.PaymentService service.
type PaymentServiceHandler interface {
	CreatePayment(context.Context, *connect.Request[biz.CreatePaymentRequest]) (*connect.Response[model.Payment], error)
	GetPayment(context.Context, *connect.Request[biz.GetPaymentRequest]) (*connect.Response[model.Payment], error)
	ListPayments(context.Context, *connect.Request[biz.ListPaymentsRequest]) (*connect.Response[biz.ListPaymentsResponse], error)
}

// NewPaymentServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPaymentServiceHandler(svc PaymentServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	paymentServiceCreatePaymentHandler := connect.NewUnaryHandler(
		PaymentServiceCreatePaymentProcedure,
		svc.CreatePayment,
		connect.WithSchema(paymentServiceCreatePaymentMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	paymentServiceGetPaymentHandler := connect.NewUnaryHandler(
		PaymentServiceGetPaymentProcedure,
		svc.GetPayment,
		connect.WithSchema(paymentServiceGetPaymentMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	paymentServiceListPaymentsHandler := connect.NewUnaryHandler(
		PaymentServiceListPaymentsProcedure,
		svc.ListPayments,
		connect.WithSchema(paymentServiceListPaymentsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/payment.PaymentService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PaymentServiceCreatePaymentProcedure:
			paymentServiceCreatePaymentHandler.ServeHTTP(w, r)
		case PaymentServiceGetPaymentProcedure:
			paymentServiceGetPaymentHandler.ServeHTTP(w, r)
		case PaymentServiceListPaymentsProcedure:
			paymentServiceListPaymentsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPaymentServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPaymentServiceHandler struct{}

func (UnimplementedPaymentServiceHandler) CreatePayment(context.Context, *connect.Request[biz.CreatePaymentRequest]) (*connect.Response[model.Payment], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("payment.PaymentService.CreatePayment is not implemented"))
}

func (UnimplementedPaymentServiceHandler) GetPayment(context.Context, *connect.Request[biz.GetPaymentRequest]) (*connect.Response[model.Payment], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("payment.PaymentService.GetPayment is not implemented"))
}

func (UnimplementedPaymentServiceHandler) ListPayments(context.Context, *connect.Request[biz.ListPaymentsRequest]) (*connect.Response[biz.ListPaymentsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("payment.PaymentService.ListPayments is not implemented"))
}
