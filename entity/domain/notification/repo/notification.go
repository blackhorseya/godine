//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/godine/entity/domain/notification/model"
	"github.com/blackhorseya/godine/pkg/utils"
)

// INotificationRepo defines the interface for notification repository operations.
type INotificationRepo interface {
	utils.IRepository[*model.Notification]
}
