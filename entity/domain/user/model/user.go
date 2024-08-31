package model

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type keyHandler struct{}

// FromContext extracts the user from the context.
func FromContext(c context.Context) (*Account, error) {
	account, ok := c.Value(keyHandler{}).(*Account)
	if !ok {
		return nil, errors.New("no user found in context")
	}

	return account, nil
}

// SetInContext sets the user in the context.
func (x *Account) SetInContext(c context.Context) context.Context {
	return context.WithValue(c, keyHandler{}, x)
}

func (x *Account) GetID() string {
	return x.Id
}

func (x *Account) SetID(id primitive.ObjectID) {
	x.Id = id.Hex()
}

func (x *Account) SetCreatedAt(t *timestamppb.Timestamp) {
	x.CreatedAt = t
}

func (x *Account) SetUpdatedAt(t *timestamppb.Timestamp) {
	x.UpdatedAt = t
}
