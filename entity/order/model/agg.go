package model

import (
	"time"

	"github.com/google/uuid"
)

// OrderAggregate represents the order aggregate root.
type OrderAggregate struct {
	// Order is the core entity of the aggregate.
	Order Order `json:"order" bson:"order"`
}

// NewOrder creates a new OrderAggregate.
func NewOrder(userID, restaurantID uuid.UUID, items []OrderItem, address Address, totalAmount float64) *OrderAggregate {
	return &OrderAggregate{
		Order: Order{
			ID:           uuid.New(),
			UserID:       userID,
			RestaurantID: restaurantID,
			Items:        items,
			Status:       "pending",
			TotalAmount:  totalAmount,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}
}

// UpdateStatus updates the status of the order.
func (oa *OrderAggregate) UpdateStatus(newStatus string) {
	oa.Order.Status = newStatus
	oa.Order.UpdatedAt = time.Now()
}

// AddItem adds an item to the order.
func (oa *OrderAggregate) AddItem(item OrderItem) {
	oa.Order.Items = append(oa.Order.Items, item)
	oa.Order.UpdatedAt = time.Now()
}
