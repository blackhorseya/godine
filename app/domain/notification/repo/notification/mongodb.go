package notification

import (
	"time"

	"github.com/blackhorseya/godine/entity/notification/model"
	"github.com/blackhorseya/godine/entity/notification/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.mongodb.org/mongo-driver/mongo"
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
	// todo: 2024/6/26|sean|implement me
	panic("implement me")
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
