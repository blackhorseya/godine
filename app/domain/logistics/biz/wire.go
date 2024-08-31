package biz

import (
	"github.com/blackhorseya/godine/app/infra/storage/mongodbx"
	"github.com/google/wire"
)

// ProviderLogisticsBizSet is biz provider set.
var ProviderLogisticsBizSet = wire.NewSet(NewLogisticsService, mongodbx.NewDeliveryRepo)
