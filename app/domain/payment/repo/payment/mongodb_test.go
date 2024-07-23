package payment

import (
	"testing"

	"github.com/blackhorseya/godine/app/infra/storage/mongodbx"
	"github.com/blackhorseya/godine/entity/domain/payment/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type suiteMongodbTester struct {
	suite.Suite

	mongodbContainer *mongodbx.Container
	rw               *mongo.Client
	repo             repo.IPaymentRepo
}

func (s *suiteMongodbTester) SetupTest() {
	mongodbC, err := mongodbx.NewContainer(contextx.Background())
	s.Require().NoError(err)
	s.mongodbContainer = mongodbC

	rw, err := mongodbC.RW(contextx.Background())
	s.Require().NoError(err)
	s.rw = rw

	s.repo = NewMongodb(s.rw)
}

func (s *suiteMongodbTester) TearDownTest() {
	if s.rw != nil {
		_ = s.rw.Disconnect(contextx.Background())
	}

	if s.mongodbContainer != nil {
		_ = s.mongodbContainer.Terminate(contextx.Background())
	}
}

func TestMongodb(t *testing.T) {
	suite.Run(t, new(suiteMongodbTester))
}
