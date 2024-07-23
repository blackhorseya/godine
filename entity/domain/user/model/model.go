package model

import (
	"errors"

	"github.com/blackhorseya/godine/pkg/contextx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// Roles is the list of roles the user has.
	Roles []*UserRole `json:"roles,omitempty" bson:"roles"`
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

// GetSubject returns the subject of the user.
func (x *User) GetSubject() string {
	return x.ID
}

// UpdateAddress updates the user's address.
func (x *User) UpdateAddress(newAddress Address) {
	x.Address = newAddress
}

func (x *User) UnmarshalBSON(bytes []byte) error {
	type Alias User
	alias := &struct {
		ID     primitive.ObjectID `bson:"_id"`
		*Alias `bson:",inline"`
	}{
		Alias: (*Alias)(x),
	}

	if err := bson.Unmarshal(bytes, alias); err != nil {
		return err
	}

	x.ID = alias.ID.Hex()

	return nil
}

func (x *User) MarshalBSON() ([]byte, error) {
	type Alias User
	alias := &struct {
		ID     primitive.ObjectID `bson:"_id"`
		*Alias `bson:",inline"`
	}{
		Alias: (*Alias)(x),
	}

	id, err := primitive.ObjectIDFromHex(x.ID)
	if err != nil {
		return nil, err
	}
	alias.ID = id

	return bson.Marshal(alias)
}

// UserRole represents the role of a user.
type UserRole struct {
	// UserID is the ID of the user.
	UserID string `json:"user_id,omitempty" bson:"user_id"`

	// RestaurantID is the ID of the restaurant.
	RestaurantID string `json:"restaurant_id,omitempty" bson:"restaurant_id"`

	// Role is the role of the user.
	Role RoleType `json:"role,omitempty" bson:"role"`
}
