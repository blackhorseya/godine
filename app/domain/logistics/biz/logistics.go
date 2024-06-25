package biz

import (
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/logistics/biz"
	"github.com/blackhorseya/godine/entity/logistics/model"
	"github.com/blackhorseya/godine/entity/logistics/repo"
	notifyB "github.com/blackhorseya/godine/entity/notification/biz"
	model2 "github.com/blackhorseya/godine/entity/notification/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.uber.org/zap"
)

type logistics struct {
	notifyService notifyB.INotificationBiz

	deliveries repo.IDeliveryRepo
}

// NewLogistics will create a new logistics biz
func NewLogistics(notifyService notifyB.INotificationBiz, deliveries repo.IDeliveryRepo) biz.ILogisticsBiz {
	return &logistics{
		notifyService: notifyService,
		deliveries:    deliveries,
	}
}

func (i *logistics) CreateDelivery(ctx contextx.Contextx, delivery *model.Delivery) error {
	ctx, span := otelx.Span(ctx, "biz.logistics.CreateDelivery")
	defer span.End()

	return i.deliveries.Create(ctx, delivery)
}

func (i *logistics) UpdateDeliveryStatus(ctx contextx.Contextx, deliveryID string, status string) error {
	ctx, span := otelx.Span(ctx, "biz.logistics.UpdateDeliveryStatus")
	defer span.End()

	delivery, err := i.deliveries.GetByID(ctx, deliveryID)
	if err != nil {
		return err
	}

	event, err := delivery.Next(ctx)
	if err != nil {
		return err
	}

	ctx.Debug("delivery next event", zap.Any("event", &event))

	err = i.deliveries.Update(ctx, delivery)
	if err != nil {
		return err
	}

	err = i.notifyService.CreateNotification(ctx, model2.NewNotify(
		delivery.DriverID,
		delivery.ID,
		delivery.OrderID,
		"delivery status changed",
	))
	if err != nil {
		return err
	}

	return nil
}

func (i *logistics) GetDelivery(ctx contextx.Contextx, deliveryID string) (item *model.Delivery, err error) {
	ctx, span := otelx.Span(ctx, "biz.logistics.GetDelivery")
	defer span.End()

	return i.deliveries.GetByID(ctx, deliveryID)
}

func (i *logistics) ListDeliveriesByDriver(
	ctx contextx.Contextx,
	driverID string,
	options biz.ListDeliveriesOptions,
) (items []*model.Delivery, total int, err error) {
	ctx, span := otelx.Span(ctx, "biz.logistics.ListDeliveriesByDriver")
	defer span.End()

	return i.deliveries.List(ctx, repo.ListCondition{
		Limit:    options.Size,
		Offset:   (options.Page - 1) * options.Size,
		DriverID: driverID,
	})
}
