package user

import (
	"time"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/user/model"
	"github.com/blackhorseya/godine/entity/user/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	ctx, span := otelx.Span(ctx, "user.mongodb.getByID")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	filter := bson.M{"_id": id}
	err = i.rw.Database(dbName).Collection(collName).FindOne(timeout, filter).Decode(&item)
	if err != nil {
		ctx.Error("get user by id from mongodb failed", zap.Error(err))
		return nil, err
	}

	return item, nil
}

func (i *mongodb) List(
	ctx contextx.Contextx,
	condition repo.ListCondition,
) (items []*model.User, total int, err error) {
	ctx, span := otelx.Span(ctx, "user.mongodb.list")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	filter := bson.M{}
	opts := options.Find()
	if condition.Limit > 0 {
		opts.SetLimit(int64(condition.Limit))
	}
	if condition.Offset > 0 {
		opts.SetSkip(int64(condition.Offset))
	}

	cursor, err := i.rw.Database(dbName).Collection(collName).Find(timeout, filter, opts)
	if err != nil {
		ctx.Error("list user from mongodb failed", zap.Error(err))
		return nil, 0, err
	}

	err = cursor.All(timeout, &items)
	if err != nil {
		ctx.Error("decode user from mongodb failed", zap.Error(err))
		return nil, 0, err
	}

	count, err := i.rw.Database(dbName).Collection(collName).CountDocuments(timeout, filter)
	if err != nil {
		ctx.Error("count user from mongodb failed", zap.Error(err))
		return nil, 0, err
	}

	return items, int(count), nil
}

func (i *mongodb) Update(ctx contextx.Contextx, user *model.User) error {
	// todo: 2024/6/14|sean|implement me
	panic("implement me")
}

func (i *mongodb) Delete(ctx contextx.Contextx, id string) error {
	// todo: 2024/6/14|sean|implement me
	panic("implement me")
}
