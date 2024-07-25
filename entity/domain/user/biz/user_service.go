//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
)

// ListUsersOptions defines the options for listing users.
type ListUsersOptions struct {
	// Page is the page number.
	Page int `form:"page" default:"1" minimum:"1"`

	// Size is the number of items per page.
	Size int `form:"size" default:"10" minimum:"1" maximum:"100"`
}

// IUserBiz defines the business operations for user management.
type IUserBiz interface {
	// Register registers a new user with the provided name, email, and password.
	Register(ctx contextx.Contextx, name string) (item *model.User, err error)

	// Login authenticates a user with the provided email and password.
	Login(ctx contextx.Contextx) (item *model.User, err error)

	// CreateUser creates a new user with the provided name, email, password, and address.
	CreateUser(ctx contextx.Contextx, name, email, password string, address model.Address) (item *model.User, err error)

	// GetUser retrieves the user with the specified ID.
	GetUser(ctx contextx.Contextx, id string) (item *model.User, err error)

	// ListUsers retrieves a list of users based on the provided options.
	ListUsers(ctx contextx.Contextx, options ListUsersOptions) (items []*model.User, total int, err error)

	// UpdateUser updates the details of an existing user.
	UpdateUser(ctx contextx.Contextx, id string, name, email, password string, address model.Address) error

	// DeleteUser deletes a user by their ID.
	DeleteUser(ctx contextx.Contextx, id string) error

	// ChangeUserStatus changes the active status of a user.
	ChangeUserStatus(ctx contextx.Contextx, userID string, isActive bool) error
}
