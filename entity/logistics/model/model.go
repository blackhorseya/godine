package model

import (
	"time"
)

// Delivery represents a delivery entity.
type Delivery struct {
	// ID is the unique identifier of the delivery.
	ID string `json:"id,omitempty" bson:"_id,omitempty"`

	// OrderID is the identifier of the order associated with the delivery.
	OrderID string `json:"order_id,omitempty" bson:"order_id"`

	// DriverID is the identifier of the driver assigned to the delivery.
	DriverID string `json:"driver_id,omitempty" bson:"driver_id"`

	// Status is the current status of the delivery (e.g., pending, in transit, delivered).
	Status string `json:"status,omitempty" bson:"status"`

	// PickupTime is the timestamp when the delivery was picked up.
	PickupTime *time.Time `json:"pickup_time,omitempty" bson:"pickup_time"`

	// DeliveryTime is the timestamp when the delivery was completed.
	DeliveryTime *time.Time `json:"delivery_time,omitempty" bson:"delivery_time"`

	// CreatedAt is the timestamp when the delivery was created.
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`

	// UpdatedAt is the timestamp when the delivery was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

// DeliveryStatus represents the status of a delivery.
type DeliveryStatus struct {
	// ID is the unique identifier of the delivery status.
	ID string `json:"id,omitempty" bson:"_id,omitempty"`

	// DeliveryID is the identifier of the delivery associated with the status.
	DeliveryID string `json:"delivery_id,omitempty" bson:"delivery_id"`

	// Status is the status of the delivery (e.g., pending, in transit, delivered).
	Status string `json:"status,omitempty" bson:"status"`

	// UpdatedAt is the timestamp when the status was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}
