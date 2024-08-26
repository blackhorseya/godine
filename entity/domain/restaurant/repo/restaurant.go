//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"context"

	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
)

// ListCondition is a struct that defines the condition for listing restaurants
type ListCondition struct {
	Limit  int64
	Offset int64
}

// IRestaurantRepo is an interface that defines the methods that the restaurant repository should implement
type IRestaurantRepo interface {
	Create(c context.Context, data *model.Restaurant) (err error)
	Update(c context.Context, data *model.Restaurant) (err error)
	Delete(c context.Context, id string) (err error)
	GetByID(c context.Context, id string) (item *model.Restaurant, err error)
	List(c context.Context, condition ListCondition) (items []*model.Restaurant, total int, err error)
}
