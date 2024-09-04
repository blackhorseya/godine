// Code generated by MockGen. DO NOT EDIT.
// Source: user.go
//
// Generated by this command:
//
//	mockgen -destination=./mock_user.go -package=repo -source=user.go
//

// Package repo is a generated GoMock package.
package repo

import (
	context "context"
	reflect "reflect"

	model "github.com/blackhorseya/godine/entity/domain/user/model"
	utils "github.com/blackhorseya/godine/pkg/persistence"
	gomock "go.uber.org/mock/gomock"
)

// MockIUserRepo is a mock of IUserRepo interface.
type MockIUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepoMockRecorder
}

// MockIUserRepoMockRecorder is the mock recorder for MockIUserRepo.
type MockIUserRepoMockRecorder struct {
	mock *MockIUserRepo
}

// NewMockIUserRepo creates a new mock instance.
func NewMockIUserRepo(ctrl *gomock.Controller) *MockIUserRepo {
	mock := &MockIUserRepo{ctrl: ctrl}
	mock.recorder = &MockIUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepo) EXPECT() *MockIUserRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIUserRepo) Create(c context.Context, item *model.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", c, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIUserRepoMockRecorder) Create(c, item any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUserRepo)(nil).Create), c, item)
}

// Delete mocks base method.
func (m *MockIUserRepo) Delete(c context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", c, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIUserRepoMockRecorder) Delete(c, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIUserRepo)(nil).Delete), c, id)
}

// GetByID mocks base method.
func (m *MockIUserRepo) GetByID(c context.Context, id string) (*model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", c, id)
	ret0, _ := ret[0].(*model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIUserRepoMockRecorder) GetByID(c, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIUserRepo)(nil).GetByID), c, id)
}

// List mocks base method.
func (m *MockIUserRepo) List(c context.Context, cond utils.Pagination) ([]*model.Account, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", c, cond)
	ret0, _ := ret[0].([]*model.Account)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockIUserRepoMockRecorder) List(c, cond any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIUserRepo)(nil).List), c, cond)
}

// Update mocks base method.
func (m *MockIUserRepo) Update(c context.Context, item *model.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", c, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIUserRepoMockRecorder) Update(c, item any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIUserRepo)(nil).Update), c, item)
}
