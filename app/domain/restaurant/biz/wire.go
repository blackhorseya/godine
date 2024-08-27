//go:build wireinject

//go:generate wire

package biz

import (
	"github.com/blackhorseya/godine/app/domain/restaurant/repo/restaurant"
	"github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProviderRestaurantBizSet is a provider set for restaurant biz.
var ProviderRestaurantBizSet = wire.NewSet(
	NewRestaurantService,
	NewMenuService,
	restaurant.NewMongodb,
)

func NewIntegration(rw *mongo.Client, rdb *redis.Client) (biz.RestaurantServiceServer, error) {
	panic(wire.Build(ProviderRestaurantBizSet))
}
