package notification

import (
	"context"
	"fmt"
	"strconv"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/notification/biz"
	"github.com/blackhorseya/godine/entity/domain/notification/model"
	"github.com/blackhorseya/godine/entity/domain/notification/repo"
	userM "github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
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
	ctx, err := contextx.FromContextLegacy(c)
	if err != nil {
		return nil, fmt.Errorf("failed to get contextx: %w", err)
	}

	ctx, span := otelx.Span(ctx, "notification.biz.SendNotification")
	defer span.End()

	handler, err := userM.FromContextLegacy(ctx)
	if err != nil {
		ctx.Error("failed to get user from context", zap.Error(err))
		return nil, err
	}

	notification := model.NewNotification(handler.Id, req.UserId, strconv.FormatInt(req.OrderId, 10), req.Message)

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
	ctx, err := contextx.FromContextLegacy(stream.Context())
	if err != nil {
		return fmt.Errorf("failed to get contextx: %w", err)
	}

	ctx, span := otelx.Span(ctx, "notification.biz.ListMyNotifications")
	defer span.End()

	handler, err := userM.FromContextLegacy(ctx)
	if err != nil {
		ctx.Error("failed to get user from context", zap.Error(err))
		return err
	}
	_ = handler

	items, total, err := i.notifications.List(ctx, utils.ListCondition{
		Limit:  req.PageSize,
		Offset: (req.Page - 1) * req.PageSize,
		// TODO: 2024/8/31|sean|fix me
		// UserID: handler.Id,
	})
	if err != nil {
		ctx.Error("list notifications failed", zap.Error(err))
		return err
	}

	err = stream.SetHeader(metadata.New(map[string]string{"total": strconv.Itoa(total)}))
	if err != nil {
		ctx.Error("set header failed", zap.Error(err))
		return err
	}

	for _, item := range items {
		err = stream.Send(item)
		if err != nil {
			ctx.Error("send notification failed", zap.Error(err))
			return err
		}
	}

	return nil
}
