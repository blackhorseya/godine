//go:build wireinject

//go:generate wire

package restaurant

import (
	"github.com/blackhorseya/godine/app/infra/storage/mongodbx"
	"github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProviderRestaurantBizSet is a provider set for restaurant biz.
var ProviderRestaurantBizSet = wire.NewSet(
	NewRestaurantService,
	NewMenuService,
	mongodbx.NewRestaurantRepo,
)

func NewIntegration(rw *mongo.Client, rdb *redis.Client) (biz.RestaurantServiceServer, error) {
	panic(wire.Build(ProviderRestaurantBizSet))
}
