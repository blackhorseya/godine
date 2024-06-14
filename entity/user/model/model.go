package model

import (
	"github.com/google/uuid"
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
}

// NewUser creates and returns a new user.
func NewUser(name, email, password string, address Address) *User {
	return &User{
		ID:       uuid.New().String(),
		Name:     name,
		Email:    email,
		Password: password,
		Address:  address,
		IsActive: true,
		Level:    0,
	}
}

// UpdateAddress updates the user's address.
func (x *User) UpdateAddress(newAddress Address) {
	x.Address = newAddress
}
