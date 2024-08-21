package biz

import (
	"github.com/blackhorseya/godine/app/domain/notification/repo/notification"
	"github.com/google/wire"
)

// ProviderNotificationBizSet is a provider set for creating a impl service.
var ProviderNotificationBizSet = wire.NewSet(
	NewNotificationService,
	notification.NewMongodb,
)
