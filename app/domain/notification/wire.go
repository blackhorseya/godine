package notification

import (
	"github.com/blackhorseya/godine/app/infra/storage/mongodbx"
	"github.com/google/wire"
)

// ProviderNotificationBizSet is a provider set for creating a impl service.
var ProviderNotificationBizSet = wire.NewSet(
	NewNotificationService,
	mongodbx.NewNotificationRepo,
)
