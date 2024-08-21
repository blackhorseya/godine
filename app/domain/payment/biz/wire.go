package biz

import (
	"github.com/blackhorseya/godine/app/domain/payment/repo/payment"
	"github.com/google/wire"
)

// ProviderPaymentBizSet for wire
var ProviderPaymentBizSet = wire.NewSet(
	NewPaymentService,
	payment.NewMongodb,
)
