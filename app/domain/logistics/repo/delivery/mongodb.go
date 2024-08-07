package delivery

import (
	"errors"
	"net/http"
	"time"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/logistics/model"
	"github.com/blackhorseya/godine/entity/domain/logistics/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

const (
	defaultTimeout = 5 * time.Second
	dbName         = "godine"
	collName       = "deliveries"
)

type mongodb struct {
	rw *mongo.Client
}

// NewMongodb is to create a mongodb instance
func NewMongodb(rw *mongo.Client) repo.IDeliveryRepo {
	return &mongodb{rw: rw}
}

func (i *mongodb) Create(ctx contextx.Contextx, item *model.Delivery) error {
	ctx, span := otelx.Span(ctx, "biz.logistics.repo.delivery.mongodb.Create")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	if item.ID == "" {
		item.ID = primitive.NewObjectID().Hex()
	}
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	_, err := i.rw.Database(dbName).Collection(collName).InsertOne(timeout, item)
	if err != nil {
		ctx.Error("insert one delivery to mongodb failed", zap.Error(err), zap.Any("delivery", &item))
		return err
	}

	return nil
}

func (i *mongodb) GetByID(ctx contextx.Contextx, id string) (item *model.Delivery, err error) {
	ctx, span := otelx.Span(ctx, "biz.logistics.repo.delivery.mongodb.GetByID")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.Error("convert id to object id failed", zap.Error(err), zap.String("id", id))
		return nil, err
	}

	filter := bson.M{"_id": hex}
	err = i.rw.Database(dbName).Collection(collName).FindOne(timeout, filter).Decode(&item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errorx.Wrap(http.StatusNotFound, 404, err)
		}

		ctx.Error("find one delivery from mongodb failed", zap.Error(err), zap.String("id", id))
		return nil, err
	}

	return item, nil
}

func (i *mongodb) List(
	ctx contextx.Contextx,
	condition repo.ListCondition,
) (items []*model.Delivery, total int, err error) {
	ctx, span := otelx.Span(ctx, "biz.logistics.repo.delivery.mongodb.List")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	filter := bson.M{}
	if condition.DriverID != "" {
		filter["driver_id"] = condition.DriverID
	}

	opts := options.Find()
	if condition.Limit > 0 {
		opts.SetLimit(int64(condition.Limit))
	}
	if condition.Offset > 0 {
		opts.SetSkip(int64(condition.Offset))
	}
	opts.SetSort(bson.M{"updated_at": -1})

	cursor, err := i.rw.Database(dbName).Collection(collName).Find(timeout, filter, opts)
	if err != nil {
		ctx.Error("find deliveries from mongodb failed", zap.Error(err))
		return nil, 0, err
	}
	defer cursor.Close(timeout)

	err = cursor.All(timeout, &items)
	if err != nil {
		ctx.Error("decode deliveries from mongodb failed", zap.Error(err))
		return nil, 0, err
	}

	count, err := i.rw.Database(dbName).Collection(collName).CountDocuments(timeout, filter)
	if err != nil {
		ctx.Error("count deliveries from mongodb failed", zap.Error(err))
		return nil, 0, err
	}

	return items, int(count), nil
}

func (i *mongodb) Update(ctx contextx.Contextx, item *model.Delivery) error {
	ctx, span := otelx.Span(ctx, "biz.logistics.repo.delivery.mongodb.Update")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	item.UpdatedAt = time.Now()

	filter := bson.M{"_id": item.ID}
	update := bson.M{"$set": item}
	_, err := i.rw.Database(dbName).Collection(collName).UpdateOne(timeout, filter, update)
	if err != nil {
		ctx.Error("update one delivery to mongodb failed", zap.Error(err), zap.Any("delivery", &item))
		return err
	}

	return nil
}

func (i *mongodb) Delete(ctx contextx.Contextx, id string) error {
	ctx, span := otelx.Span(ctx, "biz.logistics.repo.delivery.mongodb.Delete")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	filter := bson.M{"_id": id}
	_, err := i.rw.Database(dbName).Collection(collName).DeleteOne(timeout, filter)
	if err != nil {
		ctx.Error("delete one delivery from mongodb failed", zap.Error(err), zap.String("id", id))
		return err
	}

	return nil
}
