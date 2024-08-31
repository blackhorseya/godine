package mongodbx

import (
	"github.com/blackhorseya/godine/entity/domain/notification/model"
	"github.com/blackhorseya/godine/entity/domain/notification/repo"
	"github.com/blackhorseya/godine/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodbNotificationRepo struct {
	utils.IRepository[*model.Notification]

	rw   *mongo.Client
	coll *mongo.Collection
}

// NewNotificationRepo is to create a new mongodbNotificationRepo.
func NewNotificationRepo(rw *mongo.Client) repo.INotificationRepo {
	coll := rw.Database("godine").Collection("notifications")

	return &mongodbNotificationRepo{
		IRepository: utils.NewMongoRepository[*model.Notification](coll),
		rw:          rw,
		coll:        coll,
	}
}
