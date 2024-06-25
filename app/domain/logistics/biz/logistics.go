package biz

import (
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/logistics/biz"
	"github.com/blackhorseya/godine/entity/logistics/model"
	"github.com/blackhorseya/godine/entity/logistics/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
)

type logistics struct {
	deliveries repo.IDeliveryRepo
}

// NewLogistics will create a new logistics biz
func NewLogistics(deliveries repo.IDeliveryRepo) biz.ILogisticsBiz {
	return &logistics{
		deliveries: deliveries,
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

	delivery.Status = status

	return i.deliveries.Update(ctx, delivery)
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
