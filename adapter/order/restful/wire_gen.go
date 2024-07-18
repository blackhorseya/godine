// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package restful

import (
	"github.com/blackhorseya/godine/adapter/order/wirex"
	biz3 "github.com/blackhorseya/godine/app/domain/logistics/biz"
	biz4 "github.com/blackhorseya/godine/app/domain/notification/biz"
	biz5 "github.com/blackhorseya/godine/app/domain/order/biz"
	"github.com/blackhorseya/godine/app/domain/order/repo/order"
	"github.com/blackhorseya/godine/app/domain/restaurant/biz"
	biz2 "github.com/blackhorseya/godine/app/domain/user/biz"
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/app/infra/snowflakex"
	"github.com/blackhorseya/godine/app/infra/storage/postgresqlx"
	"github.com/blackhorseya/godine/app/infra/transports/httpx"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/logging"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

import (
	_ "github.com/blackhorseya/godine/api/order/restful"
)

// Injectors from wire.go:

func New(v *viper.Viper) (adapterx.Restful, error) {
	configuration, err := configx.NewConfiguration(v)
	if err != nil {
		return nil, err
	}
	application, err := initApplication(v)
	if err != nil {
		return nil, err
	}
	iRestaurantBiz := biz.NewRestaurantHTTPClient(configuration)
	iMenuBiz := biz.NewMenuHTTPClient(configuration)
	iUserBiz := biz2.NewUserHTTPClient(configuration)
	iLogisticsBiz := biz3.NewLogisticsHTTPClient(configuration)
	iNotificationBiz := biz4.NewNotificationHTTPClient(configuration)
	db, err := postgresqlx.NewClient(application)
	if err != nil {
		return nil, err
	}
	node, err := snowflakex.NewNode()
	if err != nil {
		return nil, err
	}
	iOrderRepo, err := order.NewMariadb(db, node)
	if err != nil {
		return nil, err
	}
	iOrderBiz := biz5.NewOrderBiz(iRestaurantBiz, iMenuBiz, iUserBiz, iLogisticsBiz, iNotificationBiz, iOrderRepo)
	injector := &wirex.Injector{
		C:            configuration,
		A:            application,
		OrderService: iOrderBiz,
	}
	server, err := httpx.NewServer(application)
	if err != nil {
		return nil, err
	}
	restful := newRestful(injector, server)
	return restful, nil
}

// wire.go:

func initApplication(v *viper.Viper) (*configx.Application, error) {
	app, err := configx.NewApplication(v, "orderRestful")
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
	newRestful, wire.Struct(new(wirex.Injector), "*"), configx.NewConfiguration, initApplication, httpx.NewServer, biz5.NewOrderBiz, biz.NewRestaurantHTTPClient, biz.NewMenuHTTPClient, biz2.NewUserHTTPClient, order.NewMariadb, postgresqlx.NewClient, snowflakex.NewNode, biz3.NewLogisticsHTTPClient, biz4.NewNotificationHTTPClient,
)
