package biz

import (
	"errors"
	"net/http"

	"github.com/blackhorseya/godine/app/infra/otelx"
	logisticsB "github.com/blackhorseya/godine/entity/logistics/biz"
	model2 "github.com/blackhorseya/godine/entity/logistics/model"
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
	logisticsService  logisticsB.ILogisticsBiz

	orders repo.IOrderRepo
}

// NewOrderBiz create and return a new order orderB
func NewOrderBiz(
	restaurantService restB.IRestaurantBiz,
	menuService restB.IMenuBiz,
	userService userB.IUserBiz,
	logisticsService logisticsB.ILogisticsBiz,
	orders repo.IOrderRepo,
) orderB.IOrderBiz {
	return &orderBiz{
		restaurantService: restaurantService,
		menuService:       menuService,
		userService:       userService,
		logisticsService:  logisticsService,
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

	items := make([]model.OrderItem, 0, len(options))
	for _, option := range options {
		menuItem, err2 := i.menuService.GetMenuItem(ctx, restaurant.ID, option.MenuItemID)
		if err2 != nil {
			ctx.Error(
				"get menu item from service failed",
				zap.Error(err2),
				zap.String("menu_item_id", option.MenuItemID),
			)
			return nil, err2
		}
		if menuItem == nil {
			ctx.Error(
				"menu item not found",
				zap.String("menu_item_id", option.MenuItemID),
			)
			return nil, errorx.Wrap(http.StatusNotFound, 404, errors.New("menu item not found"))
		}
		if !menuItem.IsAvailable {
			ctx.Error(
				"menu item not available",
				zap.String("menu_item_id", option.MenuItemID),
			)
			return nil, errorx.Wrap(http.StatusConflict, 409, errors.New("menu item not available"))
		}

		item := model.NewOrderItem(menuItem.ID, menuItem.Name, menuItem.Price, option.Quantity)
		items = append(items, *item)
	}

	order = model.NewOrder(user.ID, restaurant.ID, items)
	err = i.orders.Create(ctx, order)
	if err != nil {
		ctx.Error(
			"create order failed",
			zap.Error(err),
			zap.Any("order", &order),
		)
		return nil, err
	}

	delivery := model2.NewDelivery(order.ID)
	err = i.logisticsService.CreateDelivery(ctx, delivery)
	if err != nil {
		ctx.Error(
			"create delivery failed",
			zap.Error(err),
			zap.Any("order", &order),
		)
		return nil, err
	}

	order.DeliveryID = delivery.ID
	err = i.orders.Update(ctx, order)
	if err != nil {
		ctx.Error(
			"update order failed",
			zap.Error(err),
			zap.Any("order", &order),
		)
		return nil, err
	}

	return order, nil
}

func (i *orderBiz) GetOrder(ctx contextx.Contextx, id string) (order *model.Order, err error) {
	ctx, span := otelx.Span(ctx, "biz.order.get_order")
	defer span.End()

	return i.orders.GetByID(ctx, id)
}

func (i *orderBiz) ListOrders(
	ctx contextx.Contextx,
	options orderB.ListOrdersOptions,
) (orders []*model.Order, total int, err error) {
	ctx, span := otelx.Span(ctx, "biz.order.list_orders")
	defer span.End()

	return i.orders.List(ctx, repo.ListCondition{
		UserID:       options.UserID,
		RestaurantID: options.RestaurantID,
		Status:       options.Status,
		Limit:        options.Size,
		Offset:       (options.Page - 1) * options.Size,
	})
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
