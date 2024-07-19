package biz

import (
	"testing"

	logisticsB "github.com/blackhorseya/godine/entity/domain/logistics/biz"
	notifyB "github.com/blackhorseya/godine/entity/domain/notification/biz"
	"github.com/blackhorseya/godine/entity/domain/order/biz"
	"github.com/blackhorseya/godine/entity/domain/order/repo"
	rB "github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	userB "github.com/blackhorseya/godine/entity/domain/user/biz"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type suiteTester struct {
	suite.Suite

	ctrl       *gomock.Controller
	restaurant *rB.MockIRestaurantBiz
	menu       *rB.MockIMenuBiz
	user       *userB.MockIUserBiz
	logistics  *logisticsB.MockILogisticsBiz
	notify     *notifyB.MockINotificationBiz
	orders     *repo.MockIOrderRepo
	biz        biz.IOrderBiz
}

func (s *suiteTester) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.restaurant = rB.NewMockIRestaurantBiz(s.ctrl)
	s.menu = rB.NewMockIMenuBiz(s.ctrl)
	s.user = userB.NewMockIUserBiz(s.ctrl)
	s.logistics = logisticsB.NewMockILogisticsBiz(s.ctrl)
	s.notify = notifyB.NewMockINotificationBiz(s.ctrl)
	s.orders = repo.NewMockIOrderRepo(s.ctrl)
	s.biz = NewOrderBiz(
		s.restaurant,
		s.menu,
		s.user,
		s.logistics,
		s.notify,
		s.orders,
	)
}

func (s *suiteTester) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}
