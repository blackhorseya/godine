package biz

import (
	"github.com/blackhorseya/godine/app/domain/restaurant/repo/restaurant"
	"github.com/google/wire"
)

// ProviderRestaurantBizSet is a provider set for restaurant biz.
var ProviderRestaurantBizSet = wire.NewSet(
	NewRestaurantService,
	NewMenuService,
	restaurant.NewMongodb,
)
