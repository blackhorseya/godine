//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/godine/entity/notification/model"
	"github.com/blackhorseya/godine/pkg/contextx"
)

// ListCondition defines the conditions for listing notifications.
type ListCondition struct {
	Limit  int
	Offset int

	UserID string
}

// INotificationRepo defines the interface for notification repository operations.
type INotificationRepo interface {
	Create(ctx contextx.Contextx, notify *model.Notification) error
	GetByID(ctx contextx.Contextx, id string) (item *model.Notification, err error)
	List(ctx contextx.Contextx, cond ListCondition) (items []*model.Notification, total int, err error)
	UpdateStatus(ctx contextx.Contextx, id, status string) error
}
