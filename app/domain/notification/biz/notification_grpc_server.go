package biz

import (
	"context"

	"github.com/blackhorseya/godine/entity/domain/notification/biz"
	"github.com/blackhorseya/godine/entity/domain/notification/model"
)

type notificationService struct {
}

// NewNotificationService creates a new notification service.
func NewNotificationService() biz.NotificationServiceServer {
	return &notificationService{}
}

func (i *notificationService) SendNotification(
	c context.Context,
	req *biz.SendNotificationRequest,
) (*model.Notification, error) {
	// TODO: 2024/8/21|sean|implement me
	panic("implement me")
}

func (i *notificationService) ListMyNotifications(
	req *biz.ListMyNotificationsRequest,
	stream biz.NotificationService_ListMyNotificationsServer,
) error {
	// TODO: 2024/8/21|sean|implement me
	panic("implement me")
}
