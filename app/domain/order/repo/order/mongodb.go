package order

import (
	"time"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/order/model"
	"github.com/blackhorseya/godine/entity/order/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	err = i.rw.Database(dbName).Collection(collName).FindOne(timeout, filter).Decode(item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (i *mongodb) ListByUserID(
	ctx contextx.Contextx,
	userID string,
	condition repo.ListCondition,
) (items []*model.Order, total int, err error) {
	// todo: 2024/6/13|sean|implement me
	panic("implement me")
}

func (i *mongodb) ListByRestaurantID(
	ctx contextx.Contextx,
	restaurantID string,
	condition repo.ListCondition,
) (items []*model.Order, total int, err error) {
	// todo: 2024/6/13|sean|implement me
	panic("implement me")
}

func (i *mongodb) UpdateStatus(ctx contextx.Contextx, order *model.Order, status string) error {
	// todo: 2024/6/13|sean|implement me
	panic("implement me")
}
