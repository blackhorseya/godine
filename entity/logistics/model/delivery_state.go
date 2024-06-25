package model

import (
	"fmt"

	"github.com/blackhorseya/godine/pkg/contextx"
)

// DeliveryEvent represents an event during the delivery state transition.
type DeliveryEvent struct {
	Name    string `json:"name,omitempty"`
	Handler string `json:"handler,omitempty"`
}

// DeliveryState interface defines the behavior of different delivery states.
type DeliveryState interface {
	fmt.Stringer

	Next(ctx contextx.Contextx, delivery *Delivery) (event *DeliveryEvent, err error)
}
