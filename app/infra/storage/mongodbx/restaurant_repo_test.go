package mongodbx

import (
	"context"
	"testing"

	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/entity/domain/restaurant/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type suiteRestaurantRepoTester struct {
	suite.Suite

	db   *Container
	rw   *mongo.Client
	repo repo.IRestaurantRepo
}

func (s *suiteRestaurantRepoTester) SetupTest() {
	ctx := contextx.WithContextx(context.Background())

	db, err := NewContainer(ctx)
	s.Require().NoError(err)
	s.db = db

	rw, err := s.db.RW(ctx)
	s.Require().NoError(err)
	s.rw = rw

	s.repo = NewRestaurantRepo(s.rw)
}

func (s *suiteRestaurantRepoTester) TearDownTest() {
	if s.rw != nil {
		_ = s.rw.Disconnect(context.Background())
	}

	if s.db != nil {
		_ = s.db.Terminate(context.Background())
	}
}

func TestAllRestaurantRepoTester(t *testing.T) {
	suite.Run(t, new(suiteRestaurantRepoTester))
}

func (s *suiteRestaurantRepoTester) Test_mongodbRestaurantRepo_CreateReservation() {
	type args struct {
		c           context.Context
		restaurant  *model.Restaurant
		reservation *model.Order
		mock        func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "create reservation success",
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.c = context.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.repo.CreateReservation(
				tt.args.c,
				tt.args.restaurant,
				tt.args.reservation,
			); (err != nil) != tt.wantErr {
				t.Errorf("CreateReservation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
