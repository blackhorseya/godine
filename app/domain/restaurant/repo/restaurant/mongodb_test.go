package restaurant

import (
	"testing"

	"github.com/blackhorseya/godine/app/infra/storage/mongodbx"
	"github.com/blackhorseya/godine/app/infra/storage/redix"
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
