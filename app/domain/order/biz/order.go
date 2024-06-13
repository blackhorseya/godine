package biz

import (
	orderB "github.com/blackhorseya/godine/entity/order/biz"
	"github.com/blackhorseya/godine/entity/order/model"
	"github.com/blackhorseya/godine/entity/order/repo"
	restB "github.com/blackhorseya/godine/entity/restaurant/biz"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/google/uuid"
)

type orderBiz struct {
	restaurantService restB.IRestaurantBiz

	orders repo.IOrderRepo
}

// NewOrderBiz create and return a new order orderB
func NewOrderBiz(restaurantService restB.IRestaurantBiz, orders repo.IOrderRepo) orderB.IOrderBiz {
	return &orderBiz{
		restaurantService: restaurantService,
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
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
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
