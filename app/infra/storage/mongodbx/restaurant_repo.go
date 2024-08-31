package mongodbx

import (
	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/entity/domain/restaurant/repo"
	"github.com/blackhorseya/godine/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodbRestaurantRepo struct {
	utils.IRepository[*model.Restaurant]

	rw   *mongo.Client
	coll *mongo.Collection
}

// NewRestaurantRepo is to create a new mongodbRestaurantRepo.
func NewRestaurantRepo(rw *mongo.Client) repo.IRestaurantRepo {
	coll := rw.Database("godine").Collection("restaurants")

	return &mongodbRestaurantRepo{
		IRepository: utils.NewMongoRepository[*model.Restaurant](coll),
		rw:          rw,
		coll:        coll,
	}
}
