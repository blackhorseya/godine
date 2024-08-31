package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewPayment creates a new Payment.
func NewPayment(userID, orderID string, amount *PaymentAmount) (*Payment, error) {
	return &Payment{
		Id:        "",
		Amount:    amount,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
		OrderId:   orderID,
		UserId:    userID,
	}, nil
}

// Equal checks if the payment is equal to another payment.
func (x *Payment) Equal(v *Payment) bool {
	return x.Id == v.Id
}

func (x *Payment) GetID() string {
	return x.Id
}

func (x *Payment) SetID(id primitive.ObjectID) {
	x.Id = id.Hex()
}

func (x *Payment) SetCreatedAt(t *timestamppb.Timestamp) {
	x.CreatedAt = t
}

func (x *Payment) SetUpdatedAt(t *timestamppb.Timestamp) {
	x.UpdatedAt = t
}
