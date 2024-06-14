// Code generated by MockGen. DO NOT EDIT.
// Source: user_service.go
//
// Generated by this command:
//
//	mockgen -destination=./mock_user_service.go -package=biz -source=user_service.go
//

// Package biz is a generated GoMock package.
package biz

import (
	reflect "reflect"

	model "github.com/blackhorseya/godine/entity/user/model"
	contextx "github.com/blackhorseya/godine/pkg/contextx"
	gomock "go.uber.org/mock/gomock"
)

// MockIUserBiz is a mock of IUserBiz interface.
type MockIUserBiz struct {
	ctrl     *gomock.Controller
	recorder *MockIUserBizMockRecorder
}

// MockIUserBizMockRecorder is the mock recorder for MockIUserBiz.
type MockIUserBizMockRecorder struct {
	mock *MockIUserBiz
}

// NewMockIUserBiz creates a new mock instance.
func NewMockIUserBiz(ctrl *gomock.Controller) *MockIUserBiz {
	mock := &MockIUserBiz{ctrl: ctrl}
	mock.recorder = &MockIUserBizMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserBiz) EXPECT() *MockIUserBizMockRecorder {
	return m.recorder
}

// ChangeUserStatus mocks base method.
func (m *MockIUserBiz) ChangeUserStatus(ctx contextx.Contextx, userID string, isActive bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUserStatus", ctx, userID, isActive)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeUserStatus indicates an expected call of ChangeUserStatus.
func (mr *MockIUserBizMockRecorder) ChangeUserStatus(ctx, userID, isActive any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserStatus", reflect.TypeOf((*MockIUserBiz)(nil).ChangeUserStatus), ctx, userID, isActive)
}

// CreateUser mocks base method.
func (m *MockIUserBiz) CreateUser(ctx contextx.Contextx, name, email, password string, address model.Address) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, name, email, password, address)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIUserBizMockRecorder) CreateUser(ctx, name, email, password, address any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIUserBiz)(nil).CreateUser), ctx, name, email, password, address)
}

// DeleteUser mocks base method.
func (m *MockIUserBiz) DeleteUser(ctx contextx.Contextx, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockIUserBizMockRecorder) DeleteUser(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockIUserBiz)(nil).DeleteUser), ctx, id)
}

// GetUser mocks base method.
func (m *MockIUserBiz) GetUser(ctx contextx.Contextx, id string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockIUserBizMockRecorder) GetUser(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockIUserBiz)(nil).GetUser), ctx, id)
}

// ListUsers mocks base method.
func (m *MockIUserBiz) ListUsers(ctx contextx.Contextx, options ListUsersOptions) ([]model.User, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", ctx, options)
	ret0, _ := ret[0].([]model.User)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockIUserBizMockRecorder) ListUsers(ctx, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockIUserBiz)(nil).ListUsers), ctx, options)
}

// SearchUsers mocks base method.
func (m *MockIUserBiz) SearchUsers(ctx contextx.Contextx, keyword string) ([]model.User, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchUsers", ctx, keyword)
	ret0, _ := ret[0].([]model.User)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SearchUsers indicates an expected call of SearchUsers.
func (mr *MockIUserBizMockRecorder) SearchUsers(ctx, keyword any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUsers", reflect.TypeOf((*MockIUserBiz)(nil).SearchUsers), ctx, keyword)
}

// UpdateUser mocks base method.
func (m *MockIUserBiz) UpdateUser(ctx contextx.Contextx, id, name, email, password string, address model.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, id, name, email, password, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockIUserBizMockRecorder) UpdateUser(ctx, id, name, email, password, address any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockIUserBiz)(nil).UpdateUser), ctx, id, name, email, password, address)
}
