package biz

import (
	"context"
	"fmt"

	"github.com/blackhorseya/godine/app/infra/otelx"
	notifyB "github.com/blackhorseya/godine/entity/domain/notification/biz"
	"github.com/blackhorseya/godine/entity/domain/order/biz"
	"github.com/blackhorseya/godine/entity/domain/order/model"
	"github.com/blackhorseya/godine/entity/domain/order/repo"
	payB "github.com/blackhorseya/godine/entity/domain/payment/biz"
	restB "github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	userB "github.com/blackhorseya/godine/entity/domain/user/biz"
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
}

// NewOrderService returns the order service instance.
func NewOrderService(
	orders repo.IOrderRepo,
	restaurantClient restB.RestaurantServiceClient,
	menuClient restB.MenuServiceClient,
	accountClient userB.AccountServiceClient,
	notifyClient notifyB.NotificationServiceClient,
	paymentClient payB.PaymentServiceClient,
) biz.OrderServiceServer {
	return &orderService{
		orders:           orders,
		restaurantClient: restaurantClient,
		menuClient:       menuClient,
		accountClient:    accountClient,
		notifyClient:     notifyClient,
		paymentClient:    paymentClient,
	}
}

func (i *orderService) SubmitOrder(c context.Context, req *biz.SubmitOrderRequest) (*model.Order, error) {
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
