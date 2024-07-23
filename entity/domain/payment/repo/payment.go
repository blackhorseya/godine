//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/godine/entity/domain/payment/model"
	"github.com/blackhorseya/godine/pkg/contextx"
)

// ListCondition is a condition for list payment
type ListCondition struct {
	Offset int
	Limit  int
}

// IPaymentRepo is an interface for payment repository
type IPaymentRepo interface {
	GetByID(ctx contextx.Contextx, id string) (item *model.Payment, err error)
	List(ctx contextx.Contextx, cond ListCondition) (items []*model.Payment, total int, err error)
	Create(ctx contextx.Contextx, item *model.Payment) (err error)
	Update(ctx contextx.Contextx, item *model.Payment) (err error)
}
