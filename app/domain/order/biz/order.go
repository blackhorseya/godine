package biz

import (
	"errors"
	"net/http"

	"github.com/blackhorseya/godine/app/infra/otelx"
	logisticsB "github.com/blackhorseya/godine/entity/domain/logistics/biz"
	model2 "github.com/blackhorseya/godine/entity/domain/logistics/model"
	notifyB "github.com/blackhorseya/godine/entity/domain/notification/biz"
	model3 "github.com/blackhorseya/godine/entity/domain/notification/model"
	orderB "github.com/blackhorseya/godine/entity/domain/order/biz"
	model4 "github.com/blackhorseya/godine/entity/domain/order/model"
	"github.com/blackhorseya/godine/entity/domain/order/repo"
	rB "github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	userB "github.com/blackhorseya/godine/entity/domain/user/biz"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"go.uber.org/zap"
)

type orderBiz struct {
	restaurantService rB.IRestaurantBiz
	menuService       rB.IMenuBiz
	userService       userB.IUserBiz
	logisticsService  logisticsB.ILogisticsBiz
	notifyService     notifyB.INotificationBiz

	orders repo.IOrderRepo
}

// NewOrderBiz create and return a new order orderB
func NewOrderBiz(
	restaurantService rB.IRestaurantBiz,
	menuService rB.IMenuBiz,
	userService userB.IUserBiz,
	logisticsService logisticsB.ILogisticsBiz,
	notifyService notifyB.INotificationBiz,
	orders repo.IOrderRepo,
) orderB.IOrderBiz {
	return &orderBiz{
		restaurantService: restaurantService,
		menuService:       menuService,
		userService:       userService,
		logisticsService:  logisticsService,
		notifyService:     notifyService,
		orders:            orders,
	}
}

//nolint:funlen // it's okay
func (i *orderBiz) CreateOrder(
	ctx contextx.Contextx,
	userID, restaurantID string,
	options []model4.OrderItem,
	address model4.Address,
	totalAmount float64,
) (order *model4.Order, err error) {
	ctx, span := otelx.Span(ctx, "rB.order.create_order")
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

	items := make([]model4.OrderItem, 0, len(options))
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

		item := model4.NewOrderItem(menuItem.ID, menuItem.Name, menuItem.Price, option.Quantity)
		items = append(items, *item)
	}

	order = model4.NewOrder(user.ID, restaurant.ID, items)
	err = i.orders.Create(ctx, order)
	if err != nil {
		ctx.Error(
			"create order failed",
			zap.Error(err),
			zap.Any("order", &order),
		)
		return nil, err
	}

	err = i.notifyService.CreateNotification(ctx, model3.NewNotify(user.ID, user.ID, order.ID, "order created"))
	if err != nil {
		ctx.Error(
			"create notification failed",
			zap.Error(err),
			zap.Any("order", &order),
		)
		return nil, err
	}

	delivery := model2.NewDelivery(order.ID, user.ID)
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

func (i *orderBiz) GetOrder(ctx contextx.Contextx, id string) (order *model4.Order, err error) {
	ctx, span := otelx.Span(ctx, "rB.order.get_order")
	defer span.End()

	return i.orders.GetByID(ctx, id)
}

func (i *orderBiz) ListOrders(
	ctx contextx.Contextx,
	options orderB.ListOrdersOptions,
) (orders []*model4.Order, total int, err error) {
	ctx, span := otelx.Span(ctx, "rB.order.list_orders")
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
	ctx, span := otelx.Span(ctx, "rB.order.update_order_status")
	defer span.End()

	order, err := i.orders.GetByID(ctx, id)
	if err != nil {
		ctx.Error(
			"get order failed",
			zap.Error(err),
			zap.String("order_id", id),
		)
		return err
	}
	if order == nil {
		ctx.Error(
			"order not found",
			zap.String("order_id", id),
		)
		return errorx.Wrap(http.StatusNotFound, 404, errors.New("order not found"))
	}

	event, err := order.Next(ctx)
	if err != nil {
		ctx.Error(
			"next order failed",
			zap.Error(err),
			zap.String("order_id", id),
		)
		return err
	}

	ctx.Debug("order executed event", zap.Any("event", &event))

	notify := model3.NewNotify(order.UserID, order.UserID, order.ID, "order status to "+event.Name)
	err = i.notifyService.CreateNotification(ctx, notify)
	if err != nil {
		ctx.Error(
			"create notification failed",
			zap.Error(err),
			zap.Any("order", &order),
		)
		return err
	}

	err = i.orders.Update(ctx, order)
	if err != nil {
		ctx.Error(
			"update order failed",
			zap.Error(err),
			zap.Any("order", &order),
		)
		return err
	}

	return nil
}

func (i *orderBiz) ListOrdersByUser(
	ctx contextx.Contextx,
	userID string,
	options orderB.ListOrdersOptions,
) (orders []*model4.Order, total int, err error) {
	ctx, span := otelx.Span(ctx, "rB.order.list_orders_by_user")
	defer span.End()

	return i.orders.List(ctx, repo.ListCondition{
		UserID:       userID,
		RestaurantID: "",
		Status:       options.Status,
		Limit:        options.Size,
		Offset:       (options.Page - 1) * options.Size,
	})
}

func (i *orderBiz) ListOrdersByRestaurant(
	ctx contextx.Contextx,
	restaurantID string,
	options orderB.ListOrdersOptions,
) (orders []*model4.Order, total int, err error) {
	ctx, span := otelx.Span(ctx, "rB.order.list_orders_by_restaurant")
	defer span.End()

	return i.orders.List(ctx, repo.ListCondition{
		UserID:       "",
		RestaurantID: restaurantID,
		Status:       options.Status,
		Limit:        options.Size,
		Offset:       (options.Page - 1) * options.Size,
	})
}
