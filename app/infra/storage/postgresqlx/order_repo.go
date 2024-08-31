package postgresqlx

import (
	"context"

	"github.com/blackhorseya/godine/entity/domain/order/model"
	"github.com/blackhorseya/godine/entity/domain/order/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type gormOrderRepo struct {
	rw   *gorm.DB
	node *snowflake.Node
}

// NewOrderRepo create and return a new gormOrderRepo.
func NewOrderRepo(rw *gorm.DB, node *snowflake.Node) repo.IOrderRepo {
	return &gormOrderRepo{
		rw:   rw,
		node: node,
	}
}

func (i *gormOrderRepo) Create(c context.Context, item *model.Order) error {
	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	ctx := contextx.Background()

	if item.Id == 0 {
		item.Id = i.node.Generate().Int64()
	}
	item.CreatedAt = timestamppb.Now()
	item.UpdatedAt = timestamppb.Now()

	err := i.rw.WithContext(timeout).Create(item).Error
	if err != nil {
		ctx.Error("create order to gormDB failed", zap.Error(err), zap.Any("order", &item))
		return err
	}

	return nil
}

func (i *gormOrderRepo) GetByID(c context.Context, id string) (item *model.Order, err error) {
	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	ctx := contextx.Background()

	err = i.rw.WithContext(timeout).Where("id = ?", id).First(item).Error
	if err != nil {
		ctx.Error("get order by id from gormDB failed", zap.Error(err), zap.String("id", id))
		return nil, err
	}

	return item, nil
}

func (i *gormOrderRepo) List(c context.Context, cond repo.ListCondition) (items []*model.Order, total int, err error) {
	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	ctx := contextx.Background()

	query := i.rw.WithContext(timeout).Model(&model.Order{})

	// limit and offset
	limit, offset := defaultLimit, 0
	if 0 < cond.Limit && cond.Limit <= defaultMaxLimit {
		limit = cond.Limit
	}
	if 0 < cond.Offset {
		offset = cond.Offset
	}
	query = query.Limit(limit).Offset(offset)

	// order by
	query = query.Order("updated_at desc")

	var count int64
	err = query.Count(&count).Find(&items).Error
	if err != nil {
		ctx.Error("list order from gormDB failed", zap.Error(err))
		return nil, 0, err
	}

	return items, int(count), nil
}

func (i *gormOrderRepo) Update(c context.Context, item *model.Order) error {
	timeout, cancelFunc := context.WithTimeout(c, defaultTimeout)
	defer cancelFunc()

	ctx := contextx.Background()

	item.UpdatedAt = timestamppb.Now()

	err := i.rw.WithContext(timeout).Save(item).Error
	if err != nil {
		ctx.Error("update order to gormDB failed", zap.Error(err), zap.Any("order", &item))
		return err
	}

	return nil
}
