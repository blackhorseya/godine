package domain

import (
	"github.com/google/wire"
)

// ProviderAccountServiceSet is used to provide account service set.
var ProviderAccountServiceSet = wire.NewSet(NewAccountService, NewAccountServiceClient)

// ProviderRestaurantServiceSet is used to provide restaurant service set.
var ProviderRestaurantServiceSet = wire.NewSet(
	NewRestaurantService,
	NewRestaurantServiceClient,
	NewMenuService,
	NewMenuServiceClient,
)

// ProviderOrderServiceSet is used to provide order service set.
var ProviderOrderServiceSet = wire.NewSet(NewOrderService, NewOrderServiceClient)

// ProviderPaymentServiceSet is used to provide payment service set.
var ProviderPaymentServiceSet = wire.NewSet(NewPaymentService, NewPaymentServiceClient)

// ProviderLogisticsServiceSet is used to provide logistics service set.
var ProviderLogisticsServiceSet = wire.NewSet(NewLogisticsService, NewLogisticsServiceClient)

// ProviderNotificationServiceSet is used to provide notification service set.
var ProviderNotificationServiceSet = wire.NewSet(NewNotificationService, NewNotificationServiceClient)
