package biz

import (
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/notification/biz"
	"github.com/blackhorseya/godine/entity/domain/notification/model"
	"github.com/blackhorseya/godine/entity/domain/notification/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
)

type impl struct {
	notifications repo.INotificationRepo
}

// NewNotification creates a new impl service.
func NewNotification(notifications repo.INotificationRepo) biz.INotificationBiz {
	return &impl{
		notifications: notifications,
	}
}

func (i *impl) CreateNotification(ctx contextx.Contextx, notification *model.Notification) error {
	ctx, span := otelx.Span(ctx, "biz.impl.CreateNotification")
	defer span.End()

	return i.notifications.Create(ctx, notification)
}

func (i *impl) GetNotification(
	ctx contextx.Contextx,
	notificationID string,
) (item *model.Notification, err error) {
	ctx, span := otelx.Span(ctx, "biz.impl.GetNotification")
	defer span.End()

	return i.notifications.GetByID(ctx, notificationID)
}

func (i *impl) ListNotificationsByUser(
	ctx contextx.Contextx,
	userID string,
	options biz.ListNotificationsOptions,
) (items []*model.Notification, total int, err error) {
	ctx, span := otelx.Span(ctx, "biz.impl.ListNotificationsByUser")
	defer span.End()

	return i.notifications.List(ctx, repo.ListCondition{
		Limit:  options.Size,
		Offset: (options.Page - 1) * options.Size,
		UserID: userID,
	})
}
