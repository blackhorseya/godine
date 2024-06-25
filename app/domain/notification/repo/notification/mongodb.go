package notification

import (
	"time"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/notification/model"
	"github.com/blackhorseya/godine/entity/notification/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
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
		notify.ID = uuid.New().String()
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
	// todo: 2024/6/26|sean|implement me
	panic("implement me")
}

func (i *mongodb) List(
	ctx contextx.Contextx,
	cond repo.ListCondition,
) (items []*model.Notification, total int, err error) {
	// todo: 2024/6/26|sean|implement me
	panic("implement me")
}

func (i *mongodb) UpdateStatus(ctx contextx.Contextx, id, status string) error {
	// todo: 2024/6/26|sean|implement me
	panic("implement me")
}
