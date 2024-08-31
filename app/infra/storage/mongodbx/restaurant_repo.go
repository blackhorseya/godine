package mongodbx

import (
	"context"

	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/entity/domain/restaurant/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type mongodbRestaurantRepo struct {
	rw   *mongo.Client
	coll *mongo.Collection
}

// NewMongoDBRestaurantRepo is to create a new mongodbRestaurantRepo.
func NewMongoDBRestaurantRepo(rw *mongo.Client) repo.IRestaurantRepo {
	return &mongodbRestaurantRepo{
		rw:   rw,
		coll: rw.Database("godine").Collection("restaurants"),
	}
}

func (i *mongodbRestaurantRepo) Create(c context.Context, data *model.Restaurant) (err error) {
	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	ctx := contextx.Background()

	if data.GetId() == "" {
		data.Id = primitive.NewObjectID().Hex()
	}
	data.CreatedAt = timestamppb.Now()
	data.UpdatedAt = timestamppb.Now()

	_, err = i.coll.InsertOne(timeout, data)
	if err != nil {
		ctx.Error("create restaurant to mongodb failed", zap.Error(err), zap.Any("data", &data))
		return err
	}

	return nil
}

func (i *mongodbRestaurantRepo) Update(c context.Context, data *model.Restaurant) (err error) {
	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	ctx := contextx.Background()

	id, err := primitive.ObjectIDFromHex(data.GetId())
	if err != nil {
		ctx.Error("parse restaurant id failed", zap.Error(err), zap.String("id", data.GetId()))
		return err
	}
	data.UpdatedAt = timestamppb.Now()

	filter := bson.M{"_id": id}
	update := bson.M{"$set": data}
	_, err = i.coll.UpdateOne(timeout, filter, update)
	if err != nil {
		ctx.Error("update restaurant to mongodb failed", zap.Error(err), zap.Any("data", &data))
		return err
	}

	return nil
}

func (i *mongodbRestaurantRepo) Delete(c context.Context, id string) (err error) {
	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	ctx := contextx.Background()

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.Error("parse restaurant id failed", zap.Error(err), zap.String("id", id))
		return err
	}

	filter := bson.M{"_id": oid}
	_, err = i.coll.DeleteOne(timeout, filter)
	if err != nil {
		ctx.Error("delete restaurant to mongodb failed", zap.Error(err), zap.String("id", id))
		return err
	}

	return nil
}

func (i *mongodbRestaurantRepo) GetByID(c context.Context, id string) (item *model.Restaurant, err error) {
	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	ctx := contextx.Background()

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.Error("parse restaurant id failed", zap.Error(err), zap.String("id", id))
		return nil, err
	}

	filter := bson.M{"_id": oid}
	err = i.coll.FindOne(timeout, filter).Decode(&item)
	if err != nil {
		ctx.Error("get restaurant by id from mongodb failed", zap.Error(err), zap.String("id", id))
		return nil, err
	}

	return item, nil
}

func (i *mongodbRestaurantRepo) List(
	c context.Context,
	cond repo.ListCondition,
) (items []*model.Restaurant, total int, err error) {
	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	ctx := contextx.Background()

	filter := bson.M{}

	limit, skip := defaultLimit, int64(0)
	if 0 < cond.Limit && cond.Limit <= defaultMaxLimit {
		limit = cond.Limit
	}
	if cond.Offset > 0 {
		skip = cond.Offset
	}
	opts := options.Find().SetLimit(limit).SetSkip(skip).SetSort(bson.M{"_id": -1})

	cursor, err := i.coll.Find(timeout, filter, opts)
	if err != nil {
		ctx.Error("list restaurants from mongodb failed", zap.Error(err))
		return nil, 0, err
	}
	defer cursor.Close(timeout)

	err = cursor.All(timeout, &items)
	if err != nil {
		ctx.Error("decode restaurants from mongodb failed", zap.Error(err))
		return nil, 0, err
	}

	count, err := i.coll.CountDocuments(timeout, filter)
	if err != nil {
		ctx.Error("count restaurants from mongodb failed", zap.Error(err))
		return nil, 0, err
	}

	return items, int(count), nil
}