package model

import (
	"errors"

	"github.com/blackhorseya/godine/pkg/contextx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewDeliveryLegacy creates a new delivery entity.
// Deprecated: use NewDelivery instead.
func NewDeliveryLegacy(orderID string, userID string) *Delivery {
	return &Delivery{
		Id:         "",
		OrderId:    orderID,
		UserId:     userID,
		DriverId:   userID,
		Status:     DeliveryStatus_DELIVERY_STATUS_PENDING,
		PickupAt:   nil,
		DeliveryAt: nil,
		CreatedAt:  timestamppb.Now(),
		UpdatedAt:  timestamppb.Now(),
	}
}

// NewDelivery creates a new delivery entity.
func NewDelivery() (*Delivery, error) {
	// TODO: 2024/8/22|sean|add more fields
	return &Delivery{
		Id:         "",
		OrderId:    "",
		UserId:     "",
		DriverId:   "",
		Status:     0,
		PickupAt:   nil,
		DeliveryAt: nil,
		CreatedAt:  nil,
		UpdatedAt:  nil,
	}, nil
}

func (x *Delivery) UnmarshalBSON(bytes []byte) error {
	type Alias Delivery
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

func (x *Delivery) MarshalBSON() ([]byte, error) {
	type Alias Delivery
	alias := &struct {
		*Alias `bson:",inline"`
		ID     primitive.ObjectID `bson:"_id"`
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

// Next returns the next delivery event.
func (x *Delivery) Next(ctx contextx.Contextx) (event *DeliveryEvent, err error) {
	// TODO: 2024/8/20|sean|implement the delivery state transition
	return nil, errors.New("not implemented")
}
