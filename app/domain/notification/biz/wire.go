package biz

import (
	notification2 "github.com/blackhorseya/godine/app/domain/notification/repo/notification"
	"github.com/google/wire"
)

// ProviderNotificationSet is a provider set for creating a notification service.
var ProviderNotificationSet = wire.NewSet(NewNotification, notification2.NewMongodb)
