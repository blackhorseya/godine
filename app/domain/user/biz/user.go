package biz

import (
	"github.com/blackhorseya/godine/entity/user/biz"
	"github.com/blackhorseya/godine/entity/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/google/uuid"
)

type userBiz struct {
}

// NewUserBiz create and return a new user biz
func NewUserBiz() biz.IUserBiz {
	return &userBiz{}
}

func (i *userBiz) CreateUser(
	ctx contextx.Contextx,
	name, email, password string,
	address model.Address,
) (item *model.User, err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *userBiz) GetUser(ctx contextx.Contextx, id uuid.UUID) (item *model.User, err error) {
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
	id uuid.UUID,
	name, email, password string,
	address model.Address,
) error {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *userBiz) DeleteUser(ctx contextx.Contextx, id uuid.UUID) error {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *userBiz) SearchUsers(ctx contextx.Contextx, keyword string) (items []model.User, total int, err error) {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}

func (i *userBiz) ChangeUserStatus(ctx contextx.Contextx, userID uuid.UUID, isActive bool) error {
	// todo: 2024/6/11|sean|implement me
	panic("implement me")
}
