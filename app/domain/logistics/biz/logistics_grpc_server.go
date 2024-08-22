package biz

import (
	"context"

	"github.com/blackhorseya/godine/entity/domain/logistics/biz"
	"github.com/blackhorseya/godine/entity/domain/logistics/model"
	"github.com/blackhorseya/godine/entity/domain/logistics/repo"
	notifyB "github.com/blackhorseya/godine/entity/domain/notification/biz"
)

type logisticsService struct {
	deliveries repo.IDeliveryRepo

	// clients
	notifyClient notifyB.NotificationServiceClient
}

// NewLogisticsService creates a new logistics service.
func NewLogisticsService(
	deliveries repo.IDeliveryRepo,
	notifyClient notifyB.NotificationServiceClient,
) biz.LogisticsServiceServer {
	return &logisticsService{
		deliveries:   deliveries,
		notifyClient: notifyClient,
	}
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
