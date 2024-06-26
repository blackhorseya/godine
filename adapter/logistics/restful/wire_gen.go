// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package restful

import (
	"github.com/blackhorseya/godine/adapter/logistics/wirex"
	biz2 "github.com/blackhorseya/godine/app/domain/logistics/biz"
	"github.com/blackhorseya/godine/app/domain/logistics/repo/delivery"
	"github.com/blackhorseya/godine/app/domain/notification/biz"
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/app/infra/storage/mongodbx"
	"github.com/blackhorseya/godine/app/infra/transports/httpx"
	"github.com/blackhorseya/godine/app/infra/transports/kafkax"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/logging"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

import (
	_ "github.com/blackhorseya/godine/api/logistics/restful"
)

// Injectors from wire.go:

func New(v *viper.Viper) (adapterx.Restful, error) {
	application, err := initApplication()
	if err != nil {
		return nil, err
	}
	iNotificationBiz := biz.NewNotificationHTTPClient()
	client, err := mongodbx.NewClient(application)
	if err != nil {
		return nil, err
	}
	iDeliveryRepo := delivery.NewMongodb(client)
	writer, err := kafkax.NewWriter()
	if err != nil {
		return nil, err
	}
	iLogisticsBiz := biz2.NewLogistics(iNotificationBiz, iDeliveryRepo, writer)
	injector := &wirex.Injector{
		A:                application,
		LogisticsService: iLogisticsBiz,
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
	app, err := configx.LoadApplication(&configx.C.LogisticsRestful)
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
	newRestful, wire.Struct(new(wirex.Injector), "*"), initApplication, httpx.NewServer, biz2.ProviderLogisticsSet, mongodbx.NewClient, kafkax.NewWriter, biz.NewNotificationHTTPClient,
)
