package model

import (
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
	// todo: 2024/6/23|sean|implement me

	return &OrderEvent{
		Name:    "pending",
		Handler: "example",
	}, nil
}
