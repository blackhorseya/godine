package order

import (
	"github.com/blackhorseya/godine/app/infra/storage/postgresqlx"
	"github.com/google/wire"
)

// ProviderOrderBizSet is biz provider set.
var ProviderOrderBizSet = wire.NewSet(
	NewOrderService,
	postgresqlx.NewOrderRepo,
)
