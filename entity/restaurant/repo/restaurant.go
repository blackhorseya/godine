//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/godine/entity/restaurant/model"
	"github.com/blackhorseya/godine/pkg/contextx"
)

// ListCondition is a struct that defines the condition for listing restaurants
type ListCondition struct {
	Limit  int
	Offset int
}

// IRestaurantRepo is an interface that defines the methods that the restaurant repository should implement
type IRestaurantRepo interface {
	Create(ctx contextx.Contextx, data *model.Restaurant) (err error)
	Update(ctx contextx.Contextx, data *model.Restaurant) (err error)
	Delete(ctx contextx.Contextx, id string) (err error)
	GetByID(ctx contextx.Contextx, id string) (item *model.Restaurant, err error)
	List(ctx contextx.Contextx, condition ListCondition) (items []*model.Restaurant, total int, err error)
}
