package order

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/order/model"
	"github.com/blackhorseya/godine/entity/domain/order/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
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

	// 开启事务
	tx := i.rw.WithContext(timeout).Begin()
	if tx.Error != nil {
		ctx.Error("failed to begin transaction", zap.Error(tx.Error))
		return tx.Error
	}

	// 检查订单 ID
	if order.ID == "" {
		order.ID = uuid.New().String()
	}

	// 创建订单
	err := tx.Create(order).Error
	if err != nil {
		tx.Rollback()
		ctx.Error("create order to mariadb failed", zap.Error(err), zap.Any("order", &order))
		return err
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		ctx.Error("failed to commit transaction", zap.Error(err))
		return err
	}

	return nil
}

func (i *mariadb) GetByID(ctx contextx.Contextx, id string) (item *model.Order, err error) {
	ctx, span := otelx.Span(ctx, "biz.order.repo.order.mariadb.GetByID")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	// 初始化 Order 对象
	order := &model.Order{}

	// 根据 ID 查询订单
	err = i.rw.WithContext(timeout).Preload("Items").First(order, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.Warn("order not found", zap.String("id", id))
			return nil, errorx.Wrap(http.StatusNotFound, 404, err)
		}

		ctx.Error("get order by id from mariadb failed", zap.Error(err), zap.String("id", id))
		return nil, err
	}

	return order, nil
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
