//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/godine/entity/order/model"
	"github.com/blackhorseya/godine/pkg/contextx"
)

// ListOrdersOptions defines the options for listing orders.
type ListOrdersOptions struct {
	// Page is the page number.
	Page int `form:"page" default:"1" minimum:"1"`

	// Size is the number of items per page.
	Size int `form:"size" default:"10" minimum:"1" maximum:"100"`

	// UserID is the ID of the user who placed the order.
	UserID string `form:"user_id"`

	// RestaurantID is the ID of the restaurant that received the order.
	RestaurantID string `form:"restaurant_id"`

	// Status is the status of the order.
	Status string `form:"status" enums:"pending,confirmed,preparing,delivering,delivered,canceled"`
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
	ListOrders(ctx contextx.Contextx, options ListOrdersOptions) (orders []*model.Order, total int, err error)

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
