package biz

import (
	"github.com/blackhorseya/godine/entity/domain/order/biz"
	"github.com/blackhorseya/godine/entity/events"
	"github.com/blackhorseya/godine/pkg/contextx"
)

type deliveryStatusUpdatedHandler struct {
}

// NewDeliveryStatusUpdatedHandler creates a new delivery status updated event handler.
func NewDeliveryStatusUpdatedHandler() biz.DeliveryStatusChangedHandler {
	return &deliveryStatusUpdatedHandler{}
}

func (i *deliveryStatusUpdatedHandler) Handle(ctx contextx.Contextx, event *events.DomainEvent) error {
	// todo: 2024/6/26|sean|handle delivery status updated event
	return nil
}
