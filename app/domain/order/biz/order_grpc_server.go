package biz

import (
	"context"
	"fmt"

	"github.com/blackhorseya/godine/app/infra/otelx"
	opsB "github.com/blackhorseya/godine/entity/domain/logistics/biz"
	notifyB "github.com/blackhorseya/godine/entity/domain/notification/biz"
	"github.com/blackhorseya/godine/entity/domain/order/biz"
	"github.com/blackhorseya/godine/entity/domain/order/model"
	"github.com/blackhorseya/godine/entity/domain/order/repo"
	payB "github.com/blackhorseya/godine/entity/domain/payment/biz"
	restB "github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	userB "github.com/blackhorseya/godine/entity/domain/user/biz"
	userM "github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type orderService struct {
	orders repo.IOrderRepo

	// clients
	restaurantClient restB.RestaurantServiceClient
	menuClient       restB.MenuServiceClient
	accountClient    userB.AccountServiceClient
	notifyClient     notifyB.NotificationServiceClient
	paymentClient    payB.PaymentServiceClient
	logisticsClient  opsB.LogisticsServiceClient
}

// NewOrderService returns the order service instance.
func NewOrderService(
	orders repo.IOrderRepo,
	restaurantClient restB.RestaurantServiceClient,
	menuClient restB.MenuServiceClient,
	accountClient userB.AccountServiceClient,
	notifyClient notifyB.NotificationServiceClient,
	paymentClient payB.PaymentServiceClient,
	logisticsClient opsB.LogisticsServiceClient,
) biz.OrderServiceServer {
	return &orderService{
		orders:           orders,
		restaurantClient: restaurantClient,
		menuClient:       menuClient,
		accountClient:    accountClient,
		notifyClient:     notifyClient,
		paymentClient:    paymentClient,
		logisticsClient:  logisticsClient,
	}
}

func (i *orderService) SubmitOrder(c context.Context, req *biz.SubmitOrderRequest) (*model.Order, error) {
	ctx, err := contextx.FromContext(c)
	if err != nil {
		return nil, fmt.Errorf("failed to get contextx: %w", err)
	}

	ctx, span := otelx.Span(ctx, "order.biz.SubmitOrder")
	defer span.End()

	// check if the user is logged in
	handler, err := userM.FromContext(ctx)
	if err != nil {
		ctx.Error("failed to get user from context", zap.Error(err))
		return nil, err
	}
	_ = handler

	// check restaurant is open

	// TODO: 2024/8/21|sean|implement me
	panic("implement me")
}

func (i *orderService) ListOrders(req *biz.ListOrdersRequest, stream biz.OrderService_ListOrdersServer) error {
	ctx, err := contextx.FromContext(stream.Context())
	if err != nil {
		return fmt.Errorf("failed to get contextx: %w", err)
	}

	ctx, span := otelx.Span(ctx, "order.biz.ListOrders")
	defer span.End()

	items, total, err := i.orders.List(ctx, repo.ListCondition{
		UserID:       "",
		RestaurantID: "",
		Status:       "",
		Limit:        int(req.PageSize),
		Offset:       int((req.Page - 1) * req.PageSize),
	})
	if err != nil {
		ctx.Error("failed to list orders", zap.Error(err))
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
			ctx.Error("failed to send order", zap.Error(err))
			return err
		}
	}

	return nil
}
