package biz

import (
	"github.com/blackhorseya/godine/app/domain/logistics/repo/delivery"
	"github.com/google/wire"
)

// ProviderLogisticsSet is biz provider set.
var ProviderLogisticsSet = wire.NewSet(delivery.NewMongodb)
