package mongodbx

import (
	"context"
	"testing"

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
