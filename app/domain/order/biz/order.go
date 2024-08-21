package biz

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/blackhorseya/godine/app/infra/otelx"
	logisticsB "github.com/blackhorseya/godine/entity/domain/logistics/biz"
	logisticsM "github.com/blackhorseya/godine/entity/domain/logistics/model"
	notifyB "github.com/blackhorseya/godine/entity/domain/notification/biz"
	notifyM "github.com/blackhorseya/godine/entity/domain/notification/model"
	orderB "github.com/blackhorseya/godine/entity/domain/order/biz"
	orderM "github.com/blackhorseya/godine/entity/domain/order/model"
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

//nolint:funlen // this function is necessary
func (i *orderBiz) CreateOrder(
	ctx contextx.Contextx,
	userID, restaurantID string,
	options []*orderM.OrderItem,
	address *orderM.Address,
	totalAmount float64,
) (order *orderM.Order, err error) {
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

	user, err := i.userService.GetUser(ctx, userID)
	if err != nil {
		ctx.Error(
			"get user from service failed",
			zap.Error(err),
			zap.String("user_id", userID),
		)
		return nil, err
	}

	items := make([]*orderM.OrderItem, 0, len(options))
	for _, option := range options {
		menuItem, err2 := i.menuService.GetMenuItem(ctx, restaurant.GetId(), option.MenuItemId)
		if err2 != nil {
			ctx.Error(
				"get menu item from service failed",
				zap.Error(err2),
				zap.String("menu_item_id", option.MenuItemId),
			)
			return nil, err2
		}
		if !menuItem.IsAvailable {
			ctx.Error(
				"menu item not available",
				zap.String("menu_item_id", option.MenuItemId),
			)
			return nil, errorx.Wrap(http.StatusConflict, 409, errors.New("menu item not available"))
		}

		item := orderM.NewOrderItem(menuItem.GetId(), menuItem.Name, menuItem.Price, int(option.Quantity))
		items = append(items, item)
	}

	order = orderM.NewOrder(user.Id, restaurant.GetId(), items)
	err = i.orders.Create(ctx, order)
	if err != nil {
		ctx.Error(
			"create order failed",
			zap.Error(err),
			zap.Any("order", &order),
		)
		return nil, err
	}

	err = i.notifyService.CreateNotification(ctx, notifyM.NewNotification(
		user.Id,
		user.Id,
		strconv.FormatInt(order.Id, 10),
		"order created",
	))
	if err != nil {
		ctx.Error(
			"create notification failed",
			zap.Error(err),
			zap.Any("order", &order),
		)
		return nil, err
	}

	delivery := logisticsM.NewDelivery(strconv.FormatInt(order.Id, 10), user.Id)
	err = i.logisticsService.CreateDelivery(ctx, delivery)
	if err != nil {
		ctx.Error(
			"create delivery failed",
			zap.Error(err),
			zap.Any("order", &order),
		)
		return nil, err
	}

	order.DeliveryId = delivery.Id
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

func (i *orderBiz) GetOrder(ctx contextx.Contextx, id string) (order *orderM.Order, err error) {
	ctx, span := otelx.Span(ctx, "biz.order.get_order")
	defer span.End()

	return i.orders.GetByID(ctx, id)
}

func (i *orderBiz) ListOrders(
	ctx contextx.Contextx,
	options orderB.ListOrdersOptions,
) (orders []*orderM.Order, total int, err error) {
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
	ctx, span := otelx.Span(ctx, "biz.order.update_order_status")
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

	notify := notifyM.NewNotification(
		order.UserId,
		order.UserId,
		strconv.FormatInt(order.Id, 10),
		"order status to "+event.Name,
	)
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
) (orders []*orderM.Order, total int, err error) {
	ctx, span := otelx.Span(ctx, "biz.order.list_orders_by_user")
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
) (orders []*orderM.Order, total int, err error) {
	ctx, span := otelx.Span(ctx, "biz.order.list_orders_by_restaurant")
	defer span.End()

	return i.orders.List(ctx, repo.ListCondition{
		UserID:       "",
		RestaurantID: restaurantID,
		Status:       options.Status,
		Limit:        options.Size,
		Offset:       (options.Page - 1) * options.Size,
	})
}
