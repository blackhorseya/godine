package order

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/order/model"
	"github.com/blackhorseya/godine/entity/domain/order/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	defaultLimit   = 100
	defaultTimeout = 5 * time.Second
)

type gormDB struct {
	rw   *gorm.DB
	node *snowflake.Node
}

// NewMariadb create and return a new gormDB
func NewMariadb(rw *gorm.DB, node *snowflake.Node) (repo.IOrderRepo, error) {
	err := rw.AutoMigrate(&model.Order{}, &model.OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("migrate order and order item failed: %w", err)
	}

	return &gormDB{rw: rw, node: node}, nil
}

func (i *gormDB) Create(ctx contextx.Contextx, order *model.Order) (err error) {
	ctx, span := otelx.Span(ctx, "biz.order.repo.order.gormDB.Create")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	// 检查订单 ID
	if order.Id == 0 {
		order.Id = i.node.Generate().Int64()
	}

	// 创建订单
	err = i.rw.WithContext(timeout).Create(order).Error
	if err != nil {
		ctx.Error("create order to gormDB failed", zap.Error(err), zap.Any("order", &order))
		return err
	}

	return nil
}

func (i *gormDB) GetByID(ctx contextx.Contextx, id string) (item *model.Order, err error) {
	ctx, span := otelx.Span(ctx, "biz.order.repo.order.gormDB.GetByID")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	// 初始化 Order 对象
	order := &model.Order{}

	// 根据 ID 查询订单
	err = i.rw.WithContext(timeout).Preload("Items").First(order, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.Error("order not found", zap.String("id", id))
			return nil, errorx.Wrap(http.StatusNotFound, 404, err)
		}

		ctx.Error("get order by id from gormDB failed", zap.Error(err), zap.String("id", id))
		return nil, err
	}

	return order, nil
}

func (i *gormDB) List(
	ctx contextx.Contextx,
	condition repo.ListCondition,
) (items []*model.Order, total int, err error) {
	ctx, span := otelx.Span(ctx, "biz.order.repo.order.gormDB.List")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	var orders []*model.Order
	query := i.rw.WithContext(timeout).Model(&model.Order{})

	// Apply filters based on the condition
	if condition.UserID != "" {
		query = query.Where("user_id = ?", condition.UserID)
	}
	if condition.RestaurantID != "" {
		query = query.Where("restaurant_id = ?", condition.RestaurantID)
	}
	if condition.Status != "" {
		query = query.Where("status = ?", condition.Status)
	}

	// Apply limit and offset
	if condition.Limit == 0 {
		condition.Limit = defaultLimit
	}
	if condition.Offset < 0 {
		condition.Offset = 0
	}
	query = query.Limit(condition.Limit).Offset(condition.Offset)

	// Order by updated_at descending
	query = query.Order("updated_at DESC")

	// Execute the query
	var count int64
	err = query.Count(&count).Find(&orders).Error
	if err != nil {
		ctx.Error("list orders from gormDB failed", zap.Error(err))
		return nil, 0, err
	}

	return orders, int(count), nil
}

func (i *gormDB) Update(ctx contextx.Contextx, order *model.Order) (err error) {
	ctx, span := otelx.Span(ctx, "biz.order.repo.order.gormDB.Update")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	// 更新订单
	err = i.rw.WithContext(timeout).Save(order).Error
	if err != nil {
		ctx.Error("update order to gormDB failed", zap.Error(err), zap.Any("order", &order))
		return err
	}

	return nil
}
