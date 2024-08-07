package user

import (
	"errors"
	"net/http"
	"time"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/entity/domain/user/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		user.ID = primitive.NewObjectID().Hex()
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

	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.Error("convert user id to object id failed", zap.Error(err))
		return nil, errorx.Wrap(http.StatusBadRequest, 400, err)
	}

	filter := bson.M{"_id": hex}
	err = i.rw.Database(dbName).Collection(collName).FindOne(timeout, filter).Decode(&item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			ctx.Error("user not found", zap.String("user_id", id))
			return nil, errorx.Wrap(http.StatusNotFound, 404, err)
		}

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
	ctx, span := otelx.Span(ctx, "user.mongodb.update")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}

	_, err := i.rw.Database(dbName).Collection(collName).UpdateOne(timeout, filter, update)
	if err != nil {
		ctx.Error("update user to mongodb failed", zap.Error(err))
		return err
	}

	return nil
}

func (i *mongodb) Delete(ctx contextx.Contextx, id string) error {
	ctx, span := otelx.Span(ctx, "user.mongodb.delete")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	filter := bson.M{"_id": id}

	_, err := i.rw.Database(dbName).Collection(collName).DeleteOne(timeout, filter)
	if err != nil {
		ctx.Error("delete user from mongodb failed", zap.Error(err))
		return err
	}

	return nil
}
