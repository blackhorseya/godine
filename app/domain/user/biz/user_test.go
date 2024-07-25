package biz

import (
	"testing"

	"github.com/blackhorseya/godine/app/infra/authz"
	"github.com/blackhorseya/godine/entity/domain/user/biz"
	"github.com/blackhorseya/godine/entity/domain/user/repo"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type suiteTester struct {
	suite.Suite

	ctrl  *gomock.Controller
	authz *authz.Authz
	users *repo.MockIUserRepo
	biz   biz.IUserBiz
}

func (s *suiteTester) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.authz = new(authz.Authz)
	s.users = repo.NewMockIUserRepo(s.ctrl)
	s.biz = NewUserBiz(s.authz, s.users)
}

func (s *suiteTester) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}
