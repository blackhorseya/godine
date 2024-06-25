package order

import (
	"errors"
	"net/http"
	"time"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/order/model"
	"github.com/blackhorseya/godine/entity/order/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultTimeout = 5 * time.Second
	dbName         = "godine"
	collName       = "orders"
)

type mongodb struct {
	rw *mongo.Client
}

// NewMongodb create and return a new mongodb
func NewMongodb(rw *mongo.Client) repo.IOrderRepo {
	return &mongodb{rw: rw}
}

func (i *mongodb) Create(ctx contextx.Contextx, order *model.Order) error {
	ctx, span := otelx.Span(ctx, "order.mongodb.create")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	if order.ID == "" {
		order.ID = uuid.New().String()
	}

	_, err := i.rw.Database(dbName).Collection(collName).InsertOne(timeout, order)
	if err != nil {
		return err
	}

	return nil
}

func (i *mongodb) GetByID(ctx contextx.Contextx, id string) (item *model.Order, err error) {
	ctx, span := otelx.Span(ctx, "order.mongodb.get_by_id")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	filter := bson.M{"_id": id}
	err = i.rw.Database(dbName).Collection(collName).FindOne(timeout, filter).Decode(&item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errorx.Wrap(http.StatusNotFound, 404, err)
		}

		return nil, err
	}

	return item, nil
}

func (i *mongodb) List(
	ctx contextx.Contextx,
	condition repo.ListCondition,
) (items []*model.Order, total int, err error) {
	ctx, span := otelx.Span(ctx, "order.mongodb.list")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	filter := bson.M{}
	if condition.UserID != "" {
		filter["user_id"] = condition.UserID
	}
	if condition.RestaurantID != "" {
		filter["restaurant_id"] = condition.RestaurantID
	}
	if condition.Status != "" {
		filter["status"] = condition.Status
	}

	opts := options.Find()
	if condition.Limit > 0 {
		opts.SetLimit(int64(condition.Limit))
	}
	if condition.Offset > 0 {
		opts.SetSkip(int64(condition.Offset))
	}

	cursor, err := i.rw.Database(dbName).Collection(collName).Find(timeout, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(timeout)

	err = cursor.All(timeout, &items)
	if err != nil {
		return nil, 0, err
	}

	count, err := i.rw.Database(dbName).Collection(collName).CountDocuments(timeout, filter)
	if err != nil {
		return nil, 0, err
	}

	return items, int(count), nil
}

func (i *mongodb) Update(ctx contextx.Contextx, order *model.Order) error {
	ctx, span := otelx.Span(ctx, "order.mongodb.update")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	filter := bson.M{"_id": order.ID}
	update := bson.M{"$set": order}

	_, err := i.rw.Database(dbName).Collection(collName).UpdateOne(timeout, filter, update)
	if err != nil {
		return err
	}

	return nil
}
