//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	model2 "github.com/blackhorseya/godine/entity/domain/order/model"
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
		items []model2.OrderItem,
		address model2.Address,
		totalAmount float64,
	) (order *model2.Order, err error)

	// GetOrder retrieves the order with the specified ID.
	GetOrder(ctx contextx.Contextx, id string) (order *model2.Order, err error)

	// ListOrders retrieves a list of orders.
	ListOrders(ctx contextx.Contextx, options ListOrdersOptions) (orders []*model2.Order, total int, err error)

	// UpdateOrderStatus updates the status of an existing order.
	UpdateOrderStatus(ctx contextx.Contextx, id string, status string) error

	// ListOrdersByUser retrieves a list of orders placed by a specific user.
	ListOrdersByUser(
		ctx contextx.Contextx,
		userID string,
		options ListOrdersOptions,
	) (orders []*model2.Order, total int, err error)

	// ListOrdersByRestaurant retrieves a list of orders for a specific restaurant.
	ListOrdersByRestaurant(
		ctx contextx.Contextx,
		restaurantID string,
		options ListOrdersOptions,
	) (orders []*model2.Order, total int, err error)

	// OnDeliveryStatusChanged updates the order status when the delivery status changes.
	// This method is called when the delivery status changes, such as from delivering to delivered.
	// It updates the order status based on the delivery status change.
	// Usage: When the delivery status changes and the order status needs to be updated accordingly.
	// Function: Updates the order status based on the delivery status change.
	OnDeliveryStatusChanged(ctx contextx.Contextx, orderID string, status string) error
}
