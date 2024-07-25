package biz

import (
	"github.com/blackhorseya/godine/app/infra/authz"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/user/biz"
	"github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/entity/domain/user/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
)

type userBiz struct {
	authz *authz.Authz
	users repo.IUserRepo
}

// NewUserBiz create and return a new user biz
func NewUserBiz(authz *authz.Authz, users repo.IUserRepo) biz.IUserBiz {
	return &userBiz{
		authz: authz,
		users: users,
	}
}

func (i *userBiz) Register(ctx contextx.Contextx, name, email string) (item *model.User, err error) {
	// todo: 2024/7/25|sean|implement me
	panic("implement me")
}

func (i *userBiz) Login(ctx contextx.Contextx) (item *model.User, err error) {
	// todo: 2024/7/25|sean|implement me
	panic("implement me")
}

func (i *userBiz) CreateUser(
	ctx contextx.Contextx,
	name, email, password string,
	address model.Address,
) (item *model.User, err error) {
	ctx, span := otelx.Span(ctx, "userBiz.CreateUser")
	defer span.End()

	user := model.NewUser(name, email, password, address)
	err = i.users.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (i *userBiz) GetUser(ctx contextx.Contextx, id string) (item *model.User, err error) {
	ctx, span := otelx.Span(ctx, "userBiz.GetUser")
	defer span.End()

	return i.users.GetByID(ctx, id)
}

func (i *userBiz) ListUsers(
	ctx contextx.Contextx,
	options biz.ListUsersOptions,
) (items []*model.User, total int, err error) {
	ctx, span := otelx.Span(ctx, "userBiz.ListUsers")
	defer span.End()

	return i.users.List(ctx, repo.ListCondition{
		Limit:  options.Size,
		Offset: (options.Page - 1) * options.Size,
	})
}

func (i *userBiz) UpdateUser(
	ctx contextx.Contextx,
	id string,
	name, email, password string,
	address model.Address,
) error {
	ctx, span := otelx.Span(ctx, "userBiz.UpdateUser")
	defer span.End()

	user, err := i.users.GetByID(ctx, id)
	if err != nil {
		return err
	}

	user.Name = name
	user.Email = email
	user.Password = password
	user.Address = address

	return i.users.Update(ctx, user)
}

func (i *userBiz) DeleteUser(ctx contextx.Contextx, id string) error {
	ctx, span := otelx.Span(ctx, "userBiz.DeleteUser")
	defer span.End()

	return i.users.Delete(ctx, id)
}

func (i *userBiz) ChangeUserStatus(ctx contextx.Contextx, userID string, isActive bool) error {
	ctx, span := otelx.Span(ctx, "userBiz.ChangeUserStatus")
	defer span.End()

	user, err := i.users.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	user.IsActive = isActive

	return i.users.Update(ctx, user)
}
