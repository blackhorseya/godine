package biz

import (
	"testing"

	"github.com/blackhorseya/godine/app/infra/authz"
	"github.com/blackhorseya/godine/entity/domain/user/biz"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type suiteTester struct {
	suite.Suite

	ctrl  *gomock.Controller
	authz *authz.Authz
	biz   biz.IUserBiz
}

func (s *suiteTester) SetupTest() {
}

func (s *suiteTester) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}
