package user

import (
	"time"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/user/model"
	"github.com/blackhorseya/godine/entity/user/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

const (
	defaultTimeout = 5 * time.Second
	dbName         = "godine"
	collName       = "users"
)

type mongodb struct {
	rw *mongo.Client
}

// NewMongodb is used to create an instance of mongodb.
func NewMongodb(rw *mongo.Client) repo.IUserRepo {
	return &mongodb{rw: rw}
}

func (i *mongodb) Create(ctx contextx.Contextx, user *model.User) error {
	ctx, span := otelx.Span(ctx, "user.mongodb.create")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	if user.ID == "" {
		user.ID = uuid.New().String()
	}

	_, err := i.rw.Database(dbName).Collection(collName).InsertOne(timeout, user)
	if err != nil {
		ctx.Error("insert user to mongodb failed", zap.Error(err))
		return err
	}

	return nil
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
