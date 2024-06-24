package model

import (
	"time"

	"github.com/blackhorseya/godine/pkg/contextx"
)

// OrderState is the interface for order state.
type OrderState interface {
	// Next is the next state of the order.
	Next(ctx contextx.Contextx, order *Order) (event *OrderEvent, err error)
}

// PendingState is the pending state of the order.
type PendingState struct{}

func (s *PendingState) Next(ctx contextx.Contextx, order *Order) (event *OrderEvent, err error) {
	order.Status = &ConfirmedState{}
	order.UpdatedAt = time.Now()

	return &OrderEvent{
		Name:    "pending",
		Handler: "example",
	}, nil
}

// ConfirmedState is the confirmed state of the order.
type ConfirmedState struct{}

func (s *ConfirmedState) Next(ctx contextx.Contextx, order *Order) (event *OrderEvent, err error) {
	order.Status = &PreparedState{}
	order.UpdatedAt = time.Now()

	return &OrderEvent{
		Name:    "confirmed",
		Handler: "example",
	}, nil
}

// PreparedState is the prepared state of the order.
type PreparedState struct{}

func (s *PreparedState) Next(ctx contextx.Contextx, order *Order) (event *OrderEvent, err error) {
	order.Status = &OutForDeliveryState{}
	order.UpdatedAt = time.Now()

	return &OrderEvent{
		Name:    "prepared",
		Handler: "example",
	}, nil
}

// OutForDeliveryState is the out-for-delivery state of the order.
type OutForDeliveryState struct{}

func (s *OutForDeliveryState) Next(ctx contextx.Contextx, order *Order) (event *OrderEvent, err error) {
	order.Status = &DeliveredState{}
	order.UpdatedAt = time.Now()

	return &OrderEvent{
		Name:    "out_for_delivery",
		Handler: "example",
	}, nil
}

// DeliveredState is the delivered state of the order.
type DeliveredState struct{}

func (s *DeliveredState) Next(ctx contextx.Contextx, order *Order) (event *OrderEvent, err error) {
	// Delivered is a terminal state, no next state.
	order.UpdatedAt = time.Now()

	return &OrderEvent{
		Name:    "delivered",
		Handler: "example",
	}, nil
}

// CancelledState is the cancelled state of the order.
type CancelledState struct{}

func (s *CancelledState) Next(ctx contextx.Contextx, order *Order) (event *OrderEvent, err error) {
	// Cancelled is a terminal state, no next state.
	order.UpdatedAt = time.Now()

	return &OrderEvent{
		Name:    "cancelled",
		Handler: "example",
	}, nil
}
