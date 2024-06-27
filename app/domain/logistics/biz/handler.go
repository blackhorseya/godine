package biz

import (
	"github.com/blackhorseya/godine/entity/domain/logistics/biz"
	"github.com/blackhorseya/godine/entity/domain/logistics/model"
	"github.com/blackhorseya/godine/pkg/contextx"
)

type deliveryStatusChanged struct {
}

// NewDeliveryStatusChangedHandler creates a new delivery status changed handler.
func NewDeliveryStatusChangedHandler() biz.DeliveryStatusChangedHandler {
	return &deliveryStatusChanged{}
}

func (i *deliveryStatusChanged) On(ctx contextx.Contextx) (ch chan *model.DeliveryEvent, err error) {
	// todo: 2024/6/27|sean|implement the logic
	panic("implement me")
}
