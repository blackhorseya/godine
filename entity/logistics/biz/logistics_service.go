//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/godine/entity/logistics/model"
	"github.com/blackhorseya/godine/pkg/contextx"
)

// ListDeliveriesOptions defines the options for listing deliveries.
type ListDeliveriesOptions struct {
	Page int `form:"page" default:"1" minimum:"1"`
	Size int `form:"size" default:"10" minimum:"1" maximum:"100"`
}

// ILogisticsBiz defines the interface for logistics management operations.
type ILogisticsBiz interface {
	// CreateDelivery creates a new delivery.
	CreateDelivery(ctx contextx.Contextx, delivery model.Delivery) error

	// UpdateDeliveryStatus updates the status of an existing delivery.
	UpdateDeliveryStatus(ctx contextx.Contextx, deliveryID string, status string) error

	// GetDelivery retrieves a delivery by its ID.
	GetDelivery(ctx contextx.Contextx, deliveryID string) (model.Delivery, error)

	// ListDeliveriesByDriver retrieves a list of deliveries assigned to a specific driver.
	ListDeliveriesByDriver(ctx contextx.Contextx, driverID string, options ListDeliveriesOptions) ([]model.Delivery, int, error)
}
