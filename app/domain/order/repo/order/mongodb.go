package order

import (
	"github.com/blackhorseya/godine/entity/order/model"
	"github.com/blackhorseya/godine/entity/order/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodb struct {
	rw *mongo.Client
}

// NewMongodb create and return a new mongodb
func NewMongodb(rw *mongo.Client) repo.IOrderRepo {
	return &mongodb{rw: rw}
}

func (i *mongodb) Create(ctx contextx.Contextx, order *model.Order) error {
	// todo: 2024/6/13|sean|implement me
	panic("implement me")
}

func (i *mongodb) GetByID(ctx contextx.Contextx, id string) (item *model.Order, err error) {
	// todo: 2024/6/13|sean|implement me
	panic("implement me")
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
