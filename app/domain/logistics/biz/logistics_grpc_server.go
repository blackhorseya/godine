package biz

import (
	"context"

	"github.com/blackhorseya/godine/entity/domain/logistics/biz"
	"github.com/blackhorseya/godine/entity/domain/logistics/model"
)

type logisticsService struct {
}

// NewLogisticsService creates a new logistics service.
func NewLogisticsService() biz.LogisticsServiceServer {
	return &logisticsService{}
}

func (i *logisticsService) CreateDelivery(c context.Context, req *biz.CreateDeliveryRequest) (*model.Delivery, error) {
	// TODO: 2024/8/22|sean|implement me
	panic("implement me")
}

func (i *logisticsService) ListDeliveries(
	req *biz.ListDeliveriesRequest,
	stream biz.LogisticsService_ListDeliveriesServer,
) error {
	// TODO: 2024/8/22|sean|implement me
	panic("implement me")
}
