//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"context"

	"github.com/blackhorseya/godine/entity/domain/order/model"
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
	Create(c context.Context, order *model.Order) error
	GetByID(c context.Context, id string) (item *model.Order, err error)
	List(c context.Context, condition ListCondition) (items []*model.Order, total int, err error)
	Update(c context.Context, order *model.Order) error
}
