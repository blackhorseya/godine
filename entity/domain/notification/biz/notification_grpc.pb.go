// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: entity/domain/notification/biz/notification.proto

package biz

import (
	context "context"
	model "github.com/blackhorseya/godine/entity/domain/notification/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	NotificationService_SendNotification_FullMethodName    = "/notification.NotificationService/SendNotification"
	NotificationService_ListMyNotifications_FullMethodName = "/notification.NotificationService/ListMyNotifications"
)

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationServiceClient interface {
	SendNotification(ctx context.Context, in *SendNotificationRequest, opts ...grpc.CallOption) (*model.Notification, error)
	ListMyNotifications(ctx context.Context, in *ListMyNotificationsRequest, opts ...grpc.CallOption) (NotificationService_ListMyNotificationsClient, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) SendNotification(ctx context.Context, in *SendNotificationRequest, opts ...grpc.CallOption) (*model.Notification, error) {
	out := new(model.Notification)
	err := c.cc.Invoke(ctx, NotificationService_SendNotification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) ListMyNotifications(ctx context.Context, in *ListMyNotificationsRequest, opts ...grpc.CallOption) (NotificationService_ListMyNotificationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &NotificationService_ServiceDesc.Streams[0], NotificationService_ListMyNotifications_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &notificationServiceListMyNotificationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NotificationService_ListMyNotificationsClient interface {
	Recv() (*model.Notification, error)
	grpc.ClientStream
}

type notificationServiceListMyNotificationsClient struct {
	grpc.ClientStream
}

func (x *notificationServiceListMyNotificationsClient) Recv() (*model.Notification, error) {
	m := new(model.Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations should embed UnimplementedNotificationServiceServer
// for forward compatibility
type NotificationServiceServer interface {
	SendNotification(context.Context, *SendNotificationRequest) (*model.Notification, error)
	ListMyNotifications(*ListMyNotificationsRequest, NotificationService_ListMyNotificationsServer) error
}

// UnimplementedNotificationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (UnimplementedNotificationServiceServer) SendNotification(context.Context, *SendNotificationRequest) (*model.Notification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNotification not implemented")
}
func (UnimplementedNotificationServiceServer) ListMyNotifications(*ListMyNotificationsRequest, NotificationService_ListMyNotificationsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListMyNotifications not implemented")
}

// UnsafeNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServiceServer will
// result in compilation errors.
type UnsafeNotificationServiceServer interface {
	mustEmbedUnimplementedNotificationServiceServer()
}

func RegisterNotificationServiceServer(s grpc.ServiceRegistrar, srv NotificationServiceServer) {
	s.RegisterService(&NotificationService_ServiceDesc, srv)
}

func _NotificationService_SendNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).SendNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_SendNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).SendNotification(ctx, req.(*SendNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_ListMyNotifications_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListMyNotificationsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NotificationServiceServer).ListMyNotifications(m, &notificationServiceListMyNotificationsServer{stream})
}

type NotificationService_ListMyNotificationsServer interface {
	Send(*model.Notification) error
	grpc.ServerStream
}

type notificationServiceListMyNotificationsServer struct {
	grpc.ServerStream
}

func (x *notificationServiceListMyNotificationsServer) Send(m *model.Notification) error {
	return x.ServerStream.SendMsg(m)
}

// NotificationService_ServiceDesc is the grpc.ServiceDesc for NotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notification.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendNotification",
			Handler:    _NotificationService_SendNotification_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListMyNotifications",
			Handler:       _NotificationService_ListMyNotifications_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "entity/domain/notification/biz/notification.proto",
}