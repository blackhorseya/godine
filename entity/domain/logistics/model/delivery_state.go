package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/blackhorseya/godine/entity/events"
	"github.com/blackhorseya/godine/pkg/contextx"
)

var _ events.DomainEvent = (*DeliveryEvent)(nil)

// DeliveryEvent represents an event during the delivery state transition.
type DeliveryEvent struct {
	Name    string    `json:"name,omitempty"`
	Handler string    `json:"handler,omitempty"`
	Data    *Delivery `json:"data,omitempty"`
}

func (x *DeliveryEvent) Key() []byte {
	return []byte(x.Data.Id)
}

func (x *DeliveryEvent) Value() ([]byte, error) {
	return json.Marshal(x.Data)
}

func (x *DeliveryEvent) OccurredOn(ctx contextx.Contextx) time.Time {
	return time.Now()
}

func (x *DeliveryEvent) Topic() string {
	return "delivery_event"
}

// DeliveryState interface defines the behavior of different delivery states.
type DeliveryState interface {
	fmt.Stringer

	Next(ctx contextx.Contextx, delivery *Delivery) (event *DeliveryEvent, err error)
}

// UnmarshalDeliveryState unmarshals the raw string into a delivery state.
func UnmarshalDeliveryState(raw string) (DeliveryState, error) {
	switch raw {
	case "pending":
		return &PendingState{}, nil
	case "picked_up":
		return &PickedUpState{}, nil
	case "in_transit":
		return &InTransitState{}, nil
	case "completed":
		return &CompletedState{}, nil
	case "cancelled":
		return &CancelledState{}, nil
	default:
		return nil, fmt.Errorf("unknown delivery state: %s", raw)
	}
}

// PendingState represents the pending state of a delivery.
type PendingState struct{}

func (s *PendingState) String() string {
	return "pending"
}

func (s *PendingState) Next(ctx contextx.Contextx, delivery *Delivery) (event *DeliveryEvent, err error) {
	// delivery.Status = &PickedUpState{}
	// delivery.UpdatedAt = time.Now()

	return &DeliveryEvent{
		Name:    delivery.Status.String(),
		Handler: "delivery_handler",
		Data:    delivery,
	}, nil
}

// PickedUpState represents the picked-up state of a delivery.
type PickedUpState struct{}

func (s *PickedUpState) String() string {
	return "picked_up"
}

func (s *PickedUpState) Next(ctx contextx.Contextx, delivery *Delivery) (event *DeliveryEvent, err error) {
	// delivery.Status = &InTransitState{}
	// delivery.UpdatedAt = time.Now()

	return &DeliveryEvent{
		Name:    delivery.Status.String(),
		Handler: "delivery_handler",
		Data:    delivery,
	}, nil
}

// InTransitState represents the in-transit state of a delivery.
type InTransitState struct{}

func (s *InTransitState) String() string {
	return "in_transit"
}

func (s *InTransitState) Next(ctx contextx.Contextx, delivery *Delivery) (event *DeliveryEvent, err error) {
	// delivery.Status = &CompletedState{}
	// delivery.UpdatedAt = time.Now()

	return &DeliveryEvent{
		Name:    delivery.Status.String(),
		Handler: "delivery_handler",
		Data:    delivery,
	}, nil
}

// CompletedState represents the completed state of a delivery.
type CompletedState struct{}

func (s *CompletedState) String() string {
	return "completed"
}

func (s *CompletedState) Next(ctx contextx.Contextx, delivery *Delivery) (event *DeliveryEvent, err error) {
	// Completed is a terminal state, no next state.
	// delivery.UpdatedAt = time.Now()

	return &DeliveryEvent{
		Name:    delivery.Status.String(),
		Handler: "delivery_handler",
		Data:    delivery,
	}, nil
}

// CancelledState represents the cancelled state of a delivery.
type CancelledState struct{}

func (s *CancelledState) String() string {
	return "cancelled"
}

func (s *CancelledState) Next(ctx contextx.Contextx, delivery *Delivery) (event *DeliveryEvent, err error) {
	// Cancelled is a terminal state, no next state.
	// delivery.UpdatedAt = time.Now()

	return &DeliveryEvent{
		Name:    delivery.Status.String(),
		Handler: "delivery_handler",
		Data:    delivery,
	}, nil
}
