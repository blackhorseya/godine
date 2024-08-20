package model

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewPayment creates a new Payment.
func NewPayment(orderID string, amount *PaymentAmount) *Payment {
	return &Payment{
		OrderId:   orderID,
		Amount:    amount,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}
}

// Equal checks if the payment is equal to another payment.
func (x *Payment) Equal(v *Payment) bool {
	return x.Id == v.Id
}
