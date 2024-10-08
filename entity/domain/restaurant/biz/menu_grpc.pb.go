// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: domain/restaurant/biz/menu.proto

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
	MenuService_AddMenuItem_FullMethodName   = "/restaurant.MenuService/AddMenuItem"
	MenuService_GetMenuItem_FullMethodName   = "/restaurant.MenuService/GetMenuItem"
	MenuService_ListMenuItems_FullMethodName = "/restaurant.MenuService/ListMenuItems"
)

// MenuServiceClient is the client API for MenuService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MenuServiceClient interface {
	AddMenuItem(ctx context.Context, in *AddMenuItemRequest, opts ...grpc.CallOption) (*model.MenuItem, error)
	GetMenuItem(ctx context.Context, in *GetMenuItemRequest, opts ...grpc.CallOption) (*model.MenuItem, error)
	ListMenuItems(ctx context.Context, in *ListMenuItemsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[model.MenuItem], error)
}

type menuServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMenuServiceClient(cc grpc.ClientConnInterface) MenuServiceClient {
	return &menuServiceClient{cc}
}

func (c *menuServiceClient) AddMenuItem(ctx context.Context, in *AddMenuItemRequest, opts ...grpc.CallOption) (*model.MenuItem, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(model.MenuItem)
	err := c.cc.Invoke(ctx, MenuService_AddMenuItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) GetMenuItem(ctx context.Context, in *GetMenuItemRequest, opts ...grpc.CallOption) (*model.MenuItem, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(model.MenuItem)
	err := c.cc.Invoke(ctx, MenuService_GetMenuItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) ListMenuItems(ctx context.Context, in *ListMenuItemsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[model.MenuItem], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MenuService_ServiceDesc.Streams[0], MenuService_ListMenuItems_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ListMenuItemsRequest, model.MenuItem]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MenuService_ListMenuItemsClient = grpc.ServerStreamingClient[model.MenuItem]

// MenuServiceServer is the server API for MenuService service.
// All implementations should embed UnimplementedMenuServiceServer
// for forward compatibility.
type MenuServiceServer interface {
	AddMenuItem(context.Context, *AddMenuItemRequest) (*model.MenuItem, error)
	GetMenuItem(context.Context, *GetMenuItemRequest) (*model.MenuItem, error)
	ListMenuItems(*ListMenuItemsRequest, grpc.ServerStreamingServer[model.MenuItem]) error
}

// UnimplementedMenuServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMenuServiceServer struct{}

func (UnimplementedMenuServiceServer) AddMenuItem(context.Context, *AddMenuItemRequest) (*model.MenuItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMenuItem not implemented")
}
func (UnimplementedMenuServiceServer) GetMenuItem(context.Context, *GetMenuItemRequest) (*model.MenuItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMenuItem not implemented")
}
func (UnimplementedMenuServiceServer) ListMenuItems(*ListMenuItemsRequest, grpc.ServerStreamingServer[model.MenuItem]) error {
	return status.Errorf(codes.Unimplemented, "method ListMenuItems not implemented")
}
func (UnimplementedMenuServiceServer) testEmbeddedByValue() {}

// UnsafeMenuServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MenuServiceServer will
// result in compilation errors.
type UnsafeMenuServiceServer interface {
	mustEmbedUnimplementedMenuServiceServer()
}

func RegisterMenuServiceServer(s grpc.ServiceRegistrar, srv MenuServiceServer) {
	// If the following call pancis, it indicates UnimplementedMenuServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MenuService_ServiceDesc, srv)
}

func _MenuService_AddMenuItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMenuItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).AddMenuItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_AddMenuItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).AddMenuItem(ctx, req.(*AddMenuItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_GetMenuItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMenuItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).GetMenuItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_GetMenuItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).GetMenuItem(ctx, req.(*GetMenuItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_ListMenuItems_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListMenuItemsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MenuServiceServer).ListMenuItems(m, &grpc.GenericServerStream[ListMenuItemsRequest, model.MenuItem]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MenuService_ListMenuItemsServer = grpc.ServerStreamingServer[model.MenuItem]

// MenuService_ServiceDesc is the grpc.ServiceDesc for MenuService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MenuService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "restaurant.MenuService",
	HandlerType: (*MenuServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMenuItem",
			Handler:    _MenuService_AddMenuItem_Handler,
		},
		{
			MethodName: "GetMenuItem",
			Handler:    _MenuService_GetMenuItem_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListMenuItems",
			Handler:       _MenuService_ListMenuItems_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "domain/restaurant/biz/menu.proto",
}
