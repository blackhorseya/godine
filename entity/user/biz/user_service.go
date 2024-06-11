//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/godine/entity/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/google/uuid"
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
	CreateUser(ctx contextx.Contextx, name, email, password string, address model.Address) (item *model.User, err error)

	// GetUser retrieves the user with the specified ID.
	GetUser(ctx contextx.Contextx, id uuid.UUID) (item *model.User, err error)

	// ListUsers retrieves a list of users based on the provided options.
	ListUsers(ctx contextx.Contextx, options ListUsersOptions) (items []model.User, total int, err error)

	// UpdateUser updates the details of an existing user.
	UpdateUser(ctx contextx.Contextx, id uuid.UUID, name, email, password string, address model.Address) error

	// DeleteUser deletes a user by their ID.
	DeleteUser(ctx contextx.Contextx, id uuid.UUID) error

	// SearchUsers searches for users by name or email keywords.
	SearchUsers(ctx contextx.Contextx, keyword string) (items []model.User, total int, err error)

	// ChangeUserStatus changes the active status of a user.
	ChangeUserStatus(ctx contextx.Contextx, userID uuid.UUID, isActive bool) error
}
