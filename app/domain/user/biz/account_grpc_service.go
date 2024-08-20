package biz

import (
	"context"

	userB "github.com/blackhorseya/godine/entity/domain/user/biz"
	"github.com/blackhorseya/godine/entity/domain/user/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

type accountService struct {
}

// NewAccountService creates and returns a new AccountService.
func NewAccountService() userB.AccountServiceServer {
	return &accountService{}
}

func (i *accountService) WhoAmI(c context.Context, empty *emptypb.Empty) (*model.Account, error) {
	// TODO: 2024/8/21|sean|implement me
	panic("implement me")
}
