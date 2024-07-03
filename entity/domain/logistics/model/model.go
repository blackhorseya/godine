package model

import (
	"encoding/json"
	"time"

	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

// Delivery represents a delivery entity.
type Delivery struct {
	// ID is the unique identifier of the delivery.
	ID string `json:"id,omitempty" bson:"_id,omitempty"`

	// OrderID is the identifier of the order associated with the delivery.
	OrderID string `json:"order_id,omitempty" bson:"order_id"`

	// UserID is the identifier of the user who placed the order.
	UserID string `json:"user_id,omitempty" bson:"user_id"`

	// DriverID is the identifier of the driver assigned to the delivery.
	DriverID string `json:"driver_id,omitempty" bson:"driver_id"`

	// Status is the current status of the delivery (e.g., pending, in transit, delivered).
	Status DeliveryState `json:"status,omitempty" bson:"status"`

	// PickupTime is the timestamp when the delivery was picked up.
	PickupTime *time.Time `json:"pickup_time,omitempty" bson:"pickup_time"`

	// DeliveryTime is the timestamp when the delivery was completed.
	DeliveryTime *time.Time `json:"delivery_time,omitempty" bson:"delivery_time"`

	// CreatedAt is the timestamp when the delivery was created.
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`

	// UpdatedAt is the timestamp when the delivery was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

// NewDelivery creates a new delivery entity.
func NewDelivery(orderID string, userID string) *Delivery {
	return &Delivery{
		ID:           "",
		OrderID:      orderID,
		UserID:       userID,
		DriverID:     uuid.New().String(),
		Status:       &PendingState{},
		PickupTime:   nil,
		DeliveryTime: nil,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (x *Delivery) UnmarshalJSON(bytes []byte) error {
	type Alias Delivery
	alias := &struct {
		Status string `json:"status"`
		*Alias
	}{
		Alias: (*Alias)(x),
	}

	if err := json.Unmarshal(bytes, alias); err != nil {
		return err
	}

	state, err := UnmarshalDeliveryState(alias.Status)
	if err != nil {
		return err
	}
	x.Status = state

	return nil
}

func (x *Delivery) MarshalJSON() ([]byte, error) {
	type Alias Delivery
	return json.Marshal(&struct {
		*Alias `json:",inline"`
		Status string `json:"status,omitempty"`
	}{
		Alias:  (*Alias)(x),
		Status: x.Status.String(),
	})
}

func (x *Delivery) UnmarshalBSON(bytes []byte) error {
	type Alias Delivery
	alias := &struct {
		Status string `bson:"status"`
		*Alias `bson:",inline"`
	}{
		Alias: (*Alias)(x),
	}

	if err := bson.Unmarshal(bytes, alias); err != nil {
		return err
	}

	state, err := UnmarshalDeliveryState(alias.Status)
	if err != nil {
		return err
	}
	x.Status = state

	return nil
}

func (x *Delivery) MarshalBSON() ([]byte, error) {
	type Alias Delivery
	alias := &struct {
		*Alias `bson:",inline"`
		Status string `bson:"status"`
	}{
		Alias:  (*Alias)(x),
		Status: x.Status.String(),
	}

	return bson.Marshal(alias)
}

// Next will transition the delivery to the next state.
func (x *Delivery) Next(ctx contextx.Contextx) (event *DeliveryEvent, err error) {
	return x.Status.Next(ctx, x)
}
