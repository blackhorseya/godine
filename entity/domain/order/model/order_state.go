package model

import (
	"fmt"

	"github.com/blackhorseya/godine/pkg/contextx"
)

// OrderState is the interface for order state.
type OrderState interface {
	fmt.Stringer

	// Next is the next state of the order.
	Next(ctx contextx.Contextx, order *Order) (event *OrderEvent, err error)
}

func UnmarshalOrderState(raw string) (OrderState, error) {
	switch raw {
	case "pending":
		return &PendingState{}, nil
	case "confirmed":
		return &ConfirmedState{}, nil
	case "prepared":
		return &PreparedState{}, nil
	case "out_for_delivery":
		return &OutForDeliveryState{}, nil
	case "delivered":
		return &DeliveredState{}, nil
	case "cancelled":
		return &CancelledState{}, nil
	default:
		return nil, fmt.Errorf("unknown order state: %s", raw)
	}
}

var _ OrderState = &PendingState{}

// PendingState is the pending state of the order.
type PendingState struct{}

func (s *PendingState) String() string {
	return "pending"
}

func (s *PendingState) Next(ctx contextx.Contextx, order *Order) (event *OrderEvent, err error) {
	// order.Status = &ConfirmedState{}
	// order.UpdatedAt = time.Now()

	return &OrderEvent{
		Name:      order.Status.String(),
		HandlerId: "example",
	}, nil
}

var _ OrderState = &ConfirmedState{}

// ConfirmedState is the confirmed state of the order.
type ConfirmedState struct{}

func (s *ConfirmedState) String() string {
	return "confirmed"
}

func (s *ConfirmedState) Next(ctx contextx.Contextx, order *Order) (event *OrderEvent, err error) {
	// order.Status = &PreparedState{}
	// order.UpdatedAt = time.Now()

	return &OrderEvent{
		Name:      order.Status.String(),
		HandlerId: "example",
	}, nil
}

var _ OrderState = &PreparedState{}

// PreparedState is the prepared state of the order.
type PreparedState struct{}

func (s *PreparedState) String() string {
	return "prepared"
}

func (s *PreparedState) Next(ctx contextx.Contextx, order *Order) (event *OrderEvent, err error) {
	// order.Status = &OutForDeliveryState{}
	// order.UpdatedAt = time.Now()

	return &OrderEvent{
		Name:      order.Status.String(),
		HandlerId: "example",
	}, nil
}

var _ OrderState = &OutForDeliveryState{}

// OutForDeliveryState is the out-for-delivery state of the order.
type OutForDeliveryState struct{}

func (s *OutForDeliveryState) String() string {
	return "out_for_delivery"
}

func (s *OutForDeliveryState) Next(ctx contextx.Contextx, order *Order) (event *OrderEvent, err error) {
	// order.Status = &DeliveredState{}
	// order.UpdatedAt = time.Now()

	return &OrderEvent{
		Name:      order.Status.String(),
		HandlerId: "example",
	}, nil
}

var _ OrderState = &DeliveredState{}

// DeliveredState is the delivered state of the order.
type DeliveredState struct{}

func (s *DeliveredState) String() string {
	return "delivered"
}

func (s *DeliveredState) Next(ctx contextx.Contextx, order *Order) (event *OrderEvent, err error) {
	// Delivered is a terminal state, no next state.
	// order.UpdatedAt = time.Now()

	return &OrderEvent{
		Name:      order.Status.String(),
		HandlerId: "example",
	}, nil
}

var _ OrderState = &CancelledState{}

// CancelledState is the cancelled state of the order.
type CancelledState struct{}

func (s *CancelledState) String() string {
	return "cancelled"
}

func (s *CancelledState) Next(ctx contextx.Contextx, order *Order) (event *OrderEvent, err error) {
	// Cancelled is a terminal state, no next state.
	// order.UpdatedAt = time.Now()

	return &OrderEvent{
		Name:      order.Status.String(),
		HandlerId: "example",
	}, nil
}
