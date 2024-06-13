package model

import (
	"time"

	"github.com/google/uuid"
)

// Order represents an order entity.
type Order struct {
	// ID is the unique identifier of the order.
	ID string `json:"id,omitempty" bson:"_id,omitempty"`

	// UserID is the identifier of the user who placed the order.
	UserID string `json:"user_id,omitempty" bson:"user_id"`

	// RestaurantID is the identifier of the restaurant where the order was placed.
	RestaurantID string `json:"restaurant_id,omitempty" bson:"restaurant_id"`

	// Items are the list of items in the order.
	Items []OrderItem `json:"items,omitempty" bson:"items"`

	// Status is the current status of the order (e.g., pending, confirmed, delivered).
	Status string `json:"status,omitempty" bson:"status"`

	// TotalAmount is the total amount of the order.
	TotalAmount float64 `json:"total_amount,omitempty" bson:"total_amount"`

	// CreatedAt is the timestamp when the order was created.
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`

	// UpdatedAt is the timestamp when the order was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

// NewOrder creates a new order.
func NewOrder(userID, restaurantID string, items []OrderItem, address Address, totalAmount float64) *Order {
	return &Order{
		ID:           uuid.New().String(),
		UserID:       userID,
		RestaurantID: restaurantID,
		Items:        items,
		Status:       "pending",
		TotalAmount:  totalAmount,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

// UpdateStatus updates the status of the order.
func (x *Order) UpdateStatus(newStatus string) {
	x.Status = newStatus
	x.UpdatedAt = time.Now()
}

// AddItem adds an item to the order.
func (x *Order) AddItem(item OrderItem) {
	x.Items = append(x.Items, item)
	x.UpdatedAt = time.Now()
}

// OrderItem represents an item in the order.
type OrderItem struct {
	// MenuItemID is the identifier of the menu item.
	MenuItemID string `json:"menu_item_id,omitempty" bson:"menu_item_id"`

	// Quantity is the quantity of the menu item ordered.
	Quantity int `json:"quantity,omitempty" bson:"quantity"`

	// Price is the price of a single unit of the menu item.
	Price float64 `json:"price,omitempty" bson:"price"`
}
