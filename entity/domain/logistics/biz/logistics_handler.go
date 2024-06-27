//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/godine/entity/events"
)

// DeliveryStatusChangedHandler represents the interface for handling delivery status changed events.
type DeliveryStatusChangedHandler interface {
	events.EventHandler
}
