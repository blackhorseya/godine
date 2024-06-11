package restaurant

import (
	"time"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/restaurant/model"
	"github.com/blackhorseya/godine/entity/restaurant/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.mongodb.org/mongo-driver/mongo"
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
		return err
	}

	return nil
}

func (i *mongodb) Update(ctx contextx.Contextx, data *model.Restaurant) (err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *mongodb) Delete(ctx contextx.Contextx, id string) (err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *mongodb) GetByID(ctx contextx.Contextx, id string) (item *model.Restaurant, err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *mongodb) List(
	ctx contextx.Contextx,
	condition repo.ListCondition,
) (items []*model.Restaurant, total int, err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}
