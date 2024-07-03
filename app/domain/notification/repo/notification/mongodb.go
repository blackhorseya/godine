package notification

import (
	"errors"
	"net/http"
	"time"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/notification/model"
	"github.com/blackhorseya/godine/entity/domain/notification/repo"
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
	collName       = "notifications"
)

type mongodb struct {
	rw *mongo.Client
}

// NewMongodb creates a new MongoDB notification repository.
func NewMongodb(rw *mongo.Client) repo.INotificationRepo {
	return &mongodb{rw: rw}
}

func (i *mongodb) Create(ctx contextx.Contextx, notify *model.Notification) error {
	ctx, span := otelx.Span(ctx, "biz.notification.repo.notification.mongodb.Create")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	if notify.ID == "" {
		notify.ID = primitive.NewObjectID().Hex()
	}
	notify.CreatedAt = time.Now()
	notify.UpdatedAt = time.Now()

	_, err := i.rw.Database(dbName).Collection(collName).InsertOne(timeout, notify)
	if err != nil {
		ctx.Error(
			"insert one notification to mongodb failed",
			zap.Error(err),
			zap.Any("notification", &notify),
		)
		return err
	}

	return nil
}

func (i *mongodb) GetByID(ctx contextx.Contextx, id string) (item *model.Notification, err error) {
	ctx, span := otelx.Span(ctx, "biz.notification.repo.notification.mongodb.GetByID")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.Error("convert id to object id failed", zap.Error(err), zap.String("id", id))
		return nil, errorx.Wrap(http.StatusBadRequest, 400, err)
	}

	filter := bson.M{"_id": hex}
	err = i.rw.Database(dbName).Collection(collName).FindOne(timeout, filter).Decode(&item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			ctx.Error("notification not found", zap.String("id", id))
			return nil, errorx.Wrap(http.StatusNotFound, 404, err)
		}

		ctx.Error("find one notification from mongodb failed", zap.Error(err), zap.String("id", id))
		return nil, err
	}

	return item, nil
}

func (i *mongodb) List(
	ctx contextx.Contextx,
	cond repo.ListCondition,
) (items []*model.Notification, total int, err error) {
	ctx, span := otelx.Span(ctx, "biz.notification.repo.notification.mongodb.List")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	filter := bson.M{}
	if cond.UserID != "" {
		filter["user_id"] = cond.UserID
	}

	opts := options.Find()
	if cond.Limit > 0 {
		opts.SetLimit(int64(cond.Limit))
	}
	if cond.Offset > 0 {
		opts.SetSkip(int64(cond.Offset))
	}

	cursor, err := i.rw.Database(dbName).Collection(collName).Find(timeout, filter, opts)
	if err != nil {
		ctx.Error("find notifications from mongodb failed", zap.Error(err))
		return nil, 0, err
	}
	defer cursor.Close(timeout)

	err = cursor.All(timeout, &items)
	if err != nil {
		ctx.Error("decode notifications from mongodb failed", zap.Error(err))
		return nil, 0, err
	}

	count, err := i.rw.Database(dbName).Collection(collName).CountDocuments(timeout, filter)
	if err != nil {
		ctx.Error("count notifications from mongodb failed", zap.Error(err))
		return nil, 0, err
	}

	return items, int(count), nil
}

func (i *mongodb) UpdateStatus(ctx contextx.Contextx, id, status string) error {
	ctx, span := otelx.Span(ctx, "biz.notification.repo.notification.mongodb.UpdateStatus")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": status, "updated_at": time.Now()}}

	_, err := i.rw.Database(dbName).Collection(collName).UpdateOne(timeout, filter, update)
	if err != nil {
		ctx.Error(
			"update notification status in mongodb failed",
			zap.Error(err),
			zap.String("id", id),
			zap.String("status", status),
		)
		return err
	}

	return nil
}
