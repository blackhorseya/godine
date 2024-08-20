package model

import (
	"errors"

	"github.com/blackhorseya/godine/pkg/contextx"
)

// User represents a user entity.
type User struct {
	// ID is the unique identifier of the user.
	ID string `json:"id,omitempty" bson:"_id,omitempty"`

	// Name is the name of the user.
	Name string `json:"name,omitempty" bson:"name"`

	// Email is the email address of the user.
	Email string `json:"email,omitempty" bson:"email"`

	// Password is the hashed password of the user.
	Password string `json:"password,omitempty" bson:"password"`

	// Address is the address of the user.
	Address Address `json:"address,omitempty" bson:"address"`

	// IsActive is the status of the user.
	IsActive bool `json:"is_active,omitempty" bson:"is_active"`

	// Level is the level of the user.
	Level uint `json:"level" bson:"level"`

	// SocialID is the social ID of the user.
	SocialID string `json:"social_id,omitempty" bson:"social_id"`
}

// NewUser creates and returns a new user.
func NewUser(name, email, password string, address Address) *User {
	return &User{
		ID:       "",
		Name:     name,
		Email:    email,
		Password: password,
		Address:  address,
		IsActive: true,
		Level:    0,
	}
}

// FromContext extracts the user from the context.
func FromContext(ctx contextx.Contextx) (*User, error) {
	user, ok := ctx.Value(contextx.KeyHandler).(*User)
	if !ok {
		return nil, errors.New("no user found in context")
	}

	return user, nil
}
