package domain

import (
	"context"
	"errors"
	"testing"

	"github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/entity/domain/restaurant/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type suiteRestaurantServiceTester struct {
	suite.Suite

	ctrl *gomock.Controller
	repo *repo.MockIRestaurantRepo
	biz  biz.RestaurantServiceServer
}

func (s *suiteRestaurantServiceTester) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.repo = repo.NewMockIRestaurantRepo(s.ctrl)
	s.biz = NewRestaurantService(s.repo)
}

func (s *suiteRestaurantServiceTester) TearDownTest() {
	if s.ctrl != nil {
		s.ctrl.Finish()
	}
}

func TestAllRestaurantServiceTester(t *testing.T) {
	suite.Run(t, new(suiteRestaurantServiceTester))
}

func (s *suiteRestaurantServiceTester) Test_restaurantService_PlaceOrder() {
	type args struct {
		c    context.Context
		req  *biz.PlaceOrderRequest
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		want    *biz.PlaceOrderResponse
		wantErr bool
	}{
		{
			name: "not found restaurant",
			args: args{req: &biz.PlaceOrderRequest{RestaurantId: "not_found"}, mock: func() {
				s.repo.EXPECT().GetByID(gomock.Any(), "not_found").Return(nil, errors.New("not found")).Times(1)
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "reservation failed",
			args: args{req: &biz.PlaceOrderRequest{RestaurantId: "found"}, mock: func() {
				s.repo.EXPECT().GetByID(gomock.Any(), "found").Return(&model.Restaurant{Id: "found"}, nil).Times(1)
				s.repo.EXPECT().CreateReservation(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(errors.New("create reservation failed")).
					Times(1)
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "reservation success",
			args: args{req: &biz.PlaceOrderRequest{RestaurantId: "found"}, mock: func() {
				s.repo.EXPECT().GetByID(gomock.Any(), "found").Return(&model.Restaurant{Id: "found"}, nil).Times(1)
				s.repo.EXPECT().CreateReservation(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
			}},
			want:    &biz.PlaceOrderResponse{OrderId: "", Status: "reserved"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.c = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			got, err := s.biz.PlaceOrder(tt.args.c, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaceOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && got.Status != tt.want.Status {
				t.Errorf("PlaceOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}
