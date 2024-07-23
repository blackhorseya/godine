package model

import (
	"time"
)

// PaymentAmount represents a value object for payment amount.
type PaymentAmount struct {
	// Amount is the monetary value.
	Amount float64 `json:"amount,omitempty" bson:"amount"`

	// Currency is the currency of the amount.
	Currency string `json:"currency,omitempty" bson:"currency"`
}

// PaymentStatus represents a value object for payment status.
type PaymentStatus struct {
	// Status is the status of the payment.
	Status string `json:"status,omitempty" bson:"status"`
}

var (
	// PaymentStatusPending represents a pending payment status.
	PaymentStatusPending = PaymentStatus{Status: "Pending"}

	// PaymentStatusCompleted represents a completed payment status.
	PaymentStatusCompleted = PaymentStatus{Status: "Completed"}

	// PaymentStatusFailed represents a failed payment status.
	PaymentStatusFailed = PaymentStatus{Status: "Failed"}
)

// PaymentRecord represents a record of a payment transaction.
type PaymentRecord struct {
	// ID is the unique identifier of the payment record.
	ID string `json:"id,omitempty" bson:"_id,omitempty"`

	// Status is the status of the payment record.
	Status PaymentStatus `json:"status,omitempty" bson:"status"`

	// ErrorMessage is the error message in case of a failed transaction.
	ErrorMessage string `json:"error_message,omitempty" bson:"error_message"`

	// Timestamp is the time when the transaction occurred.
	Timestamp time.Time `json:"timestamp,omitempty" bson:"timestamp"`
}
