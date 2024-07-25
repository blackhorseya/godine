package biz

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/godine/app/infra/authz"
	"github.com/blackhorseya/godine/entity/domain/user/biz"
	"github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/entity/domain/user/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
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

func (s *suiteTester) Test_userBiz_Register() {
	type args struct {
		ctx  contextx.Contextx
		name string
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantItem *model.User
		wantErr  bool
	}{
		{
			name: "successful registration",
			args: args{
				name: "testUser",
				mock: func() {
					s.users.EXPECT().Create(gomock.Any(), &model.User{Name: "testUser"}).Return(nil)
				},
			},
			wantItem: &model.User{Name: "testUser"},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			tt.args.ctx = contextx.WithValue(contextx.Background(), contextx.KeyHandler, &model.User{Name: "testUser"})
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotItem, err := s.biz.Register(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("Register() gotItem = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}
