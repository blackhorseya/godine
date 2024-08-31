package user

import (
	"github.com/google/wire"
)

// ProviderUserBizSet is a provider set for user biz.
var ProviderUserBizSet = wire.NewSet(NewAccountService)
