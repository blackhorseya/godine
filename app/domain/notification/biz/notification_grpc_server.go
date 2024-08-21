package biz

import (
	"context"
	"fmt"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/notification/biz"
	"github.com/blackhorseya/godine/entity/domain/notification/model"
	"github.com/blackhorseya/godine/entity/domain/notification/repo"
	userM "github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.uber.org/zap"
)

type notificationService struct {
	notifications repo.INotificationRepo
}

// NewNotificationService creates a new notification service.
func NewNotificationService(notifications repo.INotificationRepo) biz.NotificationServiceServer {
	return &notificationService{
		notifications: notifications,
	}
}

func (i *notificationService) SendNotification(
	c context.Context,
	req *biz.SendNotificationRequest,
) (*model.Notification, error) {
	ctx, err := contextx.FromContext(c)
	if err != nil {
		return nil, fmt.Errorf("failed to get contextx: %w", err)
	}

	ctx, span := otelx.Span(ctx, "notification.biz.SendNotification")
	defer span.End()

	handler, err := userM.FromContext(ctx)
	if err != nil {
		ctx.Error("failed to get user from context", zap.Error(err))
		return nil, err
	}

	notification := model.NewNotification(handler.Id, req.UserId, req.OrderId, req.Message)

	err = i.notifications.Create(ctx, notification)
	if err != nil {
		ctx.Error("create notification failed", zap.Error(err), zap.Any("notification", notification))
		return nil, err
	}

	return notification, nil
}

func (i *notificationService) ListMyNotifications(
	req *biz.ListMyNotificationsRequest,
	stream biz.NotificationService_ListMyNotificationsServer,
) error {
	// TODO: 2024/8/21|sean|implement me
	panic("implement me")
}
