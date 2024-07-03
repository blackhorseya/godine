// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package restful

import (
	"github.com/blackhorseya/godine/adapter/restaurant/wirex"
	"github.com/blackhorseya/godine/app/domain/restaurant/biz"
	"github.com/blackhorseya/godine/app/domain/restaurant/repo/restaurant"
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/app/infra/storage/mongodbx"
	"github.com/blackhorseya/godine/app/infra/storage/redix"
	"github.com/blackhorseya/godine/app/infra/transports/httpx"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/logging"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

import (
	_ "github.com/blackhorseya/godine/api/restaurant/restful"
)

// Injectors from wire.go:

func New(v *viper.Viper) (adapterx.Restful, error) {
	application, err := initApplication()
	if err != nil {
		return nil, err
	}
	client, err := mongodbx.NewClient(application)
	if err != nil {
		return nil, err
	}
	redisClient, err := redix.NewRedis(application)
	if err != nil {
		return nil, err
	}
	iRestaurantRepo := restaurant.NewMongodb(client, redisClient)
	iRestaurantBiz := biz.NewRestaurantBiz(iRestaurantRepo)
	iMenuBiz := biz.NewMenuBiz(iRestaurantRepo)
	injector := &wirex.Injector{
		A:                 application,
		RestaurantService: iRestaurantBiz,
		MenuService:       iMenuBiz,
	}
	server, err := httpx.NewServer(application)
	if err != nil {
		return nil, err
	}
	restful := newRestful(injector, server)
	return restful, nil
}

// wire.go:

func initApplication() (*configx.Application, error) {
	app, err := configx.LoadApplication(&configx.C.RestaurantRestful)
	if err != nil {
		return nil, err
	}

	err = logging.Init(app.Log)
	if err != nil {
		return nil, err
	}

	err = otelx.SetupOTelSDK(contextx.Background(), app)
	if err != nil {
		return nil, err
	}

	return app, nil
}

var providerSet = wire.NewSet(
	newRestful, wire.Struct(new(wirex.Injector), "*"), initApplication, httpx.NewServer, biz.NewRestaurantBiz, biz.NewMenuBiz, restaurant.NewMongodb, mongodbx.NewClient, redix.NewRedis,
)
