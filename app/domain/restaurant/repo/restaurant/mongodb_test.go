package restaurant

import (
	"context"
	"testing"

	"github.com/blackhorseya/godine/app/infra/storage/mongodbx"
	"github.com/blackhorseya/godine/app/infra/storage/redix"
	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/entity/domain/restaurant/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type suiteMongodbTester struct {
	suite.Suite

	mongodbContainer *mongodbx.Container
	redisContainer   *redix.Container
	rw               *mongo.Client
	rdb              *redis.Client
	repo             repo.IRestaurantRepo
}

func (s *suiteMongodbTester) SetupTest() {
	mongodbC, err := mongodbx.NewContainer(contextx.Background())
	s.Require().NoError(err)
	s.mongodbContainer = mongodbC

	rw, err := mongodbC.RW(contextx.Background())
	s.Require().NoError(err)
	s.rw = rw

	redisC, err := redix.NewContainer(contextx.Background())
	s.Require().NoError(err)
	s.redisContainer = redisC

	rdb, err := redisC.RW(contextx.Background())
	s.Require().NoError(err)
	s.rdb = rdb

	s.repo = NewMongodb(s.rw, s.rdb)
}

func (s *suiteMongodbTester) TearDownTest() {
	if s.rw != nil {
		_ = s.rw.Disconnect(contextx.Background())
	}

	if s.rdb != nil {
		_ = s.rdb.Close()
	}

	if s.mongodbContainer != nil {
		_ = s.mongodbContainer.Terminate(contextx.Background())
	}

	if s.redisContainer != nil {
		_ = s.redisContainer.Terminate(contextx.Background())
	}
}

func TestMongodb(t *testing.T) {
	suite.Run(t, new(suiteMongodbTester))
}

func (s *suiteMongodbTester) Test_mongodb_Create() {
	type args struct {
		ctx  context.Context
		data *model.Restaurant
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "create success",
			args:    args{data: &model.Restaurant{Name: "test restaurant"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = context.WithValue(context.Background(), contextx.KeyContextx, contextx.Background())
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.repo.Create(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
