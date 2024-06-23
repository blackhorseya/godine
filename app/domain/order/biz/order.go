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
	"go.uber.org/zap"
)

type orderBiz struct {
	restaurantService restB.IRestaurantBiz
	menuService       restB.IMenuBiz
	userService       userB.IUserBiz

	orders repo.IOrderRepo
}

// NewOrderBiz create and return a new order orderB
func NewOrderBiz(
	restaurantService restB.IRestaurantBiz,
	menuService restB.IMenuBiz,
	userService userB.IUserBiz,
	orders repo.IOrderRepo,
) orderB.IOrderBiz {
	return &orderBiz{
		restaurantService: restaurantService,
		menuService:       menuService,
		userService:       userService,
		orders:            orders,
	}
}

func (i *orderBiz) CreateOrder(
	ctx contextx.Contextx,
	userID, restaurantID string,
	options []model.OrderItem,
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
			zap.String("restaurant_id", restaurantID),
		)
		return nil, err
	}
	if restaurant == nil {
		ctx.Error(
			"restaurant not found",
			zap.String("restaurant_id", restaurantID),
		)
		return nil, errorx.Wrap(http.StatusNotFound, 404, errors.New("restaurant not found"))
	}

	user, err := i.userService.GetUser(ctx, userID)
	if err != nil {
		ctx.Error(
			"get user from service failed",
			zap.Error(err),
			zap.String("user_id", userID),
		)
		return nil, err
	}
	if user == nil {
		ctx.Error(
			"user not found",
			zap.String("user_id", userID),
		)
		return nil, errorx.Wrap(http.StatusNotFound, 404, errors.New("user not found"))
	}

	order = model.NewOrder(user.ID, restaurant.ID, options, address, totalAmount)
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

func (i *orderBiz) GetOrder(ctx contextx.Contextx, id string) (order *model.Order, err error) {
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

func (i *orderBiz) UpdateOrderStatus(ctx contextx.Contextx, id string, status string) error {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *orderBiz) AddOrderItem(ctx contextx.Contextx, orderID string, item model.OrderItem) error {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *orderBiz) RemoveOrderItem(ctx contextx.Contextx, orderID string, menuItemID string) error {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *orderBiz) DeleteOrder(ctx contextx.Contextx, id string) error {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *orderBiz) ListOrdersByUser(
	ctx contextx.Contextx,
	userID string,
	options orderB.ListOrdersOptions,
) (orders []model.Order, total int, err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *orderBiz) ListOrdersByRestaurant(
	ctx contextx.Contextx,
	restaurantID string,
	options orderB.ListOrdersOptions,
) (orders []model.Order, total int, err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}
