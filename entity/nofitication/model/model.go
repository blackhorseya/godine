package model

import (
	"time"
)

// Notification represents a notification entity.
type Notification struct {
	// ID is the unique identifier of the notification.
	ID string `json:"id,omitempty" bson:"_id,omitempty"`

	// UserID is the identifier of the user to whom the notification is sent.
	UserID string `json:"user_id,omitempty" bson:"user_id"`

	// OrderID is the identifier of the order associated with the notification.
	OrderID string `json:"order_id,omitempty" bson:"order_id"`

	// Type represents the type of notification (e.g., order_status, delivery_status).
	Type string `json:"type,omitempty" bson:"type"`

	// Message is the content of the notification.
	Message string `json:"message,omitempty" bson:"message"`

	// Status is the current status of the notification (e.g., pending, sent).
	Status string `json:"status,omitempty" bson:"status"`

	// CreatedAt is the timestamp when the notification was created.
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`

	// UpdatedAt is the timestamp when the notification was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}
