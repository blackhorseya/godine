package biz

import (
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/notification/biz"
	"github.com/blackhorseya/godine/entity/domain/notification/model"
	"github.com/blackhorseya/godine/entity/domain/notification/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
)

type notification struct {
	notifications repo.INotificationRepo
}

// NewNotification creates a new notification service.
func NewNotification(notifications repo.INotificationRepo) biz.INotificationBiz {
	return &notification{
		notifications: notifications,
	}
}

func (i *notification) CreateNotification(ctx contextx.Contextx, notification *model.Notification) error {
	ctx, span := otelx.Span(ctx, "biz.notification.CreateNotification")
	defer span.End()

	return i.notifications.Create(ctx, notification)
}

func (i *notification) UpdateNotificationStatus(ctx contextx.Contextx, notificationID string, status string) error {
	ctx, span := otelx.Span(ctx, "biz.notification.UpdateNotificationStatus")
	defer span.End()

	return i.notifications.UpdateStatus(ctx, notificationID, status)
}

func (i *notification) GetNotification(
	ctx contextx.Contextx,
	notificationID string,
) (item *model.Notification, err error) {
	ctx, span := otelx.Span(ctx, "biz.notification.GetNotification")
	defer span.End()

	return i.notifications.GetByID(ctx, notificationID)
}

func (i *notification) ListNotificationsByUser(
	ctx contextx.Contextx,
	userID string,
	options biz.ListNotificationsOptions,
) (items []*model.Notification, total int, err error) {
	ctx, span := otelx.Span(ctx, "biz.notification.ListNotificationsByUser")
	defer span.End()

	return i.notifications.List(ctx, repo.ListCondition{
		Limit:  options.Size,
		Offset: (options.Page - 1) * options.Size,
		UserID: userID,
	})
}
