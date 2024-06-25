package delivery

import (
	"time"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/logistics/model"
	"github.com/blackhorseya/godine/entity/logistics/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
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
		item.ID = uuid.New().String()
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
	// todo: 2024/6/25|sean|implement me
	panic("implement me")
}

func (i *mongodb) List(
	ctx contextx.Contextx,
	condition repo.ListCondition,
) (items []*model.Delivery, total int, err error) {
	// todo: 2024/6/25|sean|implement me
	panic("implement me")
}

func (i *mongodb) Update(ctx contextx.Contextx, item *model.Delivery) error {
	// todo: 2024/6/25|sean|implement me
	panic("implement me")
}

func (i *mongodb) Delete(ctx contextx.Contextx, id string) error {
	// todo: 2024/6/25|sean|implement me
	panic("implement me")
}
