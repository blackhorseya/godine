//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/godine/entity/domain/payment/model"
	"github.com/blackhorseya/godine/pkg/utils"
)

// IPaymentRepo is an interface for payment repository
type IPaymentRepo interface {
	utils.IRepository[*model.Payment]
}
