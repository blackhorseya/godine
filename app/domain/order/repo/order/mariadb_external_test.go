//go:build external

package order

import (
	"testing"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/snowflakex"
	"github.com/blackhorseya/godine/app/infra/storage/mariadbx"
	"github.com/blackhorseya/godine/entity/domain/order/model"
	"github.com/blackhorseya/godine/entity/domain/order/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/logging"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type mariadbExternalTester struct {
	suite.Suite

	rw   *gorm.DB
	repo repo.IOrderRepo
}

func (s *mariadbExternalTester) SetupTest() {
	err := configx.LoadConfig("")
	s.Require().NoError(err)

	app, err := configx.LoadApplication(&configx.C.OrderRestful)
	s.Require().NoError(err)

	err = logging.Init(app.Log)
	s.Require().NoError(err)

	rw, err := mariadbx.NewClient(app)
	s.Require().NoError(err)
	s.rw = rw

	node, err := snowflakex.NewNode()
	s.Require().NoError(err)

	orderRepo, err := NewMariadb(s.rw, node)
	s.Require().NoError(err)
	s.repo = orderRepo
}

func (s *mariadbExternalTester) TearDownTest() {
}

func TestMariadbExternal(t *testing.T) {
	suite.Run(t, new(mariadbExternalTester))
}

func (s *mariadbExternalTester) TestCreate() {
	order := model.NewOrder(primitive.NewObjectID().Hex(), primitive.NewObjectID().Hex(), []model.OrderItem{
		*model.NewOrderItem(primitive.NewObjectID().Hex(), "item 1", 10, 2),
		*model.NewOrderItem(primitive.NewObjectID().Hex(), "item 1", 20, 4),
	})

	ctx := contextx.Background()
	err := s.repo.Create(ctx, order)
	s.Require().NoError(err)

	ctx.Debug("create order success", zap.Any("order", &order))
}

func (s *mariadbExternalTester) TestGetByID() {
	order := model.NewOrder(primitive.NewObjectID().Hex(), primitive.NewObjectID().Hex(), []model.OrderItem{
		*model.NewOrderItem(primitive.NewObjectID().Hex(), "item 1", 10, 2),
		*model.NewOrderItem(primitive.NewObjectID().Hex(), "item 1", 20, 4),
	})

	ctx := contextx.Background()
	err := s.repo.Create(ctx, order)
	s.Require().NoError(err)

	ctx.Debug("create order success", zap.Any("order", &order))

	find, err := s.repo.GetByID(ctx, order.ID)
	s.Require().NoError(err)

	ctx.Debug("find order success", zap.Any("order", &find))
}

func (s *mariadbExternalTester) TestList() {
	order := model.NewOrder(primitive.NewObjectID().Hex(), primitive.NewObjectID().Hex(), []model.OrderItem{
		*model.NewOrderItem(primitive.NewObjectID().Hex(), "item 1", 10, 2),
		*model.NewOrderItem(primitive.NewObjectID().Hex(), "item 1", 20, 4),
	})

	ctx := contextx.Background()
	err := s.repo.Create(ctx, order)
	s.Require().NoError(err)

	ctx.Debug("create order success", zap.Any("order", &order))

	list, total, err := s.repo.List(ctx, repo.ListCondition{
		UserID:       "",
		RestaurantID: "",
		Status:       "",
		Limit:        0,
		Offset:       0,
	})
	s.Require().NoError(err)

	ctx.Debug("list order success", zap.Any("list", &list), zap.Int("total", total))
}
