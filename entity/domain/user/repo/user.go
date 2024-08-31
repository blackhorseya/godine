//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/pkg/utils"
)

// IUserRepo is an interface for user repository.
type IUserRepo interface {
	utils.IRepository[*model.Account]
}
