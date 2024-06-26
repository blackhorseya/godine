// Code generated by MockGen. DO NOT EDIT.
// Source: logistics_service.go
//
// Generated by this command:
//
//	mockgen -destination=./mock_logistics_service.go -package=biz -source=logistics_service.go
//

// Package biz is a generated GoMock package.
package biz

import (
	"reflect"

	"github.com/blackhorseya/godine/entity/domain/logistics/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.uber.org/mock/gomock"
)

// MockILogisticsBiz is a mock of ILogisticsBiz interface.
type MockILogisticsBiz struct {
	ctrl     *gomock.Controller
	recorder *MockILogisticsBizMockRecorder
}

// MockILogisticsBizMockRecorder is the mock recorder for MockILogisticsBiz.
type MockILogisticsBizMockRecorder struct {
	mock *MockILogisticsBiz
}

// NewMockILogisticsBiz creates a new mock instance.
func NewMockILogisticsBiz(ctrl *gomock.Controller) *MockILogisticsBiz {
	mock := &MockILogisticsBiz{ctrl: ctrl}
	mock.recorder = &MockILogisticsBizMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockILogisticsBiz) EXPECT() *MockILogisticsBizMockRecorder {
	return m.recorder
}

// CreateDelivery mocks base method.
func (m *MockILogisticsBiz) CreateDelivery(ctx contextx.Contextx, delivery *model.Delivery) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDelivery", ctx, delivery)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDelivery indicates an expected call of CreateDelivery.
func (mr *MockILogisticsBizMockRecorder) CreateDelivery(ctx, delivery any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDelivery", reflect.TypeOf((*MockILogisticsBiz)(nil).CreateDelivery), ctx, delivery)
}

// GetDelivery mocks base method.
func (m *MockILogisticsBiz) GetDelivery(ctx contextx.Contextx, deliveryID string) (*model.Delivery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDelivery", ctx, deliveryID)
	ret0, _ := ret[0].(*model.Delivery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDelivery indicates an expected call of GetDelivery.
func (mr *MockILogisticsBizMockRecorder) GetDelivery(ctx, deliveryID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDelivery", reflect.TypeOf((*MockILogisticsBiz)(nil).GetDelivery), ctx, deliveryID)
}

// ListDeliveriesByDriver mocks base method.
func (m *MockILogisticsBiz) ListDeliveriesByDriver(ctx contextx.Contextx, driverID string, options ListDeliveriesOptions) ([]*model.Delivery, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeliveriesByDriver", ctx, driverID, options)
	ret0, _ := ret[0].([]*model.Delivery)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListDeliveriesByDriver indicates an expected call of ListDeliveriesByDriver.
func (mr *MockILogisticsBizMockRecorder) ListDeliveriesByDriver(ctx, driverID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeliveriesByDriver", reflect.TypeOf((*MockILogisticsBiz)(nil).ListDeliveriesByDriver), ctx, driverID, options)
}

// UpdateDeliveryStatus mocks base method.
func (m *MockILogisticsBiz) UpdateDeliveryStatus(ctx contextx.Contextx, deliveryID, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDeliveryStatus", ctx, deliveryID, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDeliveryStatus indicates an expected call of UpdateDeliveryStatus.
func (mr *MockILogisticsBizMockRecorder) UpdateDeliveryStatus(ctx, deliveryID, status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeliveryStatus", reflect.TypeOf((*MockILogisticsBiz)(nil).UpdateDeliveryStatus), ctx, deliveryID, status)
}
