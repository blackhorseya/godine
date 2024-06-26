//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	model2 "github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
)

// ListUsersOptions defines the options for listing users.
type ListUsersOptions struct {
	// Page is the page number.
	Page int

	// PageSize is the number of items per page.
	PageSize int
}

// IUserBiz defines the business operations for user management.
type IUserBiz interface {
	// CreateUser creates a new user with the provided name, email, password, and address.
	CreateUser(ctx contextx.Contextx, name, email, password string, address model2.Address) (item *model2.User, err error)

	// GetUser retrieves the user with the specified ID.
	GetUser(ctx contextx.Contextx, id string) (item *model2.User, err error)

	// ListUsers retrieves a list of users based on the provided options.
	ListUsers(ctx contextx.Contextx, options ListUsersOptions) (items []*model2.User, total int, err error)

	// UpdateUser updates the details of an existing user.
	UpdateUser(ctx contextx.Contextx, id string, name, email, password string, address model2.Address) error

	// DeleteUser deletes a user by their ID.
	DeleteUser(ctx contextx.Contextx, id string) error

	// ChangeUserStatus changes the active status of a user.
	ChangeUserStatus(ctx contextx.Contextx, userID string, isActive bool) error
}
