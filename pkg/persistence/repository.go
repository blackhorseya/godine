package persistence

import (
	"context"
	"time"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	defaultTimeout  = 5 * time.Second
	defaultLimit    = int64(10)
	defaultMaxLimit = int64(100)
)

// BaseModelInterface ensures that the type has an embedded BaseModel or equivalent fields.
type BaseModelInterface interface {
	GetID() string
	SetID(id primitive.ObjectID)
	GetCreatedAt() *timestamppb.Timestamp
	SetCreatedAt(t *timestamppb.Timestamp)
	GetUpdatedAt() *timestamppb.Timestamp
	SetUpdatedAt(t *timestamppb.Timestamp)
}

// Pagination is a struct for pagination.
type Pagination struct {
	Limit  int64
	Offset int64
}

// IRepository is a generic interface for repositories.
type IRepository[T BaseModelInterface] interface {
	Create(c context.Context, item T) error
	GetByID(c context.Context, id string) (item T, err error)
	List(c context.Context, cond Pagination) (items []T, total int, err error)
	Update(c context.Context, item T) error
	Delete(c context.Context, id string) error
}

type mongoRepository[T BaseModelInterface] struct {
	coll *mongo.Collection
}

// NewMongoRepository is to create a new mongo repository.
func NewMongoRepository[T BaseModelInterface](coll *mongo.Collection) IRepository[T] {
	return &mongoRepository[T]{coll: coll}
}

func (x *mongoRepository[T]) Create(c context.Context, item T) error {
	_, span := otelx.Tracer.Start(c, "Create")
	defer span.End()

	logger := ctxzap.Extract(c)
	logger.Debug("create item", zap.Any("item", &item))

	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	if item.GetID() == "" {
		item.SetID(primitive.NewObjectID())
	}
	item.SetCreatedAt(timestamppb.Now())
	item.SetUpdatedAt(timestamppb.Now())

	_, err := x.coll.InsertOne(timeout, item)
	if err != nil {
		logger.Error("failed to insert item", zap.Error(err))
		return err
	}

	return nil
}

func (x *mongoRepository[T]) GetByID(c context.Context, id string) (item T, err error) {
	_, span := otelx.Tracer.Start(c, "GetByID")
	defer span.End()

	logger := ctxzap.Extract(c)
	logger.Debug("get item by ID", zap.String("id", id))

	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("invalid ObjectID", zap.Error(err), zap.String("id", id))
		return item, err
	}

	var result T
	err = x.coll.FindOne(timeout, bson.M{"_id": objectID}).Decode(&result)
	if err != nil {
		logger.Error("failed to get item by ID", zap.Error(err), zap.String("id", id))
		return item, err
	}

	return result, nil
}

func (x *mongoRepository[T]) List(c context.Context, cond Pagination) (items []T, total int, err error) {
	_, span := otelx.Tracer.Start(c, "List")
	defer span.End()

	logger := ctxzap.Extract(c)
	logger.Debug("list items", zap.Any("condition", cond))

	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	limit, skip := defaultLimit, int64(0)
	if 0 < cond.Limit && cond.Limit <= defaultMaxLimit {
		limit = cond.Limit
	}
	if cond.Offset > 0 {
		skip = cond.Offset
	}

	opts := options.Find().SetLimit(limit).SetSkip(skip).SetSort(bson.M{"_id": -1})
	cursor, err := x.coll.Find(timeout, bson.M{}, opts)
	if err != nil {
		logger.Error("failed to list items", zap.Error(err))
		return nil, 0, err
	}
	defer cursor.Close(timeout)

	err = cursor.All(timeout, &items)
	if err != nil {
		logger.Error("failed to decode items", zap.Error(err))
		return nil, 0, err
	}

	count, err := x.coll.CountDocuments(timeout, bson.M{})
	if err != nil {
		logger.Error("failed to count items", zap.Error(err))
		return nil, 0, err
	}

	return items, int(count), nil
}

func (x *mongoRepository[T]) Update(c context.Context, item T) error {
	_, span := otelx.Tracer.Start(c, "Update")
	defer span.End()

	logger := ctxzap.Extract(c)
	logger.Debug("update item", zap.Any("item", &item))

	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	item.SetUpdatedAt(timestamppb.Now())
	oid, err := primitive.ObjectIDFromHex(item.GetID())
	if err != nil {
		logger.Error("invalid ObjectID", zap.Error(err), zap.String("id", item.GetID()))
		return err
	}

	filter := bson.M{"_id": oid}
	update := bson.M{"$set": item}

	_, err = x.coll.UpdateOne(timeout, filter, update)
	if err != nil {
		logger.Error("failed to update item", zap.Error(err))
		return err
	}

	return nil
}

func (x *mongoRepository[T]) Delete(c context.Context, id string) error {
	_, span := otelx.Tracer.Start(c, "Delete")
	defer span.End()

	logger := ctxzap.Extract(c)
	logger.Debug("delete item", zap.String("id", id))

	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("invalid ObjectID", zap.Error(err), zap.String("id", id))
		return err
	}

	_, err = x.coll.DeleteOne(timeout, bson.M{"_id": oid})
	if err != nil {
		logger.Error("failed to delete item", zap.Error(err), zap.String("id", id))
		return err
	}

	return nil
}
