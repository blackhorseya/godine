package biz

import (
	"github.com/blackhorseya/godine/app/domain/order/repo/order"
	"github.com/google/wire"
)

// ProviderOrderBizSet is biz provider set.
var ProviderOrderBizSet = wire.NewSet(
	NewOrderService,
	order.NewGORM,
)
