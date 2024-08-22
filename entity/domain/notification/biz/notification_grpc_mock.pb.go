// Code generated by protoc-gen-go-grpc-mock. DO NOT EDIT.
// source: entity/domain/notification/biz/notification.proto

package biz

import (
	context "context"
	reflect "reflect"

	model "github.com/blackhorseya/godine/entity/domain/notification/model"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockNotificationService_ListMyNotificationsClient is a mock of NotificationService_ListMyNotificationsClient interface.
type MockNotificationService_ListMyNotificationsClient struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationService_ListMyNotificationsClientMockRecorder
}

// MockNotificationService_ListMyNotificationsClientMockRecorder is the mock recorder for MockNotificationService_ListMyNotificationsClient.
type MockNotificationService_ListMyNotificationsClientMockRecorder struct {
	mock *MockNotificationService_ListMyNotificationsClient
}

// NewMockNotificationService_ListMyNotificationsClient creates a new mock instance.
func NewMockNotificationService_ListMyNotificationsClient(ctrl *gomock.Controller) *MockNotificationService_ListMyNotificationsClient {
	mock := &MockNotificationService_ListMyNotificationsClient{ctrl: ctrl}
	mock.recorder = &MockNotificationService_ListMyNotificationsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationService_ListMyNotificationsClient) EXPECT() *MockNotificationService_ListMyNotificationsClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockNotificationService_ListMyNotificationsClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockNotificationService_ListMyNotificationsClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockNotificationService_ListMyNotificationsClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockNotificationService_ListMyNotificationsClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockNotificationService_ListMyNotificationsClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockNotificationService_ListMyNotificationsClient)(nil).Context))
}

// Header mocks base method.
func (m *MockNotificationService_ListMyNotificationsClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockNotificationService_ListMyNotificationsClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockNotificationService_ListMyNotificationsClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockNotificationService_ListMyNotificationsClient) Recv() (*model.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*model.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockNotificationService_ListMyNotificationsClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockNotificationService_ListMyNotificationsClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockNotificationService_ListMyNotificationsClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockNotificationService_ListMyNotificationsClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockNotificationService_ListMyNotificationsClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method.
func (m *MockNotificationService_ListMyNotificationsClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockNotificationService_ListMyNotificationsClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockNotificationService_ListMyNotificationsClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockNotificationService_ListMyNotificationsClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockNotificationService_ListMyNotificationsClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockNotificationService_ListMyNotificationsClient)(nil).Trailer))
}

// MockNotificationService_ListMyNotificationsServer is a mock of NotificationService_ListMyNotificationsServer interface.
type MockNotificationService_ListMyNotificationsServer struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationService_ListMyNotificationsServerMockRecorder
}

// MockNotificationService_ListMyNotificationsServerMockRecorder is the mock recorder for MockNotificationService_ListMyNotificationsServer.
type MockNotificationService_ListMyNotificationsServerMockRecorder struct {
	mock *MockNotificationService_ListMyNotificationsServer
}

// NewMockNotificationService_ListMyNotificationsServer creates a new mock instance.
func NewMockNotificationService_ListMyNotificationsServer(ctrl *gomock.Controller) *MockNotificationService_ListMyNotificationsServer {
	mock := &MockNotificationService_ListMyNotificationsServer{ctrl: ctrl}
	mock.recorder = &MockNotificationService_ListMyNotificationsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationService_ListMyNotificationsServer) EXPECT() *MockNotificationService_ListMyNotificationsServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockNotificationService_ListMyNotificationsServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockNotificationService_ListMyNotificationsServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockNotificationService_ListMyNotificationsServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m *MockNotificationService_ListMyNotificationsServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockNotificationService_ListMyNotificationsServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockNotificationService_ListMyNotificationsServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockNotificationService_ListMyNotificationsServer) Send(arg0 *model.Notification) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockNotificationService_ListMyNotificationsServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockNotificationService_ListMyNotificationsServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockNotificationService_ListMyNotificationsServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockNotificationService_ListMyNotificationsServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockNotificationService_ListMyNotificationsServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockNotificationService_ListMyNotificationsServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockNotificationService_ListMyNotificationsServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockNotificationService_ListMyNotificationsServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockNotificationService_ListMyNotificationsServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockNotificationService_ListMyNotificationsServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockNotificationService_ListMyNotificationsServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockNotificationService_ListMyNotificationsServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockNotificationService_ListMyNotificationsServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockNotificationService_ListMyNotificationsServer)(nil).SetTrailer), arg0)
}

// MockNotificationServiceClient is a mock of NotificationServiceClient interface.
type MockNotificationServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationServiceClientMockRecorder
}

// MockNotificationServiceClientMockRecorder is the mock recorder for MockNotificationServiceClient.
type MockNotificationServiceClientMockRecorder struct {
	mock *MockNotificationServiceClient
}

// NewMockNotificationServiceClient creates a new mock instance.
func NewMockNotificationServiceClient(ctrl *gomock.Controller) *MockNotificationServiceClient {
	mock := &MockNotificationServiceClient{ctrl: ctrl}
	mock.recorder = &MockNotificationServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationServiceClient) EXPECT() *MockNotificationServiceClientMockRecorder {
	return m.recorder
}

// ListMyNotifications mocks base method.
func (m *MockNotificationServiceClient) ListMyNotifications(ctx context.Context, in *ListMyNotificationsRequest, opts ...grpc.CallOption) (NotificationService_ListMyNotificationsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMyNotifications", varargs...)
	ret0, _ := ret[0].(NotificationService_ListMyNotificationsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMyNotifications indicates an expected call of ListMyNotifications.
func (mr *MockNotificationServiceClientMockRecorder) ListMyNotifications(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMyNotifications", reflect.TypeOf((*MockNotificationServiceClient)(nil).ListMyNotifications), varargs...)
}

// SendNotification mocks base method.
func (m *MockNotificationServiceClient) SendNotification(ctx context.Context, in *SendNotificationRequest, opts ...grpc.CallOption) (*model.Notification, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendNotification", varargs...)
	ret0, _ := ret[0].(*model.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendNotification indicates an expected call of SendNotification.
func (mr *MockNotificationServiceClientMockRecorder) SendNotification(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendNotification", reflect.TypeOf((*MockNotificationServiceClient)(nil).SendNotification), varargs...)
}

// MockNotificationServiceServer is a mock of NotificationServiceServer interface.
type MockNotificationServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationServiceServerMockRecorder
}

// MockNotificationServiceServerMockRecorder is the mock recorder for MockNotificationServiceServer.
type MockNotificationServiceServerMockRecorder struct {
	mock *MockNotificationServiceServer
}

// NewMockNotificationServiceServer creates a new mock instance.
func NewMockNotificationServiceServer(ctrl *gomock.Controller) *MockNotificationServiceServer {
	mock := &MockNotificationServiceServer{ctrl: ctrl}
	mock.recorder = &MockNotificationServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationServiceServer) EXPECT() *MockNotificationServiceServerMockRecorder {
	return m.recorder
}

// ListMyNotifications mocks base method.
func (m *MockNotificationServiceServer) ListMyNotifications(blob *ListMyNotificationsRequest, server NotificationService_ListMyNotificationsServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMyNotifications", blob, server)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListMyNotifications indicates an expected call of ListMyNotifications.
func (mr *MockNotificationServiceServerMockRecorder) ListMyNotifications(blob, server interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMyNotifications", reflect.TypeOf((*MockNotificationServiceServer)(nil).ListMyNotifications), blob, server)
}

// SendNotification mocks base method.
func (m *MockNotificationServiceServer) SendNotification(ctx context.Context, in *SendNotificationRequest) (*model.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendNotification", ctx, in)
	ret0, _ := ret[0].(*model.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendNotification indicates an expected call of SendNotification.
func (mr *MockNotificationServiceServerMockRecorder) SendNotification(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendNotification", reflect.TypeOf((*MockNotificationServiceServer)(nil).SendNotification), ctx, in)
}