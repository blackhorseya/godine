// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: entity/domain/order/biz/order.proto

package biz

import (
	context "context"
	model "github.com/blackhorseya/godine/entity/domain/order/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	OrderService_SubmitOrder_FullMethodName = "/order.OrderService/SubmitOrder"
	OrderService_ListOrders_FullMethodName  = "/order.OrderService/ListOrders"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	SubmitOrder(ctx context.Context, in *SubmitOrderRequest, opts ...grpc.CallOption) (*model.Order, error)
	ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (OrderService_ListOrdersClient, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) SubmitOrder(ctx context.Context, in *SubmitOrderRequest, opts ...grpc.CallOption) (*model.Order, error) {
	out := new(model.Order)
	err := c.cc.Invoke(ctx, OrderService_SubmitOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (OrderService_ListOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderService_ServiceDesc.Streams[0], OrderService_ListOrders_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &orderServiceListOrdersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderService_ListOrdersClient interface {
	Recv() (*model.Order, error)
	grpc.ClientStream
}

type orderServiceListOrdersClient struct {
	grpc.ClientStream
}

func (x *orderServiceListOrdersClient) Recv() (*model.Order, error) {
	m := new(model.Order)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations should embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	SubmitOrder(context.Context, *SubmitOrderRequest) (*model.Order, error)
	ListOrders(*ListOrdersRequest, OrderService_ListOrdersServer) error
}

// UnimplementedOrderServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) SubmitOrder(context.Context, *SubmitOrderRequest) (*model.Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitOrder not implemented")
}
func (UnimplementedOrderServiceServer) ListOrders(*ListOrdersRequest, OrderService_ListOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_SubmitOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).SubmitOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_SubmitOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).SubmitOrder(ctx, req.(*SubmitOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ListOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListOrdersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderServiceServer).ListOrders(m, &orderServiceListOrdersServer{stream})
}

type OrderService_ListOrdersServer interface {
	Send(*model.Order) error
	grpc.ServerStream
}

type orderServiceListOrdersServer struct {
	grpc.ServerStream
}

func (x *orderServiceListOrdersServer) Send(m *model.Order) error {
	return x.ServerStream.SendMsg(m)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitOrder",
			Handler:    _OrderService_SubmitOrder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListOrders",
			Handler:       _OrderService_ListOrders_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "entity/domain/order/biz/order.proto",
}