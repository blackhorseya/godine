//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/godine/entity/domain/payment/model"
	"github.com/blackhorseya/godine/pkg/contextx"
)

// ListPaymentsOptions defines the options for listing payments.
type ListPaymentsOptions struct {
	// Page is the page number.
	Page int `form:"page" default:"1" minimum:"1"`

	// Size is the number of items per page.
	Size int `form:"size" default:"10" minimum:"1" maximum:"100"`
}

// IPaymentBiz is an interface for payment service
type IPaymentBiz interface {
	// GetPaymentByID gets payment by id
	GetPaymentByID(ctx contextx.Contextx, id string) (item *model.Payment, err error)

	// CreatePayment creates a new payment
	CreatePayment(ctx contextx.Contextx, orderID string, amount model.PaymentAmount) (item *model.Payment, err error)

	// ListPayments lists payments
	ListPayments(ctx contextx.Contextx, options ListPaymentsOptions) (items []*model.Payment, total int, err error)
}
