package domain

import (
	"testing"

	"github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	"github.com/blackhorseya/godine/entity/domain/restaurant/repo"
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
