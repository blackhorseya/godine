package biz

import (
	"context"

	"github.com/blackhorseya/godine/entity/domain/payment/biz"
	"github.com/blackhorseya/godine/entity/domain/payment/model"
)

type paymentService struct {
}

// NewPaymentService creates a new payment service.
func NewPaymentService() biz.PaymentServiceServer {
	return &paymentService{}
}

func (i *paymentService) CreatePayment(c context.Context, req *biz.CreatePaymentRequest) (*model.Payment, error) {
	// TODO: 2024/8/21|sean|implement me
	panic("implement me")
}

func (i *paymentService) GetPayment(c context.Context, req *biz.GetPaymentRequest) (*model.Payment, error) {
	// TODO: 2024/8/21|sean|implement me
	panic("implement me")
}
