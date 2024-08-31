package mongodbx

import (
	"github.com/blackhorseya/godine/entity/domain/payment/model"
	"github.com/blackhorseya/godine/entity/domain/payment/repo"
	"github.com/blackhorseya/godine/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodbPaymentRepo struct {
	utils.IRepository[*model.Payment]

	rw   *mongo.Client
	coll *mongo.Collection
}

// NewPaymentRepo is to create a new mongodbPaymentRepo.
func NewPaymentRepo(rw *mongo.Client) repo.IPaymentRepo {
	coll := rw.Database("godine").Collection("payments")

	return &mongodbPaymentRepo{
		IRepository: utils.NewMongoRepository[*model.Payment](coll),
		rw:          rw,
		coll:        coll,
	}
}
