package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewNotification creates a new notification entity.
func NewNotification(from, to string, orderID string, message string) *Notification {
	return &Notification{
		Id:        "",
		SenderId:  from,
		UserId:    to,
		OrderId:   orderID,
		Type:      "order_status",
		Message:   message,
		Status:    "pending",
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}
}

func (x *Notification) UnmarshalBSON(bytes []byte) error {
	type Alias Notification
	alias := &struct {
		ID     primitive.ObjectID `bson:"_id"`
		*Alias `bson:",inline"`
	}{
		Alias: (*Alias)(x),
	}

	if err := bson.Unmarshal(bytes, alias); err != nil {
		return err
	}

	x.Id = alias.ID.Hex()

	return nil
}

func (x *Notification) MarshalBSON() ([]byte, error) {
	type Alias Notification
	alias := &struct {
		ID     primitive.ObjectID `bson:"_id"`
		*Alias `bson:",inline"`
	}{
		Alias: (*Alias)(x),
	}

	id, err := primitive.ObjectIDFromHex(x.Id)
	if err != nil {
		return nil, err
	}
	alias.ID = id

	return bson.Marshal(alias)
}
