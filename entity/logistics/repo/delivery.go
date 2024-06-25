//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/godine/entity/logistics/model"
	"github.com/blackhorseya/godine/pkg/contextx"
)

// ListCondition represents a condition for listing deliveries.
type ListCondition struct {
	Limit  int
	Offset int

	DriverID string
}

// IDeliveryRepo represents a repository interface for managing delivery entities.
type IDeliveryRepo interface {
	Create(ctx contextx.Contextx, item *model.Delivery) error
	GetByID(ctx contextx.Contextx, id string) (item *model.Delivery, err error)
	List(ctx contextx.Contextx, condition ListCondition) (items []*model.Delivery, total int, err error)
	Update(ctx contextx.Contextx, item *model.Delivery) error
	Delete(ctx contextx.Contextx, id string) error
}
