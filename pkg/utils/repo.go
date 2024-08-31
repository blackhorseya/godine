package utils

import (
	"context"
	"time"
)

// BaseModel is a struct that defines the base model.
type BaseModel struct {
	ID        string    `bson:"_id" json:"id"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

// ListCondition is a struct that defines the condition for listing items.
type ListCondition struct {
	Limit  int64
	Offset int64
}

// IRepository is a generic interface for repositories.
type IRepository[T BaseModel] interface {
	Create(c context.Context, item *T) error
	GetByID(c context.Context, id string) (item *T, err error)
	List(c context.Context, cond ListCondition) (items []*T, total int, err error)
	Update(c context.Context, item *T) error
	Delete(c context.Context, id string) error
}
