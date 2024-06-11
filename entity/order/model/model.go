package model

import (
	"time"

	"github.com/google/uuid"
)

// Order represents an order entity.
type Order struct {
	// ID is the unique identifier of the order.
	ID uuid.UUID `json:"id,omitempty" bson:"_id,omitempty"`

	// UserID is the identifier of the user who placed the order.
	UserID uuid.UUID `json:"user_id,omitempty" bson:"userId"`

	// RestaurantID is the identifier of the restaurant where the order was placed.
	RestaurantID uuid.UUID `json:"restaurant_id,omitempty" bson:"restaurantId"`

	// Items are the list of items in the order.
	Items []OrderItem `json:"items,omitempty" bson:"items"`

	// Status is the current status of the order (e.g., pending, confirmed, delivered).
	Status string `json:"status,omitempty" bson:"status"`

	// TotalAmount is the total amount of the order.
	TotalAmount float64 `json:"total_amount,omitempty" bson:"totalAmount"`

	// CreatedAt is the timestamp when the order was created.
	CreatedAt time.Time `json:"created_at,omitempty" bson:"createdAt"`

	// UpdatedAt is the timestamp when the order was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updatedAt"`
}

// OrderItem represents an item in the order.
type OrderItem struct {
	// MenuItemID is the identifier of the menu item.
	MenuItemID uuid.UUID `json:"menu_item_id,omitempty" bson:"menuItemId"`

	// Quantity is the quantity of the menu item ordered.
	Quantity int `json:"quantity,omitempty" bson:"quantity"`

	// Price is the price of a single unit of the menu item.
	Price float64 `json:"price,omitempty" bson:"price"`
}
