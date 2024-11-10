package mongodbx

import (
	"context"

	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/entity/domain/restaurant/repo"
	"github.com/blackhorseya/godine/pkg/persistence"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodbRestaurantRepo struct {
	persistence.IRepository[*model.Restaurant]

	rw   *mongo.Client
	coll *mongo.Collection
}

// NewRestaurantRepo is to create a new mongodbRestaurantRepo.
func NewRestaurantRepo(rw *mongo.Client) repo.IRestaurantRepo {
	coll := rw.Database("godine").Collection("restaurants")

	return &mongodbRestaurantRepo{
		IRepository: persistence.NewMongoRepository[*model.Restaurant](coll),
		rw:          rw,
		coll:        coll,
	}
}

func (x *mongodbRestaurantRepo) CreateReservation(
	c context.Context,
	restaurant *model.Restaurant,
	reservation *model.Order,
) error {
	// TODO: 2024/11/10|sean|implement CreateReservation
	return nil
}
