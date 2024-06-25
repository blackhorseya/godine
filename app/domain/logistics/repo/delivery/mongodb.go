package delivery

import (
	"github.com/blackhorseya/godine/entity/logistics/model"
	"github.com/blackhorseya/godine/entity/logistics/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodb struct {
	rw *mongo.Client
}

// NewMongodb is to create a mongodb instance
func NewMongodb(rw *mongo.Client) repo.IDeliveryRepo {
	return &mongodb{rw: rw}
}

func (i *mongodb) Create(ctx contextx.Contextx, item *model.Delivery) error {
	// todo: 2024/6/25|sean|implement me
	panic("implement me")
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
