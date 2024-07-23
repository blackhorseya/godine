package payment

import (
	"time"

	"github.com/blackhorseya/godine/entity/domain/payment/model"
	"github.com/blackhorseya/godine/entity/domain/payment/repo"
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

// NewMongodb is a function that returns a new mongodb instance that implements the IPaymentRepo interface
func NewMongodb(rw *mongo.Client) repo.IPaymentRepo {
	return &mongodb{rw: rw}
}

func (i *mongodb) GetByID(ctx contextx.Contextx, id string) (item *model.Payment, err error) {
	// todo: 2024/7/23|sean|implement me
	panic("implement me")
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
