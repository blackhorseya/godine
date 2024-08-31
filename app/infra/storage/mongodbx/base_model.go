package mongodbx

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BaseModelInterface ensures that the type has an embedded BaseModel or equivalent fields.
type BaseModelInterface interface {
	GetID() primitive.ObjectID
	SetID(id primitive.ObjectID)
	GetCreatedAt() time.Time
	SetCreatedAt(t time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(t time.Time)
}

// BaseModel is a struct that defines the base model.
type BaseModel struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

func (x *BaseModel) GetID() primitive.ObjectID {
	return x.ID
}

func (x *BaseModel) SetID(id primitive.ObjectID) {
	x.ID = id
}

func (x *BaseModel) GetCreatedAt() time.Time {
	return x.CreatedAt
}

func (x *BaseModel) SetCreatedAt(t time.Time) {
	x.CreatedAt = t
}

func (x *BaseModel) GetUpdatedAt() time.Time {
	return x.UpdatedAt
}

func (x *BaseModel) SetUpdatedAt(t time.Time) {
	x.UpdatedAt = t
}
