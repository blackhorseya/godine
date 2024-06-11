package restaurant

import (
	"time"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/restaurant/model"
	"github.com/blackhorseya/godine/entity/restaurant/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

const (
	defaultTimeout = 5 * time.Second
	dbName         = "godine"
	collName       = "restaurants"
)

type mongodb struct {
	rw *mongo.Client
}

// NewMongodb is a function that returns a new mongodb instance that implements the IRestaurantRepo interface
func NewMongodb(rw *mongo.Client) repo.IRestaurantRepo {
	return &mongodb{rw: rw}
}

func (i *mongodb) Create(ctx contextx.Contextx, data *model.Restaurant) (err error) {
	ctx, span := otelx.Span(ctx, "restaurant.mongodb.create")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	_, err = i.rw.Database(dbName).Collection(collName).InsertOne(timeout, data)
	if err != nil {
		ctx.Error("create restaurant to mongodb failed", zap.Error(err), zap.Any("data", &data))
		return err
	}

	return nil
}

func (i *mongodb) Update(ctx contextx.Contextx, data *model.Restaurant) (err error) {
	ctx, span := otelx.Span(ctx, "restaurant.mongodb.update")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	filter := bson.M{"_id": data.ID}
	update := bson.M{"$set": data}
	_, err = i.rw.Database(dbName).Collection(collName).UpdateOne(timeout, filter, update)
	if err != nil {
		ctx.Error("update restaurant to mongodb failed", zap.Error(err), zap.Any("data", &data))
		return err
	}

	return nil
}

func (i *mongodb) Delete(ctx contextx.Contextx, id string) (err error) {
	ctx, span := otelx.Span(ctx, "restaurant.mongodb.delete")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	filter := bson.M{"_id": id}
	_, err = i.rw.Database(dbName).Collection(collName).DeleteOne(timeout, filter)
	if err != nil {
		ctx.Error("delete restaurant from mongodb failed", zap.Error(err), zap.String("id", id))
		return err
	}

	return nil
}

func (i *mongodb) GetByID(ctx contextx.Contextx, id string) (item *model.Restaurant, err error) {
	ctx, span := otelx.Span(ctx, "restaurant.mongodb.get_by_id")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	filter := bson.M{"_id": id}
	err = i.rw.Database(dbName).Collection(collName).FindOne(timeout, filter).Decode(&item)
	if err != nil {
		ctx.Error("get restaurant by id from mongodb failed", zap.Error(err), zap.String("id", id))
		return nil, err
	}

	return item, nil
}

func (i *mongodb) List(
	ctx contextx.Contextx,
	condition repo.ListCondition,
) (items []*model.Restaurant, total int, err error) {
	ctx, span := otelx.Span(ctx, "restaurant.mongodb.list")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	filter := bson.M{}
	opts := options.Find()
	if condition.Limit > 0 {
		opts.SetLimit(int64(condition.Limit))
	}
	if condition.Offset > 0 {
		opts.SetSkip(int64(condition.Offset))
	}

	cursor, err := i.rw.Database(dbName).Collection(collName).Find(timeout, filter, opts)
	if err != nil {
		ctx.Error("list restaurants from mongodb failed", zap.Error(err))
		return nil, 0, err
	}

	err = cursor.All(timeout, &items)
	if err != nil {
		ctx.Error("decode restaurants from mongodb failed", zap.Error(err))
		return nil, 0, err
	}

	count, err := i.rw.Database(dbName).Collection(collName).CountDocuments(timeout, filter)
	if err != nil {
		ctx.Error("count restaurants from mongodb failed", zap.Error(err))
		return nil, 0, err
	}

	return items, int(count), nil
}
