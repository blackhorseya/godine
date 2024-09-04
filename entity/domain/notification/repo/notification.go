//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"context"

	"github.com/blackhorseya/godine/entity/domain/notification/model"
	"github.com/blackhorseya/godine/pkg/persistence"
)

// INotificationRepo defines the interface for notification repository operations.
type INotificationRepo interface {
	persistence.IRepository[*model.Notification]

	ListByReceiverID(
		c context.Context,
		receiverID string,
		cond persistence.Pagination,
	) (items []*model.Notification, total int64, err error)
}
