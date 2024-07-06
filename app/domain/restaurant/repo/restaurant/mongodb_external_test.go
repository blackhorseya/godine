//go:build external

package restaurant

import (
	"testing"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/storage/mongodbx"
	"github.com/blackhorseya/godine/app/infra/storage/redix"
	model2 "github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/entity/domain/restaurant/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/logging"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type suiteMongodbExternal struct {
	suite.Suite

	rw   *mongo.Client
	rdb  *redis.Client
	repo repo.IRestaurantRepo
}

func (s *suiteMongodbExternal) SetupTest() {
	app, err := configx.NewApplication(viper.GetViper(), "restaurantRestful")
	s.Require().NoError(err)

	err = logging.Init(app.Log)
	s.Require().NoError(err)

	rw, err := mongodbx.NewClient(app)
	s.Require().NoError(err)
	s.rw = rw

	rdb, err := redix.NewRedis(app)
	s.Require().NoError(err)
	s.rdb = rdb

	s.repo = NewMongodb(s.rw, s.rdb)
}

func (s *suiteMongodbExternal) TearDownTest() {
	if s.rw != nil {
		_ = s.rw.Disconnect(contextx.Background())
	}

	if s.rdb != nil {
		_ = s.rdb.Close()
	}
}

func TestMongodbExternal(t *testing.T) {
	suite.Run(t, new(suiteMongodbExternal))
}

func (s *suiteMongodbExternal) Test_mongodb_Create() {
	restaurant := model2.NewRestaurant("test", model2.Address{})

	type args struct {
		ctx  contextx.Contextx
		data *model2.Restaurant
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "create success",
			args:    args{data: restaurant},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.Background()
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.repo.Create(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			_, _ = s.rw.Database(dbName).Collection(collName).DeleteMany(contextx.Background(), bson.M{})
		})
	}
}
