// Code generated by MockGen. DO NOT EDIT.
// Source: order.go
//
// Generated by this command:
//
//	mockgen -destination=./mock_order.go -package=repo -source=order.go
//

// Package repo is a generated GoMock package.
package repo

import (
	reflect "reflect"

	model "github.com/blackhorseya/godine/entity/order/model"
	contextx "github.com/blackhorseya/godine/pkg/contextx"
	gomock "go.uber.org/mock/gomock"
)

// MockIOrderRepo is a mock of IOrderRepo interface.
type MockIOrderRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIOrderRepoMockRecorder
}

// MockIOrderRepoMockRecorder is the mock recorder for MockIOrderRepo.
type MockIOrderRepoMockRecorder struct {
	mock *MockIOrderRepo
}

// NewMockIOrderRepo creates a new mock instance.
func NewMockIOrderRepo(ctrl *gomock.Controller) *MockIOrderRepo {
	mock := &MockIOrderRepo{ctrl: ctrl}
	mock.recorder = &MockIOrderRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIOrderRepo) EXPECT() *MockIOrderRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIOrderRepo) Create(ctx contextx.Contextx, order *model.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIOrderRepoMockRecorder) Create(ctx, order any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIOrderRepo)(nil).Create), ctx, order)
}

// GetByID mocks base method.
func (m *MockIOrderRepo) GetByID(ctx contextx.Contextx, id string) (*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIOrderRepoMockRecorder) GetByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIOrderRepo)(nil).GetByID), ctx, id)
}

// ListByRestaurantID mocks base method.
func (m *MockIOrderRepo) ListByRestaurantID(ctx contextx.Contextx, restaurantID string, condition ListCondition) ([]*model.Order, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByRestaurantID", ctx, restaurantID, condition)
	ret0, _ := ret[0].([]*model.Order)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListByRestaurantID indicates an expected call of ListByRestaurantID.
func (mr *MockIOrderRepoMockRecorder) ListByRestaurantID(ctx, restaurantID, condition any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByRestaurantID", reflect.TypeOf((*MockIOrderRepo)(nil).ListByRestaurantID), ctx, restaurantID, condition)
}

// ListByUserID mocks base method.
func (m *MockIOrderRepo) ListByUserID(ctx contextx.Contextx, userID string, condition ListCondition) ([]*model.Order, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByUserID", ctx, userID, condition)
	ret0, _ := ret[0].([]*model.Order)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListByUserID indicates an expected call of ListByUserID.
func (mr *MockIOrderRepoMockRecorder) ListByUserID(ctx, userID, condition any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByUserID", reflect.TypeOf((*MockIOrderRepo)(nil).ListByUserID), ctx, userID, condition)
}

// UpdateStatus mocks base method.
func (m *MockIOrderRepo) UpdateStatus(ctx contextx.Contextx, order *model.Order, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", ctx, order, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockIOrderRepoMockRecorder) UpdateStatus(ctx, order, status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockIOrderRepo)(nil).UpdateStatus), ctx, order, status)
}
