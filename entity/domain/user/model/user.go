package model

import (
	"context"
	"errors"

	"github.com/blackhorseya/godine/pkg/contextx"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type keyHandler struct{}

// SetInContext sets the user in the context.
func (x *Account) SetInContext(c context.Context) context.Context {
	return context.WithValue(c, keyHandler{}, x)
}

// FromContextLegacy extracts the user from the context.
func FromContextLegacy(ctx contextx.Contextx) (*Account, error) {
	user, ok := ctx.Value(contextx.KeyHandler{}).(*Account)
	if !ok {
		return nil, errors.New("no user found in context")
	}

	return user, nil
}

// FromContext extracts the user from the context.
func FromContext(c context.Context) (*Account, error) {
	account, ok := c.Value(keyHandler{}).(*Account)
	if !ok {
		return nil, errors.New("no user found in context")
	}

	return account, nil
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
