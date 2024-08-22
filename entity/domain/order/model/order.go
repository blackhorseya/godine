package model

import (
	"errors"

	"github.com/blackhorseya/godine/pkg/contextx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewOrder creates a new order.
func NewOrder(userID, restaurantID string, items []*OrderItem) *Order {
	totalAmount := 0.0
	for _, item := range items {
		totalAmount += item.Price * float64(item.Quantity)
	}

	return &Order{
		UserId:       userID,
		RestaurantId: restaurantID,
		Items:        items,
		TotalAmount:  totalAmount,
		Status:       OrderStatus_ORDER_STATUS_PENDING,
		CreatedAt:    timestamppb.Now(),
		UpdatedAt:    timestamppb.Now(),
	}
}

// Next transitions the order to the next state.
func (x *Order) Next(ctx contextx.Contextx) (event *OrderEvent, err error) {
	// TODO: 2024/8/21|sean|implement the order state transition logic
	return nil, errors.New("not implemented")
}

// NewOrderItem creates a new order item.
func NewOrderItem(menuItemID string, price float64, quantity int) *OrderItem {
	return &OrderItem{
		OrderId:    0,
		MenuItemId: menuItemID,
		Quantity:   int64(quantity),
		Price:      price,
	}
}
