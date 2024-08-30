// Code generated by protoc-gen-go-grpc-mock. DO NOT EDIT.
// source: domain/restaurant/biz/restaurant.proto

package biz

import (
	context "context"
	reflect "reflect"

	model "github.com/blackhorseya/godine/entity/domain/restaurant/model"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockRestaurantService_ListRestaurantsClient is a mock of RestaurantService_ListRestaurantsClient interface.
type MockRestaurantService_ListRestaurantsClient struct {
	ctrl     *gomock.Controller
	recorder *MockRestaurantService_ListRestaurantsClientMockRecorder
}

// MockRestaurantService_ListRestaurantsClientMockRecorder is the mock recorder for MockRestaurantService_ListRestaurantsClient.
type MockRestaurantService_ListRestaurantsClientMockRecorder struct {
	mock *MockRestaurantService_ListRestaurantsClient
}

// NewMockRestaurantService_ListRestaurantsClient creates a new mock instance.
func NewMockRestaurantService_ListRestaurantsClient(ctrl *gomock.Controller) *MockRestaurantService_ListRestaurantsClient {
	mock := &MockRestaurantService_ListRestaurantsClient{ctrl: ctrl}
	mock.recorder = &MockRestaurantService_ListRestaurantsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRestaurantService_ListRestaurantsClient) EXPECT() *MockRestaurantService_ListRestaurantsClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockRestaurantService_ListRestaurantsClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockRestaurantService_ListRestaurantsClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockRestaurantService_ListRestaurantsClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockRestaurantService_ListRestaurantsClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockRestaurantService_ListRestaurantsClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockRestaurantService_ListRestaurantsClient)(nil).Context))
}

// Header mocks base method.
func (m *MockRestaurantService_ListRestaurantsClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockRestaurantService_ListRestaurantsClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockRestaurantService_ListRestaurantsClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockRestaurantService_ListRestaurantsClient) Recv() (*model.Restaurant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*model.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockRestaurantService_ListRestaurantsClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockRestaurantService_ListRestaurantsClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockRestaurantService_ListRestaurantsClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockRestaurantService_ListRestaurantsClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockRestaurantService_ListRestaurantsClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method.
func (m *MockRestaurantService_ListRestaurantsClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockRestaurantService_ListRestaurantsClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockRestaurantService_ListRestaurantsClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockRestaurantService_ListRestaurantsClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockRestaurantService_ListRestaurantsClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockRestaurantService_ListRestaurantsClient)(nil).Trailer))
}

// MockRestaurantService_ListRestaurantsServer is a mock of RestaurantService_ListRestaurantsServer interface.
type MockRestaurantService_ListRestaurantsServer struct {
	ctrl     *gomock.Controller
	recorder *MockRestaurantService_ListRestaurantsServerMockRecorder
}

// MockRestaurantService_ListRestaurantsServerMockRecorder is the mock recorder for MockRestaurantService_ListRestaurantsServer.
type MockRestaurantService_ListRestaurantsServerMockRecorder struct {
	mock *MockRestaurantService_ListRestaurantsServer
}

// NewMockRestaurantService_ListRestaurantsServer creates a new mock instance.
func NewMockRestaurantService_ListRestaurantsServer(ctrl *gomock.Controller) *MockRestaurantService_ListRestaurantsServer {
	mock := &MockRestaurantService_ListRestaurantsServer{ctrl: ctrl}
	mock.recorder = &MockRestaurantService_ListRestaurantsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRestaurantService_ListRestaurantsServer) EXPECT() *MockRestaurantService_ListRestaurantsServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockRestaurantService_ListRestaurantsServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockRestaurantService_ListRestaurantsServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockRestaurantService_ListRestaurantsServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m *MockRestaurantService_ListRestaurantsServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockRestaurantService_ListRestaurantsServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockRestaurantService_ListRestaurantsServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockRestaurantService_ListRestaurantsServer) Send(arg0 *model.Restaurant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockRestaurantService_ListRestaurantsServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockRestaurantService_ListRestaurantsServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockRestaurantService_ListRestaurantsServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockRestaurantService_ListRestaurantsServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockRestaurantService_ListRestaurantsServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockRestaurantService_ListRestaurantsServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockRestaurantService_ListRestaurantsServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockRestaurantService_ListRestaurantsServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockRestaurantService_ListRestaurantsServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockRestaurantService_ListRestaurantsServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockRestaurantService_ListRestaurantsServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockRestaurantService_ListRestaurantsServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockRestaurantService_ListRestaurantsServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockRestaurantService_ListRestaurantsServer)(nil).SetTrailer), arg0)
}

// MockRestaurantServiceClient is a mock of RestaurantServiceClient interface.
type MockRestaurantServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockRestaurantServiceClientMockRecorder
}

// MockRestaurantServiceClientMockRecorder is the mock recorder for MockRestaurantServiceClient.
type MockRestaurantServiceClientMockRecorder struct {
	mock *MockRestaurantServiceClient
}

// NewMockRestaurantServiceClient creates a new mock instance.
func NewMockRestaurantServiceClient(ctrl *gomock.Controller) *MockRestaurantServiceClient {
	mock := &MockRestaurantServiceClient{ctrl: ctrl}
	mock.recorder = &MockRestaurantServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRestaurantServiceClient) EXPECT() *MockRestaurantServiceClientMockRecorder {
	return m.recorder
}

// CreateRestaurant mocks base method.
func (m *MockRestaurantServiceClient) CreateRestaurant(ctx context.Context, in *CreateRestaurantRequest, opts ...grpc.CallOption) (*model.Restaurant, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateRestaurant", varargs...)
	ret0, _ := ret[0].(*model.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRestaurant indicates an expected call of CreateRestaurant.
func (mr *MockRestaurantServiceClientMockRecorder) CreateRestaurant(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRestaurant", reflect.TypeOf((*MockRestaurantServiceClient)(nil).CreateRestaurant), varargs...)
}

// GetRestaurant mocks base method.
func (m *MockRestaurantServiceClient) GetRestaurant(ctx context.Context, in *GetRestaurantRequest, opts ...grpc.CallOption) (*model.Restaurant, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRestaurant", varargs...)
	ret0, _ := ret[0].(*model.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurant indicates an expected call of GetRestaurant.
func (mr *MockRestaurantServiceClientMockRecorder) GetRestaurant(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurant", reflect.TypeOf((*MockRestaurantServiceClient)(nil).GetRestaurant), varargs...)
}

// ListRestaurants mocks base method.
func (m *MockRestaurantServiceClient) ListRestaurants(ctx context.Context, in *ListRestaurantsRequest, opts ...grpc.CallOption) (RestaurantService_ListRestaurantsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRestaurants", varargs...)
	ret0, _ := ret[0].(RestaurantService_ListRestaurantsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRestaurants indicates an expected call of ListRestaurants.
func (mr *MockRestaurantServiceClientMockRecorder) ListRestaurants(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRestaurants", reflect.TypeOf((*MockRestaurantServiceClient)(nil).ListRestaurants), varargs...)
}

// ListRestaurantsNonStream mocks base method.
func (m *MockRestaurantServiceClient) ListRestaurantsNonStream(ctx context.Context, in *ListRestaurantsRequest, opts ...grpc.CallOption) (*ListRestaurantsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRestaurantsNonStream", varargs...)
	ret0, _ := ret[0].(*ListRestaurantsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRestaurantsNonStream indicates an expected call of ListRestaurantsNonStream.
func (mr *MockRestaurantServiceClientMockRecorder) ListRestaurantsNonStream(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRestaurantsNonStream", reflect.TypeOf((*MockRestaurantServiceClient)(nil).ListRestaurantsNonStream), varargs...)
}

// MockRestaurantServiceServer is a mock of RestaurantServiceServer interface.
type MockRestaurantServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockRestaurantServiceServerMockRecorder
}

// MockRestaurantServiceServerMockRecorder is the mock recorder for MockRestaurantServiceServer.
type MockRestaurantServiceServerMockRecorder struct {
	mock *MockRestaurantServiceServer
}

// NewMockRestaurantServiceServer creates a new mock instance.
func NewMockRestaurantServiceServer(ctrl *gomock.Controller) *MockRestaurantServiceServer {
	mock := &MockRestaurantServiceServer{ctrl: ctrl}
	mock.recorder = &MockRestaurantServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRestaurantServiceServer) EXPECT() *MockRestaurantServiceServerMockRecorder {
	return m.recorder
}

// CreateRestaurant mocks base method.
func (m *MockRestaurantServiceServer) CreateRestaurant(ctx context.Context, in *CreateRestaurantRequest) (*model.Restaurant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRestaurant", ctx, in)
	ret0, _ := ret[0].(*model.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRestaurant indicates an expected call of CreateRestaurant.
func (mr *MockRestaurantServiceServerMockRecorder) CreateRestaurant(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRestaurant", reflect.TypeOf((*MockRestaurantServiceServer)(nil).CreateRestaurant), ctx, in)
}

// GetRestaurant mocks base method.
func (m *MockRestaurantServiceServer) GetRestaurant(ctx context.Context, in *GetRestaurantRequest) (*model.Restaurant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurant", ctx, in)
	ret0, _ := ret[0].(*model.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurant indicates an expected call of GetRestaurant.
func (mr *MockRestaurantServiceServerMockRecorder) GetRestaurant(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurant", reflect.TypeOf((*MockRestaurantServiceServer)(nil).GetRestaurant), ctx, in)
}

// ListRestaurants mocks base method.
func (m *MockRestaurantServiceServer) ListRestaurants(blob *ListRestaurantsRequest, server RestaurantService_ListRestaurantsServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRestaurants", blob, server)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListRestaurants indicates an expected call of ListRestaurants.
func (mr *MockRestaurantServiceServerMockRecorder) ListRestaurants(blob, server interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRestaurants", reflect.TypeOf((*MockRestaurantServiceServer)(nil).ListRestaurants), blob, server)
}

// ListRestaurantsNonStream mocks base method.
func (m *MockRestaurantServiceServer) ListRestaurantsNonStream(ctx context.Context, in *ListRestaurantsRequest) (*ListRestaurantsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRestaurantsNonStream", ctx, in)
	ret0, _ := ret[0].(*ListRestaurantsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRestaurantsNonStream indicates an expected call of ListRestaurantsNonStream.
func (mr *MockRestaurantServiceServerMockRecorder) ListRestaurantsNonStream(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRestaurantsNonStream", reflect.TypeOf((*MockRestaurantServiceServer)(nil).ListRestaurantsNonStream), ctx, in)
}
