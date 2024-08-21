package model

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewPaymentLegacy creates a new Payment.
// Deprecated: use NewPayment instead.
func NewPaymentLegacy(orderID string, amount *PaymentAmount) *Payment {
	return &Payment{
		OrderId:   orderID,
		Amount:    amount,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}
}

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
