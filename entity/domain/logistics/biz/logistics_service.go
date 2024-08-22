//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/godine/entity/domain/logistics/model"
	"github.com/blackhorseya/godine/pkg/contextx"
)

// ListDeliveriesOptions defines the options for listing deliveries.
type ListDeliveriesOptions struct {
	// Page specifies the page number for pagination.
	Page int `form:"page" default:"1" minimum:"1"`

	// Size specifies the number of items per page.
	Size int `form:"size" default:"10" minimum:"1" maximum:"100"`
}

// ILogisticsBiz defines the interface for logistics management operations. This
// interface abstracts the logistics management functionality, making it easier
// to manage deliveries and their statuses. The design allows for flexibility,
// scalability, and maintainability of the logistics operations in the GoDine
// system.
// Deprecated: use LogisticsServiceServer instead.
type ILogisticsBiz interface {
	// CreateDelivery creates a new delivery. This method is responsible for
	// initializing a new delivery entity with provided details such as OrderID,
	// DriverID, and timestamps. It ensures that all necessary information is
	// captured and stored correctly when a delivery is created. Usage: When a user
	// places an order and a delivery needs to be scheduled. Function: Creates a new
	// delivery record, assigns a driver, and initializes the delivery status.
	CreateDelivery(ctx contextx.Contextx, delivery *model.Delivery) error

	// UpdateDeliveryStatus updates the status of an existing delivery. This method
	// changes the current status of the delivery (e.g., from pending to in transit)
	// and updates the timestamp to reflect the change. It ensures that the delivery
	// status is accurately tracked and updated in real-time. Usage: When the
	// delivery status changes, such as from pending to in transit. Function:
	// Updates the current status of the delivery record and logs the timestamp.
	UpdateDeliveryStatus(ctx contextx.Contextx, deliveryID string, status string) error

	// GetDelivery retrieves a delivery by its ID. This method fetches the details
	// of a specific delivery based on the unique identifier provided. It returns
	// the Delivery entity containing all relevant information. Usage: When you need
	// to query the details of a specific delivery record. Function: Retrieves the
	// delivery record based on its unique identifier.
	GetDelivery(ctx contextx.Contextx, deliveryID string) (item *model.Delivery, err error)

	// ListDeliveriesByDriver retrieves a list of deliveries assigned to a specific
	// driver. This method returns a paginated list of deliveries for a given
	// driver, allowing for easier management and tracking of multiple deliveries.
	// Usage: When you need to manage and track deliveries assigned to a specific
	// driver. Function: Provides a list of deliveries assigned to the specified
	// driver with pagination options.
	ListDeliveriesByDriver(
		ctx contextx.Contextx,
		driverID string,
		options ListDeliveriesOptions,
	) (items []*model.Delivery, total int, err error)
}
