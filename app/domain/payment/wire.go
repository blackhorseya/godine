package payment

import (
	"github.com/blackhorseya/godine/app/infra/storage/mongodbx"
	"github.com/google/wire"
)

// ProviderPaymentBizSet for wire
var ProviderPaymentBizSet = wire.NewSet(
	NewPaymentService,
	mongodbx.NewPaymentRepo,
)
