package biz

import (
	"errors"
	"net/http"

	"github.com/blackhorseya/godine/app/infra/otelx"
	orderB "github.com/blackhorseya/godine/entity/order/biz"
	"github.com/blackhorseya/godine/entity/order/model"
	"github.com/blackhorseya/godine/entity/order/repo"
	restB "github.com/blackhorseya/godine/entity/restaurant/biz"
	userB "github.com/blackhorseya/godine/entity/user/biz"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type orderBiz struct {
	restaurantService restB.IRestaurantBiz
	userService       userB.IUserBiz

	orders repo.IOrderRepo
}

// NewOrderBiz create and return a new order orderB
func NewOrderBiz(
	restaurantService restB.IRestaurantBiz,
	userService userB.IUserBiz,
	orders repo.IOrderRepo,
) orderB.IOrderBiz {
	return &orderBiz{
		restaurantService: restaurantService,
		userService:       userService,
		orders:            orders,
	}
}

func (i *orderBiz) CreateOrder(
	ctx contextx.Contextx,
	userID, restaurantID uuid.UUID,
	items []model.OrderItem,
	address model.Address,
	totalAmount float64,
) (order *model.Order, err error) {
	ctx, span := otelx.Span(ctx, "biz.order.create_order")
	defer span.End()

	restaurant, err := i.restaurantService.GetRestaurant(ctx, restaurantID)
	if err != nil {
		ctx.Error(
			"get restaurant from service failed",
			zap.Error(err),
			zap.String("restaurant_id", restaurantID.String()),
		)
		return nil, err
	}
	if restaurant == nil {
		ctx.Error(
			"restaurant not found",
			zap.String("restaurant_id", restaurantID.String()),
		)
		return nil, errorx.Wrap(http.StatusNotFound, 404, errors.New("restaurant not found"))
	}

	order = model.NewOrder(userID.String(), restaurantID.String(), items, address, totalAmount)
	err = i.orders.Create(ctx, order)
	if err != nil {
		ctx.Error(
			"create order failed",
			zap.Error(err),
			zap.Any("order", &order),
		)
		return nil, err
	}

	return order, nil
}

func (i *orderBiz) GetOrder(ctx contextx.Contextx, id uuid.UUID) (order *model.Order, err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *orderBiz) ListOrders(
	ctx contextx.Contextx,
	options orderB.ListOrdersOptions,
) (orders []model.Order, total int, err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *orderBiz) UpdateOrderStatus(ctx contextx.Contextx, id uuid.UUID, status string) error {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *orderBiz) AddOrderItem(ctx contextx.Contextx, orderID uuid.UUID, item model.OrderItem) error {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *orderBiz) RemoveOrderItem(ctx contextx.Contextx, orderID uuid.UUID, menuItemID uuid.UUID) error {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *orderBiz) DeleteOrder(ctx contextx.Contextx, id uuid.UUID) error {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *orderBiz) ListOrdersByUser(
	ctx contextx.Contextx,
	userID uuid.UUID,
	options orderB.ListOrdersOptions,
) (orders []model.Order, total int, err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *orderBiz) ListOrdersByRestaurant(
	ctx contextx.Contextx,
	restaurantID uuid.UUID,
	options orderB.ListOrdersOptions,
) (orders []model.Order, total int, err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}
