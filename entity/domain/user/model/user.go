package model

import (
	"errors"

	"github.com/blackhorseya/godine/pkg/contextx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewUser creates and returns a new user.
func NewUser(name, email, password string, address *Address) *Account {
	return &Account{
		Id:          "",
		Username:    "",
		Email:       email,
		Password:    password,
		Address:     address,
		IsActive:    true,
		Level:       0,
		Roles:       nil,
		SocialIDMap: nil,
		CreatedAt:   timestamppb.Now(),
		UpdatedAt:   timestamppb.Now(),
	}
}

// FromContext extracts the user from the context.
func FromContext(ctx contextx.Contextx) (*Account, error) {
	user, ok := ctx.Value(contextx.KeyHandler{}).(*Account)
	if !ok {
		return nil, errors.New("no user found in context")
	}

	return user, nil
}
