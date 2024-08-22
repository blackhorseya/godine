package biz

import (
	"github.com/blackhorseya/godine/app/domain/logistics/repo/delivery"
	"github.com/google/wire"
)

// ProviderLogisticsBizSet is biz provider set.
var ProviderLogisticsBizSet = wire.NewSet(NewLogisticsService, delivery.NewMongodb)
