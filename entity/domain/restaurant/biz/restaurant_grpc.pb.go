// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: domain/restaurant/biz/restaurant.proto

package biz

import (
	context "context"
	model "github.com/blackhorseya/godine/entity/domain/restaurant/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RestaurantService_CreateRestaurant_FullMethodName         = "/restaurant.RestaurantService/CreateRestaurant"
	RestaurantService_ListRestaurants_FullMethodName          = "/restaurant.RestaurantService/ListRestaurants"
	RestaurantService_GetRestaurant_FullMethodName            = "/restaurant.RestaurantService/GetRestaurant"
	RestaurantService_ListRestaurantsNonStream_FullMethodName = "/restaurant.RestaurantService/ListRestaurantsNonStream"
	RestaurantService_PlaceOrder_FullMethodName               = "/restaurant.RestaurantService/PlaceOrder"
	RestaurantService_ListOrders_FullMethodName               = "/restaurant.RestaurantService/ListOrders"
)

// RestaurantServiceClient is the client API for RestaurantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RestaurantServiceClient interface {
	CreateRestaurant(ctx context.Context, in *CreateRestaurantRequest, opts ...grpc.CallOption) (*model.Restaurant, error)
	ListRestaurants(ctx context.Context, in *ListRestaurantsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[model.Restaurant], error)
	GetRestaurant(ctx context.Context, in *GetRestaurantRequest, opts ...grpc.CallOption) (*model.Restaurant, error)
	ListRestaurantsNonStream(ctx context.Context, in *ListRestaurantsRequest, opts ...grpc.CallOption) (*ListRestaurantsResponse, error)
	PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...grpc.CallOption) (*PlaceOrderResponse, error)
	ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[model.Order], error)
}

type restaurantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRestaurantServiceClient(cc grpc.ClientConnInterface) RestaurantServiceClient {
	return &restaurantServiceClient{cc}
}

func (c *restaurantServiceClient) CreateRestaurant(ctx context.Context, in *CreateRestaurantRequest, opts ...grpc.CallOption) (*model.Restaurant, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(model.Restaurant)
	err := c.cc.Invoke(ctx, RestaurantService_CreateRestaurant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) ListRestaurants(ctx context.Context, in *ListRestaurantsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[model.Restaurant], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &RestaurantService_ServiceDesc.Streams[0], RestaurantService_ListRestaurants_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ListRestaurantsRequest, model.Restaurant]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RestaurantService_ListRestaurantsClient = grpc.ServerStreamingClient[model.Restaurant]

func (c *restaurantServiceClient) GetRestaurant(ctx context.Context, in *GetRestaurantRequest, opts ...grpc.CallOption) (*model.Restaurant, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(model.Restaurant)
	err := c.cc.Invoke(ctx, RestaurantService_GetRestaurant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) ListRestaurantsNonStream(ctx context.Context, in *ListRestaurantsRequest, opts ...grpc.CallOption) (*ListRestaurantsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRestaurantsResponse)
	err := c.cc.Invoke(ctx, RestaurantService_ListRestaurantsNonStream_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...grpc.CallOption) (*PlaceOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PlaceOrderResponse)
	err := c.cc.Invoke(ctx, RestaurantService_PlaceOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[model.Order], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &RestaurantService_ServiceDesc.Streams[1], RestaurantService_ListOrders_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ListOrdersRequest, model.Order]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RestaurantService_ListOrdersClient = grpc.ServerStreamingClient[model.Order]

// RestaurantServiceServer is the server API for RestaurantService service.
// All implementations should embed UnimplementedRestaurantServiceServer
// for forward compatibility.
type RestaurantServiceServer interface {
	CreateRestaurant(context.Context, *CreateRestaurantRequest) (*model.Restaurant, error)
	ListRestaurants(*ListRestaurantsRequest, grpc.ServerStreamingServer[model.Restaurant]) error
	GetRestaurant(context.Context, *GetRestaurantRequest) (*model.Restaurant, error)
	ListRestaurantsNonStream(context.Context, *ListRestaurantsRequest) (*ListRestaurantsResponse, error)
	PlaceOrder(context.Context, *PlaceOrderRequest) (*PlaceOrderResponse, error)
	ListOrders(*ListOrdersRequest, grpc.ServerStreamingServer[model.Order]) error
}

// UnimplementedRestaurantServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRestaurantServiceServer struct{}

func (UnimplementedRestaurantServiceServer) CreateRestaurant(context.Context, *CreateRestaurantRequest) (*model.Restaurant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRestaurant not implemented")
}
func (UnimplementedRestaurantServiceServer) ListRestaurants(*ListRestaurantsRequest, grpc.ServerStreamingServer[model.Restaurant]) error {
	return status.Errorf(codes.Unimplemented, "method ListRestaurants not implemented")
}
func (UnimplementedRestaurantServiceServer) GetRestaurant(context.Context, *GetRestaurantRequest) (*model.Restaurant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRestaurant not implemented")
}
func (UnimplementedRestaurantServiceServer) ListRestaurantsNonStream(context.Context, *ListRestaurantsRequest) (*ListRestaurantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRestaurantsNonStream not implemented")
}
func (UnimplementedRestaurantServiceServer) PlaceOrder(context.Context, *PlaceOrderRequest) (*PlaceOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}
func (UnimplementedRestaurantServiceServer) ListOrders(*ListOrdersRequest, grpc.ServerStreamingServer[model.Order]) error {
	return status.Errorf(codes.Unimplemented, "method ListOrders not implemented")
}
func (UnimplementedRestaurantServiceServer) testEmbeddedByValue() {}

// UnsafeRestaurantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RestaurantServiceServer will
// result in compilation errors.
type UnsafeRestaurantServiceServer interface {
	mustEmbedUnimplementedRestaurantServiceServer()
}

func RegisterRestaurantServiceServer(s grpc.ServiceRegistrar, srv RestaurantServiceServer) {
	// If the following call pancis, it indicates UnimplementedRestaurantServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RestaurantService_ServiceDesc, srv)
}

func _RestaurantService_CreateRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRestaurantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).CreateRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RestaurantService_CreateRestaurant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).CreateRestaurant(ctx, req.(*CreateRestaurantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_ListRestaurants_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRestaurantsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RestaurantServiceServer).ListRestaurants(m, &grpc.GenericServerStream[ListRestaurantsRequest, model.Restaurant]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RestaurantService_ListRestaurantsServer = grpc.ServerStreamingServer[model.Restaurant]

func _RestaurantService_GetRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRestaurantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).GetRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RestaurantService_GetRestaurant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).GetRestaurant(ctx, req.(*GetRestaurantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_ListRestaurantsNonStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRestaurantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).ListRestaurantsNonStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RestaurantService_ListRestaurantsNonStream_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).ListRestaurantsNonStream(ctx, req.(*ListRestaurantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RestaurantService_PlaceOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).PlaceOrder(ctx, req.(*PlaceOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_ListOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListOrdersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RestaurantServiceServer).ListOrders(m, &grpc.GenericServerStream[ListOrdersRequest, model.Order]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RestaurantService_ListOrdersServer = grpc.ServerStreamingServer[model.Order]

// RestaurantService_ServiceDesc is the grpc.ServiceDesc for RestaurantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RestaurantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "restaurant.RestaurantService",
	HandlerType: (*RestaurantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRestaurant",
			Handler:    _RestaurantService_CreateRestaurant_Handler,
		},
		{
			MethodName: "GetRestaurant",
			Handler:    _RestaurantService_GetRestaurant_Handler,
		},
		{
			MethodName: "ListRestaurantsNonStream",
			Handler:    _RestaurantService_ListRestaurantsNonStream_Handler,
		},
		{
			MethodName: "PlaceOrder",
			Handler:    _RestaurantService_PlaceOrder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListRestaurants",
			Handler:       _RestaurantService_ListRestaurants_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListOrders",
			Handler:       _RestaurantService_ListOrders_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "domain/restaurant/biz/restaurant.proto",
}
