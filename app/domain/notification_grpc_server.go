package domain

import (
	"context"
	"strconv"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/notification/biz"
	"github.com/blackhorseya/godine/entity/domain/notification/model"
	"github.com/blackhorseya/godine/entity/domain/notification/repo"
	userM "github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/persistence"
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
	next, span := otelx.Tracer.Start(c, "notification.biz.SendNotification")
	defer span.End()

	ctx := contextx.WithContextx(c)

	handler, err := userM.FromContext(c)
	if err != nil {
		ctx.Error("failed to get user from context", zap.Error(err))
		return nil, err
	}

	notification := model.NewNotification(handler.Id, req.UserId, req.OrderId, req.Message)

	err = i.notifications.Create(next, notification)
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
	c := stream.Context()
	next, span := otelx.Tracer.Start(c, "notification.biz.ListMyNotifications")
	defer span.End()

	ctx := contextx.WithContextx(c)

	handler, err := userM.FromContext(stream.Context())
	if err != nil {
		ctx.Error("failed to get user from context", zap.Error(err))
		return err
	}

	items, total, err := i.notifications.ListByReceiverID(next, handler.Id, persistence.Pagination{
		Limit:  req.PageSize,
		Offset: (req.Page - 1) * req.PageSize,
	})
	if err != nil {
		ctx.Error("list notifications failed", zap.Error(err))
		return err
	}

	err = stream.SetHeader(metadata.New(map[string]string{"total": strconv.FormatInt(total, 10)}))
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
