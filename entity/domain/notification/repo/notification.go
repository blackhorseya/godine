//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"context"

	"github.com/blackhorseya/godine/entity/domain/notification/model"
	"github.com/blackhorseya/godine/pkg/utils"
)

// INotificationRepo defines the interface for notification repository operations.
type INotificationRepo interface {
	utils.IRepository[*model.Notification]

	ListByReceiverID(
		c context.Context,
		receiverID string,
		cond utils.Pagination,
	) (items []*model.Notification, total int64, err error)
}
