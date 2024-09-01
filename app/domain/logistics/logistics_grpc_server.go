package logistics

import (
	"context"
	"fmt"
	"strconv"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/logistics/biz"
	"github.com/blackhorseya/godine/entity/domain/logistics/model"
	"github.com/blackhorseya/godine/entity/domain/logistics/repo"
	notifyB "github.com/blackhorseya/godine/entity/domain/notification/biz"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
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
	ctx, err := contextx.FromContextLegacy(c)
	if err != nil {
		return nil, fmt.Errorf("failed to get contextx: %w", err)
	}

	next, span := otelx.Tracer.Start(c, "biz.logistics.CreateDelivery")
	defer span.End()

	delivery, err := model.NewDelivery(strconv.FormatInt(req.OrderId, 10), req.UserId, req.Address)
	if err != nil {
		ctx.Error("failed to create new delivery", zap.Error(err))
		return nil, err
	}

	err = i.deliveries.Create(next, delivery)
	if err != nil {
		ctx.Error("failed to create delivery", zap.Error(err))
		return nil, err
	}

	return delivery, nil
}

func (i *logisticsService) ListDeliveries(
	req *biz.ListDeliveriesRequest,
	stream biz.LogisticsService_ListDeliveriesServer,
) error {
	ctx, err := contextx.FromContextLegacy(stream.Context())
	if err != nil {
		return fmt.Errorf("failed to get contextx: %w", err)
	}

	next, span := otelx.Tracer.Start(stream.Context(), "biz.logistics.ListDeliveries")
	defer span.End()

	items, total, err := i.deliveries.List(next, utils.Pagination{
		Limit:  req.PageSize,
		Offset: (req.Page - 1) * req.PageSize,
	})
	if err != nil {
		ctx.Error("failed to list deliveries", zap.Error(err))
		return err
	}

	err = stream.SetHeader(metadata.New(map[string]string{"total": fmt.Sprintf("%d", total)}))
	if err != nil {
		ctx.Error("failed to set header", zap.Error(err))
		return err
	}

	for _, item := range items {
		err = stream.Send(item)
		if err != nil {
			ctx.Error("failed to send delivery", zap.Error(err))
			return err
		}
	}

	return nil
}
