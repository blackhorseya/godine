//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/godine/entity/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
)

type ListCondition struct {
	Limit  int
	Offset int
}

// IUserRepo is an interface for user repository.
type IUserRepo interface {
	Create(ctx contextx.Contextx, user *model.User) error
	GetByID(ctx contextx.Contextx, id string) (item *model.User, err error)
	List(ctx contextx.Contextx, condition ListCondition) (items []*model.User, total int, err error)
	Update(ctx contextx.Contextx, user *model.User) error
	Delete(ctx contextx.Contextx, id string) error
}
