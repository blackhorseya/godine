package model

import (
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
