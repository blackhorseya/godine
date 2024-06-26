//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/godine/entity/domain/order/model"
	"github.com/blackhorseya/godine/pkg/contextx"
)

// ListCondition is a struct that defines the conditions for listing orders
type ListCondition struct {
	UserID       string
	RestaurantID string
	Status       string

	Limit  int
	Offset int
}

// IOrderRepo is an interface that defines the methods that the order repository should implement
type IOrderRepo interface {
	Create(ctx contextx.Contextx, order *model.Order) error

	GetByID(ctx contextx.Contextx, id string) (item *model.Order, err error)

	List(ctx contextx.Contextx, condition ListCondition) (items []*model.Order, total int, err error)

	Update(ctx contextx.Contextx, order *model.Order) error
}
