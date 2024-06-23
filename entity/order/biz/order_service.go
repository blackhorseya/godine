//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/godine/entity/order/model"
	"github.com/blackhorseya/godine/pkg/contextx"
)

// ListOrdersOptions defines the options for listing orders.
type ListOrdersOptions struct {
	// Page is the page number.
	Page int

	// PageSize is the number of items per page.
	PageSize int
}

// IOrderBiz defines the business operations for order management.
type IOrderBiz interface {
	// CreateOrder creates a new order.
	CreateOrder(
		ctx contextx.Contextx,
		userID, restaurantID string,
		items []model.OrderItem,
		address model.Address,
		totalAmount float64,
	) (order *model.Order, err error)

	// GetOrder retrieves the order with the specified ID.
	GetOrder(ctx contextx.Contextx, id string) (order *model.Order, err error)

	// ListOrders retrieves a list of orders.
	ListOrders(ctx contextx.Contextx, options ListOrdersOptions) (orders []model.Order, total int, err error)

	// UpdateOrderStatus updates the status of an existing order.
	UpdateOrderStatus(ctx contextx.Contextx, id string, status string) error

	// AddOrderItem adds an item to an existing order.
	AddOrderItem(ctx contextx.Contextx, orderID string, item model.OrderItem) error

	// RemoveOrderItem removes an item from an existing order.
	RemoveOrderItem(ctx contextx.Contextx, orderID string, menuItemID string) error

	// DeleteOrder deletes an order by its ID.
	DeleteOrder(ctx contextx.Contextx, id string) error

	// ListOrdersByUser retrieves a list of orders placed by a specific user.
	ListOrdersByUser(
		ctx contextx.Contextx,
		userID string,
		options ListOrdersOptions,
	) (orders []model.Order, total int, err error)

	// ListOrdersByRestaurant retrieves a list of orders for a specific restaurant.
	ListOrdersByRestaurant(
		ctx contextx.Contextx,
		restaurantID string,
		options ListOrdersOptions,
	) (orders []model.Order, total int, err error)
}
