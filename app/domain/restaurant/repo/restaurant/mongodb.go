package restaurant

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/entity/domain/restaurant/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	defaultTimeout = 5 * time.Second
	defaultLimit   = 100
	dbName         = "godine"
	collName       = "restaurants"
)

type mongodb struct {
	rw  *mongo.Client
	rdb *redis.Client
}

// NewMongodb is a function that returns a new mongodb instance that implements the IRestaurantRepo interface
func NewMongodb(rw *mongo.Client, rdb *redis.Client) repo.IRestaurantRepo {
	return &mongodb{
		rw:  rw,
		rdb: rdb,
	}
}

func (i *mongodb) Create(ctx contextx.Contextx, data *model.Restaurant) (err error) {
	ctx, span := otelx.Span(ctx, "restaurant.mongodb.create")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	if data.GetId() == "" {
		data.Id = primitive.NewObjectID().Hex()
	}
	data.CreatedAt = timestamppb.Now()
	data.UpdatedAt = timestamppb.Now()

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

	id, err := primitive.ObjectIDFromHex(data.GetId())
	if err != nil {
		ctx.Error("parse restaurant id failed", zap.Error(err), zap.String("id", data.GetId()))
		return err
	}
	data.UpdatedAt = timestamppb.Now()

	filter := bson.M{"_id": id}
	update := bson.M{"$set": data}
	_, err = i.rw.Database(dbName).Collection(collName).UpdateOne(timeout, filter, update)
	if err != nil {
		ctx.Error("update restaurant to mongodb failed", zap.Error(err), zap.Any("data", &data))
		return err
	}

	err = cacheRestaurant(ctx, i.rdb, data.GetId(), data)
	if err != nil {
		ctx.Error("cache restaurant to redis failed", zap.Error(err), zap.String("id", data.GetId()))
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

	// get from redis
	val, err := i.rdb.Get(ctx, id).Result()
	switch {
	case errors.Is(err, redis.Nil):
		var hex primitive.ObjectID
		hex, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			ctx.Error("parse restaurant id failed", zap.Error(err), zap.String("id", id))
			return nil, err
		}

		filter := bson.M{"_id": hex}
		err = i.rw.Database(dbName).Collection(collName).FindOne(timeout, filter).Decode(&item)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				ctx.Error("restaurant not found", zap.String("id", id))
				return nil, errorx.Wrap(http.StatusNotFound, 404, err)
			}

			ctx.Error("get restaurant by id from mongodb failed", zap.Error(err), zap.String("id", id))
			return nil, err
		}
		err = cacheRestaurant(ctx, i.rdb, id, item)
		if err != nil {
			ctx.Error("cache restaurant to redis failed", zap.Error(err), zap.String("id", id))
		}
	case err != nil:
		ctx.Error("get restaurant by id from redis failed", zap.Error(err), zap.String("id", id))
		return nil, err
	default:
		err = json.Unmarshal([]byte(val), &item)
		if err != nil {
			ctx.Error("decode restaurant from redis failed", zap.Error(err), zap.String("id", id))
			return nil, err
		}
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
	if condition.Limit <= 0 {
		condition.Limit = defaultLimit
	}
	if condition.Offset < 0 {
		condition.Offset = 0
	}
	opts := options.Find().SetLimit(condition.Limit).SetSkip(condition.Offset).SetSort(bson.M{"updated_at": -1})

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

func cacheRestaurant(ctx contextx.Contextx, rdb *redis.Client, id string, restaurant *model.Restaurant) error {
	data, err := json.Marshal(restaurant)
	if err != nil {
		return err
	}

	return rdb.Set(ctx, id, data, 10*time.Minute).Err()
}
