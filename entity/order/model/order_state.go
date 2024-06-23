package model

import (
	"github.com/blackhorseya/godine/pkg/contextx"
)

// OrderState is the interface for order state.
type OrderState interface {
	// Update is used to update the order.
	Update(ctx contextx.Contextx, order *Order) (event *OrderEvent, err error)

	// Next is the next state of the order.
	Next(ctx contextx.Contextx, order *Order) (event *OrderEvent, err error)
}
