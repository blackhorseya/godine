package payment

import (
	"errors"
	"net/http"
	"time"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/payment/model"
	"github.com/blackhorseya/godine/entity/domain/payment/repo"
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
	collName       = "payments"
)

type mongodb struct {
	rw *mongo.Client
}

// NewMongodb is a function that returns a new mongodb instance that implements the IPaymentRepo interface
func NewMongodb(rw *mongo.Client) repo.IPaymentRepo {
	return &mongodb{rw: rw}
}

func (i *mongodb) GetByID(ctx contextx.Contextx, id string) (item *model.Payment, err error) {
	ctx, span := otelx.Span(ctx, "biz.payment.mongodb.get_by_id")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	filter := bson.M{"_id": id}
	err = i.rw.Database(dbName).Collection(collName).FindOne(timeout, filter).Decode(&item)
	if errors.Is(err, mongo.ErrNoDocuments) {
		ctx.Error("payment not found", zap.Error(err), zap.String("id", id))
		return nil, errorx.Wrap(http.StatusNotFound, 404, err)
	} else if err != nil {
		ctx.Error("failed to find payment", zap.Error(err), zap.String("id", id))
		return nil, err
	}

	return item, nil
}

func (i *mongodb) List(ctx contextx.Contextx, cond repo.ListCondition) (items []*model.Payment, total int, err error) {
	ctx, span := otelx.Span(ctx, "biz.payment.mongodb.list")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	if cond.Limit == 0 {
		cond.Limit = 100
	}
	if cond.Offset < 0 {
		cond.Offset = 0
	}

	filter := bson.M{}
	opts := options.Find()
	opts.SetLimit(int64(cond.Limit))
	opts.SetSkip(int64(cond.Offset))
	opts.SetSort(bson.M{"updated_at": -1})

	cursor, err := i.rw.Database(dbName).Collection(collName).Find(timeout, filter)
	if err != nil {
		ctx.Error("failed to find payments", zap.Error(err), zap.Any("condition", &cond))
		return nil, 0, err
	}

	err = cursor.All(timeout, &items)
	if err != nil {
		ctx.Error("failed to decode payments", zap.Error(err))
		return nil, 0, err
	}

	count, err := i.rw.Database(dbName).Collection(collName).CountDocuments(timeout, filter)
	if err != nil {
		ctx.Error("failed to count payments", zap.Error(err))
		return nil, 0, err
	}

	return items, int(count), nil
}

func (i *mongodb) Create(ctx contextx.Contextx, item *model.Payment) (err error) {
	ctx, span := otelx.Span(ctx, "biz.payment.mongodb.create")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	if item.ID == "" {
		item.ID = primitive.NewObjectID().Hex()
	}

	_, err = i.rw.Database(dbName).Collection(collName).InsertOne(timeout, item)
	if err != nil {
		ctx.Error("failed to create payment", zap.Error(err), zap.Any("payment", &item))
		return err
	}

	return nil
}

func (i *mongodb) Update(ctx contextx.Contextx, item *model.Payment) (err error) {
	ctx, span := otelx.Span(ctx, "biz.payment.mongodb.update")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	filter := bson.M{"_id": item.ID}
	update := bson.M{"$set": item}

	_, err = i.rw.Database(dbName).Collection(collName).UpdateOne(timeout, filter, update)
	if err != nil {
		ctx.Error("failed to update payment", zap.Error(err), zap.Any("payment", &item))
		return err
	}

	return nil
}
