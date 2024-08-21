package biz

import (
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/payment/biz"
	"github.com/blackhorseya/godine/entity/domain/payment/model"
	"github.com/blackhorseya/godine/entity/domain/payment/repo"
	"github.com/blackhorseya/godine/pkg/contextx"
)

type impl struct {
	payments repo.IPaymentRepo
}

// NewPaymentBiz creates a new payment service.
func NewPaymentBiz(payments repo.IPaymentRepo) biz.IPaymentBiz {
	return &impl{
		payments: payments,
	}
}

func (i *impl) GetPaymentByID(ctx contextx.Contextx, id string) (item *model.Payment, err error) {
	ctx, span := otelx.Span(ctx, "biz.payment.GetPaymentByID")
	defer span.End()

	return i.payments.GetByID(ctx, id)
}

func (i *impl) CreatePayment(
	ctx contextx.Contextx,
	orderID string,
	amount *model.PaymentAmount,
) (item *model.Payment, err error) {
	ctx, span := otelx.Span(ctx, "biz.payment.CreatePayment")
	defer span.End()

	payment := model.NewPaymentLegacy(orderID, amount)
	err = i.payments.Create(ctx, payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (i *impl) ListPayments(
	ctx contextx.Contextx,
	options biz.ListPaymentsOptions,
) (items []*model.Payment, total int, err error) {
	ctx, span := otelx.Span(ctx, "biz.payment.ListPayments")
	defer span.End()

	return i.payments.List(ctx, repo.ListCondition{
		Offset: options.Size,
		Limit:  (options.Page - 1) * options.Size,
	})
}
