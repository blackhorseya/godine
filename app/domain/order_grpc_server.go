package domain

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
	payM "github.com/blackhorseya/godine/entity/domain/payment/model"
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

//nolint:funlen // it's okay
func (i *orderService) SubmitOrder(c context.Context, req *biz.SubmitOrderRequest) (*model.Order, error) {
	next, span := otelx.Tracer.Start(c, "order.biz.SubmitOrder")
	defer span.End()

	ctx := contextx.WithContextx(c)

	// check if the user is logged in
	handler, err := userM.FromContext(c)
	if err != nil {
		ctx.Error("failed to get user from context", zap.Error(err))
		return nil, err
	}

	// check restaurant is open
	restaurant, err := i.restaurantClient.GetRestaurant(next, &restB.GetRestaurantRequest{RestaurantId: req.RestaurantId})
	if err != nil {
		ctx.Error("failed to get restaurant", zap.Error(err), zap.String("restaurant_id", req.RestaurantId))
		return nil, err
	}
	if !restaurant.IsOpen {
		ctx.Error("restaurant is not open", zap.String("restaurant_id", restaurant.Id))
		return nil, fmt.Errorf("restaurant %s is not open", req.RestaurantId)
	}

	// check menu is available and collect order items
	var orderItems []*model.OrderItem
	for _, item := range req.Items {
		menuItem, err2 := i.menuClient.GetMenuItem(next, &restB.GetMenuItemRequest{
			RestaurantId: restaurant.Id,
			MenuItemId:   item.MenuItemId,
		})
		if err2 != nil {
			ctx.Error("failed to get menu item", zap.Error(err2))
			return nil, err2
		}

		if !menuItem.IsAvailable {
			ctx.Error("menu item is not available", zap.Any("menu_item", menuItem))
			return nil, fmt.Errorf("menu item %s is not available", item.MenuItemId)
		}

		orderItems = append(orderItems, model.NewOrderItem(menuItem.Id, menuItem.Price, int(item.Quantity)))
	}

	// new order with the user, order items
	order := model.NewOrder(handler.Id, restaurant.Id, orderItems)

	// store the order
	err = i.orders.Create(next, order)
	if err != nil {
		ctx.Error("failed to create order", zap.Error(err))
		return nil, err
	}

	payment, err := i.paymentClient.CreatePayment(next, &payB.CreatePaymentRequest{
		OrderId: order.Id,
		Amount: &payM.PaymentAmount{
			Value:    order.TotalAmount,
			Currency: "USD",
		},
	})
	if err != nil {
		ctx.Error("failed to create payment", zap.Error(err))
		return nil, err
	}
	order.PaymentId = payment.Id

	err = i.orders.Update(next, order)
	if err != nil {
		ctx.Error("failed to update order", zap.Error(err))
		return nil, err
	}

	// book the delivery
	delivery, err := i.logisticsClient.CreateDelivery(next, &opsB.CreateDeliveryRequest{
		OrderId: order.Id,
		UserId:  handler.Id,
		Address: req.Address,
		Phone:   "",
		Note:    "",
	})
	if err != nil {
		ctx.Error("failed to create delivery", zap.Error(err))
		return nil, err
	}
	order.DeliveryId = delivery.Id

	err = i.orders.Update(next, order)
	if err != nil {
		ctx.Error("failed to update order", zap.Error(err))
		return nil, err
	}

	// send notification
	_, err = i.notifyClient.SendNotification(next, &notifyB.SendNotificationRequest{
		UserId:  handler.Id,
		OrderId: order.Id,
		Type:    "",
		Message: fmt.Sprintf("order %v is submitted", order.Id),
	})
	if err != nil {
		ctx.Error("failed to send notification", zap.Error(err))
		return nil, err
	}

	return order, nil
}

func (i *orderService) ListOrders(req *biz.ListOrdersRequest, stream biz.OrderService_ListOrdersServer) error {
	c := stream.Context()
	next, span := otelx.Tracer.Start(c, "order.biz.ListOrders")
	defer span.End()

	ctx := contextx.WithContextx(c)

	items, total, err := i.orders.List(next, repo.ListCondition{
		UserID:       "",
		RestaurantID: "",
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

func (i *orderService) GetOrder(c context.Context, req *biz.GetOrderRequest) (*model.Order, error) {
	next, span := otelx.Tracer.Start(c, "order.biz.GetOrder")
	defer span.End()

	ctx := contextx.WithContextx(c)

	item, err := i.orders.GetByID(next, req.OrderId)
	if err != nil {
		ctx.Error("failed to get order", zap.Error(err))
		return nil, err
	}

	return item, nil
}
