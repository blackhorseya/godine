package biz

import (
	"github.com/google/wire"
)

// ProviderPaymentBizSet for wire
var ProviderPaymentBizSet = wire.NewSet(
	NewPaymentService,
)
