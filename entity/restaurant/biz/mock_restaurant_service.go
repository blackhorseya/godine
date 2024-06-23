// Code generated by MockGen. DO NOT EDIT.
// Source: restaurant_service.go
//
// Generated by this command:
//
//	mockgen -destination=./mock_restaurant_service.go -package=biz -source=restaurant_service.go
//

// Package biz is a generated GoMock package.
package biz

import (
	reflect "reflect"

	model "github.com/blackhorseya/godine/entity/restaurant/model"
	contextx "github.com/blackhorseya/godine/pkg/contextx"
	gomock "go.uber.org/mock/gomock"
)

// MockIRestaurantBiz is a mock of IRestaurantBiz interface.
type MockIRestaurantBiz struct {
	ctrl     *gomock.Controller
	recorder *MockIRestaurantBizMockRecorder
}

// MockIRestaurantBizMockRecorder is the mock recorder for MockIRestaurantBiz.
type MockIRestaurantBizMockRecorder struct {
	mock *MockIRestaurantBiz
}

// NewMockIRestaurantBiz creates a new mock instance.
func NewMockIRestaurantBiz(ctrl *gomock.Controller) *MockIRestaurantBiz {
	mock := &MockIRestaurantBiz{ctrl: ctrl}
	mock.recorder = &MockIRestaurantBizMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRestaurantBiz) EXPECT() *MockIRestaurantBizMockRecorder {
	return m.recorder
}

// ChangeRestaurantStatus mocks base method.
func (m *MockIRestaurantBiz) ChangeRestaurantStatus(ctx contextx.Contextx, restaurantID string, isOpen bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeRestaurantStatus", ctx, restaurantID, isOpen)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeRestaurantStatus indicates an expected call of ChangeRestaurantStatus.
func (mr *MockIRestaurantBizMockRecorder) ChangeRestaurantStatus(ctx, restaurantID, isOpen any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeRestaurantStatus", reflect.TypeOf((*MockIRestaurantBiz)(nil).ChangeRestaurantStatus), ctx, restaurantID, isOpen)
}

// CreateRestaurant mocks base method.
func (m *MockIRestaurantBiz) CreateRestaurant(ctx contextx.Contextx, name, address string) (*model.Restaurant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRestaurant", ctx, name, address)
	ret0, _ := ret[0].(*model.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRestaurant indicates an expected call of CreateRestaurant.
func (mr *MockIRestaurantBizMockRecorder) CreateRestaurant(ctx, name, address any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRestaurant", reflect.TypeOf((*MockIRestaurantBiz)(nil).CreateRestaurant), ctx, name, address)
}

// DeleteRestaurant mocks base method.
func (m *MockIRestaurantBiz) DeleteRestaurant(ctx contextx.Contextx, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRestaurant", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRestaurant indicates an expected call of DeleteRestaurant.
func (mr *MockIRestaurantBizMockRecorder) DeleteRestaurant(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRestaurant", reflect.TypeOf((*MockIRestaurantBiz)(nil).DeleteRestaurant), ctx, id)
}

// GetRestaurant mocks base method.
func (m *MockIRestaurantBiz) GetRestaurant(ctx contextx.Contextx, id string) (*model.Restaurant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurant", ctx, id)
	ret0, _ := ret[0].(*model.Restaurant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurant indicates an expected call of GetRestaurant.
func (mr *MockIRestaurantBizMockRecorder) GetRestaurant(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurant", reflect.TypeOf((*MockIRestaurantBiz)(nil).GetRestaurant), ctx, id)
}

// ListRestaurants mocks base method.
func (m *MockIRestaurantBiz) ListRestaurants(ctx contextx.Contextx, options ListRestaurantsOptions) ([]*model.Restaurant, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRestaurants", ctx, options)
	ret0, _ := ret[0].([]*model.Restaurant)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListRestaurants indicates an expected call of ListRestaurants.
func (mr *MockIRestaurantBizMockRecorder) ListRestaurants(ctx, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRestaurants", reflect.TypeOf((*MockIRestaurantBiz)(nil).ListRestaurants), ctx, options)
}

// UpdateRestaurant mocks base method.
func (m *MockIRestaurantBiz) UpdateRestaurant(ctx contextx.Contextx, id, name string, address model.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRestaurant", ctx, id, name, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRestaurant indicates an expected call of UpdateRestaurant.
func (mr *MockIRestaurantBizMockRecorder) UpdateRestaurant(ctx, id, name, address any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRestaurant", reflect.TypeOf((*MockIRestaurantBiz)(nil).UpdateRestaurant), ctx, id, name, address)
}
