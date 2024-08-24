package biz

import (
	"context"

	"github.com/blackhorseya/godine/entity/domain/user/biz"
	"github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type accountService struct {
}

// NewAccountService creates and returns a new AccountService.
func NewAccountService() biz.AccountServiceServer {
	return &accountService{}
}

func (i *accountService) WhoAmI(c context.Context, empty *emptypb.Empty) (*model.Account, error) {
	ctx, err := contextx.FromContext(c)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get context: %v", err)
	}

	handler, err := model.FromContext(ctx)
	if err != nil {
		ctx.Error("failed to get user from context", zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return handler, nil
}
