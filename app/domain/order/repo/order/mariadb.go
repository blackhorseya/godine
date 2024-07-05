package order

import (
	"github.com/blackhorseya/godine/entity/domain/order/model"
	"github.com/blackhorseya/godine/entity/domain/order/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/jmoiron/sqlx"
)

type mariadb struct {
	rw *sqlx.DB
}

// NewMariadb create and return a new mariadb
func NewMariadb(rw *sqlx.DB) repo.IOrderRepo {
	return &mariadb{rw: rw}
}

func (i *mariadb) Create(ctx contextx.Contextx, order *model.Order) error {
	// todo: 2024/7/5|sean|implement me
	panic("implement me")
}

func (i *mariadb) GetByID(ctx contextx.Contextx, id string) (item *model.Order, err error) {
	// todo: 2024/7/5|sean|implement me
	panic("implement me")
}

func (i *mariadb) List(
	ctx contextx.Contextx,
	condition repo.ListCondition,
) (items []*model.Order, total int, err error) {
	// todo: 2024/7/5|sean|implement me
	panic("implement me")
}

func (i *mariadb) Update(ctx contextx.Contextx, order *model.Order) error {
	// todo: 2024/7/5|sean|implement me
	panic("implement me")
}
