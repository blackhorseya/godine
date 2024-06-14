package model

import (
	"github.com/google/uuid"
)

// User represents a user entity.
type User struct {
	// ID is the unique identifier of the user.
	ID uuid.UUID `json:"id,omitempty" bson:"_id,omitempty"`

	// Name is the name of the user.
	Name string `json:"name,omitempty" bson:"name"`

	// Email is the email address of the user.
	Email string `json:"email,omitempty" bson:"email"`

	// Password is the hashed password of the user.
	Password string `json:"password,omitempty" bson:"password"`

	// Address is the address of the user.
	Address Address `json:"address,omitempty" bson:"address"`
}

// NewUser creates and returns a new user.
func NewUser(name, email, password string, address Address) *User {
	return &User{
		ID:       uuid.New(),
		Name:     name,
		Email:    email,
		Password: password,
		Address:  address,
	}
}

// UpdateAddress updates the user's address.
func (x *User) UpdateAddress(newAddress Address) {
	x.Address = newAddress
}
