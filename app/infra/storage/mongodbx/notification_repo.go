package mongodbx

import (
	"context"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/notification/model"
	"github.com/blackhorseya/godine/entity/domain/notification/repo"
	"github.com/blackhorseya/godine/pkg/persistence"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type mongodbNotificationRepo struct {
	persistence.IRepository[*model.Notification]

	rw   *mongo.Client
	coll *mongo.Collection
}

// NewNotificationRepo is to create a new mongodbNotificationRepo.
func NewNotificationRepo(rw *mongo.Client) repo.INotificationRepo {
	coll := rw.Database("godine").Collection("notifications")

	return &mongodbNotificationRepo{
		IRepository: persistence.NewMongoRepository[*model.Notification](coll),
		rw:          rw,
		coll:        coll,
	}
}

func (i *mongodbNotificationRepo) ListByReceiverID(
	c context.Context,
	receiverID string,
	cond persistence.Pagination,
) (items []*model.Notification, total int64, err error) {
	_, span := otelx.Tracer.Start(c, "mongodbNotificationRepo.ListByReceiverID")
	defer span.End()

	logger := ctxzap.Extract(c)

	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	filter := bson.M{"user_id": receiverID}

	total, err = i.coll.CountDocuments(timeout, filter)
	if err != nil {
		logger.Error("count documents failed", zap.Error(err))
		return nil, 0, err
	}

	limit, skip := defaultLimit, int64(0)
	if 0 < cond.Limit && cond.Limit <= defaultMaxLimit {
		limit = cond.Limit
	}
	if cond.Offset > 0 {
		skip = cond.Offset
	}
	opts := options.Find().SetLimit(limit).SetSkip(skip).SetSort(bson.M{"created_at": -1})

	cur, err := i.coll.Find(timeout, filter, opts)
	if err != nil {
		logger.Error("find documents failed", zap.Error(err))
		return nil, 0, err
	}
	defer cur.Close(timeout)

	err = cur.All(timeout, &items)
	if err != nil {
		logger.Error("decode documents failed", zap.Error(err))
		return nil, 0, err
	}

	return items, total, nil
}
