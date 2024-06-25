package biz

import (
	"github.com/blackhorseya/godine/entity/notification/biz"
	"github.com/blackhorseya/godine/entity/notification/model"
	"github.com/blackhorseya/godine/pkg/contextx"
)

type notificationHTTPClient struct {
}

// NewNotificationHTTPClient creates a new notification service.
func NewNotificationHTTPClient() biz.INotificationBiz {
	return &notificationHTTPClient{}
}

func (i *notificationHTTPClient) CreateNotification(ctx contextx.Contextx, notification *model.Notification) error {
	// todo: 2024/6/26|sean|implement me
	panic("implement me")
}

func (i *notificationHTTPClient) UpdateNotificationStatus(
	ctx contextx.Contextx,
	notificationID string,
	status string,
) error {
	// todo: 2024/6/26|sean|implement me
	panic("implement me")
}

func (i *notificationHTTPClient) GetNotification(
	ctx contextx.Contextx,
	notificationID string,
) (item *model.Notification, err error) {
	// todo: 2024/6/26|sean|implement me
	panic("implement me")
}

func (i *notificationHTTPClient) ListNotificationsByUser(
	ctx contextx.Contextx,
	userID string,
	options biz.ListNotificationsOptions,
) (items []*model.Notification, total int, err error) {
	// todo: 2024/6/26|sean|implement me
	panic("implement me")
}
