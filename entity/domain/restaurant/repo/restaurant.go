//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/pkg/persistence"
)

// IRestaurantRepo is an interface that defines the methods that the restaurant repository should implement
type IRestaurantRepo interface {
	persistence.IRepository[*model.Restaurant]
}
