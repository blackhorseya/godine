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
	"go.mongodb.org/mongo-driver/mongo"
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
	// todo: 2024/7/23|sean|implement me
	panic("implement me")
}

func (i *mongodb) Create(ctx contextx.Contextx, item *model.Payment) (err error) {
	// todo: 2024/7/23|sean|implement me
	panic("implement me")
}

func (i *mongodb) Update(ctx contextx.Contextx, item *model.Payment) (err error) {
	// todo: 2024/7/23|sean|implement me
	panic("implement me")
}
