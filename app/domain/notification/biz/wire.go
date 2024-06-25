package biz

import (
	"github.com/google/wire"
)

// ProviderNotificationSet is a provider set for creating a notification service.
var ProviderNotificationSet = wire.NewSet(NewNotification)
