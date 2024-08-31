//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
)

type ListCondition struct {
	Limit  int64
	Offset int64
}

// IUserRepoLegacy is an interface for user repository.
type IUserRepoLegacy interface {
	Create(ctx contextx.Contextx, user *model.Account) error
	GetByID(ctx contextx.Contextx, id string) (item *model.Account, err error)
	List(ctx contextx.Contextx, condition ListCondition) (items []*model.Account, total int, err error)
	Update(ctx contextx.Contextx, user *model.Account) error
	Delete(ctx contextx.Contextx, id string) error
}
