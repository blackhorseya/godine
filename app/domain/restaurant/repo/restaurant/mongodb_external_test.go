//go:build external

package restaurant

import (
	"testing"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/storage/mongodbx"
	"github.com/blackhorseya/godine/entity/restaurant/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/logging"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type suiteMongodbExternal struct {
	suite.Suite

	rw   *mongo.Client
	repo repo.IRestaurantRepo
}

func (s *suiteMongodbExternal) SetupTest() {
	err := configx.LoadConfig("")
	s.Require().NoError(err)

	app, err := configx.LoadApplication(&configx.C.RestaurantRestful)
	s.Require().NoError(err)

	err = logging.Init(app.Log)
	s.Require().NoError(err)

	rw, err := mongodbx.NewClient(app)
	s.Require().NoError(err)
	s.rw = rw

	s.repo = NewMongodb(s.rw)
}

func (s *suiteMongodbExternal) TearDownTest() {
	if s.rw != nil {
		_ = s.rw.Disconnect(contextx.Background())
	}
}

func TestMongodbExternal(t *testing.T) {
	suite.Run(t, new(suiteMongodbExternal))
}
