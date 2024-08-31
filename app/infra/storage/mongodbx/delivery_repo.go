package mongodbx

import (
	"github.com/blackhorseya/godine/entity/domain/logistics/model"
	"github.com/blackhorseya/godine/entity/domain/logistics/repo"
	"github.com/blackhorseya/godine/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodbDeliveryRepo struct {
	utils.IRepository[*model.Delivery]

	rw   *mongo.Client
	coll *mongo.Collection
}

// NewDeliveryRepo will create a new delivery repository.
func NewDeliveryRepo(rw *mongo.Client) repo.IDeliveryRepo {
	coll := rw.Database("godine").Collection("deliveries")

	return &mongodbDeliveryRepo{
		IRepository: utils.NewMongoRepository[*model.Delivery](coll),
		rw:          rw,
		coll:        coll,
	}
}
