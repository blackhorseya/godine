package biz

import (
	"errors"
	"reflect"
	"testing"

	logisticsB "github.com/blackhorseya/godine/entity/domain/logistics/biz"
	notifyB "github.com/blackhorseya/godine/entity/domain/notification/biz"
	"github.com/blackhorseya/godine/entity/domain/order/biz"
	model4 "github.com/blackhorseya/godine/entity/domain/order/model"
	"github.com/blackhorseya/godine/entity/domain/order/repo"
	rB "github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	userB "github.com/blackhorseya/godine/entity/domain/user/biz"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

var errMock = errors.New("mock error")

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

func (s *suiteTester) Test_orderBiz_GetOrder() {
	type args struct {
		ctx  contextx.Contextx
		id   string
		mock func()
	}
	tests := []struct {
		name      string
		args      args
		wantOrder *model4.Order
		wantErr   bool
	}{
		{
			name: "get order failed",
			args: args{id: "1", mock: func() {
				s.orders.EXPECT().GetByID(gomock.Any(), "1").Return(nil, errMock)
			}},
			wantOrder: nil,
			wantErr:   true,
		},
		{
			name: "get order success",
			args: args{id: "1", mock: func() {
				s.orders.EXPECT().GetByID(gomock.Any(), "1").Return(&model4.Order{Id: 1}, nil)
			}},
			wantOrder: &model4.Order{Id: 1},
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotOrder, err := s.biz.GetOrder(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrder, tt.wantOrder) {
				t.Errorf("GetOrder() gotOrder = %v, want %v", gotOrder, tt.wantOrder)
			}
		})
	}
}
