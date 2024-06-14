package biz

import (
	"github.com/blackhorseya/godine/entity/user/biz"
	"github.com/blackhorseya/godine/entity/user/model"
	"github.com/blackhorseya/godine/entity/user/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
)

type userBiz struct {
	users repo.IUserRepo
}

// NewUserBiz create and return a new user biz
func NewUserBiz(users repo.IUserRepo) biz.IUserBiz {
	return &userBiz{
		users: users,
	}
}

func (i *userBiz) CreateUser(
	ctx contextx.Contextx,
	name, email, password string,
	address model.Address,
) (item *model.User, err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *userBiz) GetUser(ctx contextx.Contextx, id string) (item *model.User, err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *userBiz) ListUsers(
	ctx contextx.Contextx,
	options biz.ListUsersOptions,
) (items []model.User, total int, err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *userBiz) UpdateUser(
	ctx contextx.Contextx,
	id string,
	name, email, password string,
	address model.Address,
) error {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *userBiz) DeleteUser(ctx contextx.Contextx, id string) error {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *userBiz) SearchUsers(ctx contextx.Contextx, keyword string) (items []model.User, total int, err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *userBiz) ChangeUserStatus(ctx contextx.Contextx, userID string, isActive bool) error {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}
