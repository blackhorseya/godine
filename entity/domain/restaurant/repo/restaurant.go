//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/godine/app/infra/storage/mongodbx"
	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
)

// ListCondition is a struct that defines the condition for listing restaurants
type ListCondition struct {
	Limit  int64
	Offset int64
}

// IRestaurantRepo is an interface that defines the methods that the restaurant repository should implement
type IRestaurantRepo interface {
	mongodbx.IRepository[model.Restaurant]
}
