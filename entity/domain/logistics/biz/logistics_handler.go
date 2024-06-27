package biz

import (
	"github.com/blackhorseya/godine/entity/domain/logistics/model"
	"github.com/blackhorseya/godine/pkg/contextx"
)

// DeliveryStatusChangedHandler represents the interface for handling delivery status changed events.
type DeliveryStatusChangedHandler interface {
	On(ctx contextx.Contextx) (ch chan *model.DeliveryEvent, err error)
}
