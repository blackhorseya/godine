package utils

import (
	"context"
)

// ListCondition is a struct that defines the condition for listing items.
type ListCondition struct {
	Limit  int64
	Offset int64
}

// IRepository is a generic interface for repositories.
type IRepository[T any] interface {
	Create(c context.Context, item *T) error
	GetByID(c context.Context, id string) (item *T, err error)
	List(c context.Context, cond ListCondition) (items []*T, total int, err error)
	Update(c context.Context, item *T) error
	Delete(c context.Context, id string) error
}
