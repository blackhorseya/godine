//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/godine/entity/domain/logistics/model"
	"github.com/blackhorseya/godine/pkg/utils"
)

// IDeliveryRepo represents a repository interface for managing delivery entities.
type IDeliveryRepo interface {
	utils.IRepository[*model.Delivery]
}
