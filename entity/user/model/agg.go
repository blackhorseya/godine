package model

import (
	"github.com/google/uuid"
)

// UserAggregate represents the user aggregate root.
type UserAggregate struct {
	// User is the core entity of the aggregate.
	User User `json:"user" bson:"user"`
}

// NewUser creates a new UserAggregate.
func NewUser(name, email, password string, address Address) *UserAggregate {
	return &UserAggregate{
		User: User{
			ID:       uuid.New(),
			Name:     name,
			Email:    email,
			Password: password,
			Address:  address,
		},
	}
}

// UpdateAddress updates the user's address.
func (ua *UserAggregate) UpdateAddress(newAddress Address) {
	ua.User.Address = newAddress
}
