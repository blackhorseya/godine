package order

import (
	"fmt"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/order/model"
	"github.com/blackhorseya/godine/entity/domain/order/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type mariadb struct {
	rw *gorm.DB
}

// NewMariadb create and return a new mariadb
func NewMariadb(rw *gorm.DB) (repo.IOrderRepo, error) {
	err := rw.AutoMigrate(&model.Order{}, &model.OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("migrate order and order item failed: %w", err)
	}

	return &mariadb{rw: rw}, nil
}

func (i *mariadb) Create(ctx contextx.Contextx, order *model.Order) error {
	ctx, span := otelx.Span(ctx, "biz.order.repo.order.mariadb.Create")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	if order.ID == "" {
		order.ID = uuid.New().String()
	}
	err := i.rw.WithContext(timeout).Create(order).Error
	if err != nil {
		ctx.Error("create order to mariadb failed", zap.Error(err), zap.Any("order", &order))
		return err
	}

	return nil
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
