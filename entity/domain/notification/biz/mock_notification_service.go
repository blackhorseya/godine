// Code generated by MockGen. DO NOT EDIT.
// Source: notification_service.go
//
// Generated by this command:
//
//	mockgen -destination=./mock_notification_service.go -package=biz -source=notification_service.go
//

// Package biz is a generated GoMock package.
package biz

import (
	reflect "reflect"

	model "github.com/blackhorseya/godine/entity/domain/notification/model"
	contextx "github.com/blackhorseya/godine/pkg/contextx"
	gomock "go.uber.org/mock/gomock"
)

// MockINotificationBiz is a mock of INotificationBiz interface.
type MockINotificationBiz struct {
	ctrl     *gomock.Controller
	recorder *MockINotificationBizMockRecorder
}

// MockINotificationBizMockRecorder is the mock recorder for MockINotificationBiz.
type MockINotificationBizMockRecorder struct {
	mock *MockINotificationBiz
}

// NewMockINotificationBiz creates a new mock instance.
func NewMockINotificationBiz(ctrl *gomock.Controller) *MockINotificationBiz {
	mock := &MockINotificationBiz{ctrl: ctrl}
	mock.recorder = &MockINotificationBizMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockINotificationBiz) EXPECT() *MockINotificationBizMockRecorder {
	return m.recorder
}

// CreateNotification mocks base method.
func (m *MockINotificationBiz) CreateNotification(ctx contextx.Contextx, notification *model.Notification) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNotification", ctx, notification)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNotification indicates an expected call of CreateNotification.
func (mr *MockINotificationBizMockRecorder) CreateNotification(ctx, notification any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNotification", reflect.TypeOf((*MockINotificationBiz)(nil).CreateNotification), ctx, notification)
}

// GetNotification mocks base method.
func (m *MockINotificationBiz) GetNotification(ctx contextx.Contextx, notificationID string) (*model.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotification", ctx, notificationID)
	ret0, _ := ret[0].(*model.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotification indicates an expected call of GetNotification.
func (mr *MockINotificationBizMockRecorder) GetNotification(ctx, notificationID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotification", reflect.TypeOf((*MockINotificationBiz)(nil).GetNotification), ctx, notificationID)
}

// ListNotificationsByUser mocks base method.
func (m *MockINotificationBiz) ListNotificationsByUser(ctx contextx.Contextx, userID string, options ListNotificationsOptions) ([]*model.Notification, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNotificationsByUser", ctx, userID, options)
	ret0, _ := ret[0].([]*model.Notification)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListNotificationsByUser indicates an expected call of ListNotificationsByUser.
func (mr *MockINotificationBizMockRecorder) ListNotificationsByUser(ctx, userID, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNotificationsByUser", reflect.TypeOf((*MockINotificationBiz)(nil).ListNotificationsByUser), ctx, userID, options)
}
