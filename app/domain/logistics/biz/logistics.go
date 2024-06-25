package biz

import (
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
	// todo: 2024/6/25|sean|implement me
	panic("implement me")
}

func (i *logistics) UpdateDeliveryStatus(ctx contextx.Contextx, deliveryID string, status string) error {
	// todo: 2024/6/25|sean|implement me
	panic("implement me")
}

func (i *logistics) GetDelivery(ctx contextx.Contextx, deliveryID string) (item *model.Delivery, err error) {
	// todo: 2024/6/25|sean|implement me
	panic("implement me")
}

func (i *logistics) ListDeliveriesByDriver(
	ctx contextx.Contextx,
	driverID string,
	options biz.ListDeliveriesOptions,
) (items []*model.Delivery, total int, err error) {
	// todo: 2024/6/25|sean|implement me
	panic("implement me")
}
