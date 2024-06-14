package user

import (
	"github.com/blackhorseya/godine/entity/user/model"
	"github.com/blackhorseya/godine/entity/user/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodb struct {
	rw *mongo.Client
}

// NewMongodb is used to create an instance of mongodb.
func NewMongodb(rw *mongo.Client) repo.IUserRepo {
	return &mongodb{rw: rw}
}

func (i *mongodb) Create(ctx contextx.Contextx, user *model.User) error {
	// todo: 2024/6/14|sean|implement me
	panic("implement me")
}

func (i *mongodb) GetByID(ctx contextx.Contextx, id string) (item *model.User, err error) {
	// todo: 2024/6/14|sean|implement me
	panic("implement me")
}

func (i *mongodb) List(
	ctx contextx.Contextx,
	condition repo.ListCondition,
) (items []*model.User, total int, err error) {
	// todo: 2024/6/14|sean|implement me
	panic("implement me")
}

func (i *mongodb) Update(ctx contextx.Contextx, user *model.User) error {
	// todo: 2024/6/14|sean|implement me
	panic("implement me")
}

func (i *mongodb) Delete(ctx contextx.Contextx, id string) error {
	// todo: 2024/6/14|sean|implement me
	panic("implement me")
}
