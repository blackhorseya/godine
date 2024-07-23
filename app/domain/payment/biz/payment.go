package biz

import (
	"github.com/blackhorseya/godine/entity/domain/payment/biz"
	"github.com/blackhorseya/godine/entity/domain/payment/model"
	"github.com/blackhorseya/godine/pkg/contextx"
)

type impl struct {
}

// NewPaymentBiz creates a new payment service.
func NewPaymentBiz() biz.IPaymentBiz {
	return &impl{}
}

func (i *impl) GetPaymentByID(ctx contextx.Contextx, id string) (item *model.Payment, err error) {
	// todo: 2024/7/23|sean|implement me
	panic("implement me")
}

func (i *impl) CreatePayment(
	ctx contextx.Contextx,
	orderID string,
	amount model.PaymentAmount,
) (item *model.Payment, err error) {
	// todo: 2024/7/23|sean|implement me
	panic("implement me")
}

func (i *impl) ListPayments(
	ctx contextx.Contextx,
	options biz.ListPaymentsOptions,
) (items []*model.Payment, total int, err error) {
	// todo: 2024/7/23|sean|implement me
	panic("implement me")
}
