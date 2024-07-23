package model

import (
	"time"
)

// Payment represents a payment entity.
type Payment struct {
	// ID is the unique identifier of the payment.
	ID string `json:"id,omitempty" bson:"_id,omitempty"`

	// OrderID is the identifier of the associated order.
	OrderID string `json:"order_id,omitempty" bson:"orderId"`

	// Amount is the value object representing the payment amount.
	Amount PaymentAmount `json:"amount,omitempty" bson:"amount"`

	// Status is the value object representing the payment status.
	Status PaymentStatus `json:"status,omitempty" bson:"status"`

	// Records is the list of payment records.
	Records []*PaymentRecord `json:"records,omitempty" bson:"records"`

	// CreatedAt is the time when the payment was created.
	CreatedAt time.Time `json:"created_at,omitempty" bson:"createdAt"`

	// UpdatedAt is the time when the payment was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updatedAt"`
}

// NewPayment creates a new Payment.
func NewPayment(orderID string, amount PaymentAmount) *Payment {
	return &Payment{
		OrderID:   orderID,
		Amount:    amount,
		Status:    PaymentStatusPending,
		Records:   []*PaymentRecord{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// AddPaymentRecord adds a new record to the payment.
func (p *Payment) AddPaymentRecord(status PaymentStatus, errorMessage string) {
	record := &PaymentRecord{
		Status:       status,
		ErrorMessage: errorMessage,
		Timestamp:    time.Now(),
	}
	p.Records = append(p.Records, record)
	p.Status = status
	p.UpdatedAt = time.Now()
}

// Equal checks if the payment is equal to another payment.
func (p *Payment) Equal(v *Payment) bool {
	return p.ID == v.ID
}
