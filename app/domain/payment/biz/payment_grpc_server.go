package biz

import (
	"context"
	"fmt"
	"strconv"

	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/payment/biz"
	"github.com/blackhorseya/godine/entity/domain/payment/model"
	"github.com/blackhorseya/godine/entity/domain/payment/repo"
	userM "github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"go.uber.org/zap"
)

type paymentService struct {
	payments repo.IPaymentRepo
}

// NewPaymentService creates a new payment service.
func NewPaymentService(payments repo.IPaymentRepo) biz.PaymentServiceServer {
	return &paymentService{
		payments: payments,
	}
}

func (i *paymentService) CreatePayment(c context.Context, req *biz.CreatePaymentRequest) (*model.Payment, error) {
	ctx, err := contextx.FromContextLegacy(c)
	if err != nil {
		return nil, fmt.Errorf("failed to get contextx: %w", err)
	}

	ctx, span := otelx.Span(ctx, "payment.biz.CreatePayment")
	defer span.End()

	handler, err := userM.FromContext(ctx)
	if err != nil {
		ctx.Error("failed to get user from context", zap.Error(err))
		return nil, err
	}

	payment, err := model.NewPayment(handler.Id, strconv.FormatInt(req.OrderId, 10), req.Amount)
	if err != nil {
		ctx.Error("failed to create payment", zap.Error(err))
		return nil, err
	}

	err = i.payments.Create(ctx, payment)
	if err != nil {
		ctx.Error("failed to create payment", zap.Error(err))
		return nil, err
	}

	return payment, nil
}

func (i *paymentService) GetPayment(c context.Context, req *biz.GetPaymentRequest) (*model.Payment, error) {
	ctx, err := contextx.FromContextLegacy(c)
	if err != nil {
		return nil, fmt.Errorf("failed to get contextx: %w", err)
	}

	ctx, span := otelx.Span(ctx, "payment.biz.GetPayment")
	defer span.End()

	payment, err := i.payments.GetByID(ctx, req.PaymentId)
	if err != nil {
		ctx.Error("failed to get payment", zap.Error(err))
		return nil, err
	}

	return payment, nil
}

func (i *paymentService) ListPayments(
	c context.Context,
	req *biz.ListPaymentsRequest,
) (*biz.ListPaymentsResponse, error) {
	// TODO: 2024/8/31|sean|implement me
	panic("implement me")
}
