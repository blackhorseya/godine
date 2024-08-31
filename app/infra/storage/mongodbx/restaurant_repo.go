package mongodbx

import (
	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/entity/domain/restaurant/repo"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodbRestaurantRepo struct {
	IRepository[model.Restaurant]

	rw   *mongo.Client
	coll *mongo.Collection
}

// NewMongoDBRestaurantRepo is to create a new mongodbRestaurantRepo.
func NewMongoDBRestaurantRepo(rw *mongo.Client) repo.IRestaurantRepo {
	coll := rw.Database("godine").Collection("restaurants")

	return &mongodbRestaurantRepo{
		IRepository: NewMongoRepository[model.Restaurant](coll),
		rw:          rw,
		coll:        coll,
	}
}
