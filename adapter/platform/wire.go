//go:build wireinject

//go:generate wire

package platform

import (
	"fmt"

	"github.com/blackhorseya/godine/adapter/platform/wirex"
	"github.com/blackhorseya/godine/app/domain"
	"github.com/blackhorseya/godine/app/infra/authx"
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/app/infra/snowflakex"
	"github.com/blackhorseya/godine/app/infra/storage/mongodbx"
	"github.com/blackhorseya/godine/app/infra/storage/postgresqlx"
	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
	"github.com/blackhorseya/godine/app/infra/transports/httpx"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

const serviceName = "platform"

// InitApplication is used to initialize the application.
func InitApplication(config *configx.Configuration) (*configx.Application, error) {
	app, err := config.GetService(serviceName)
	if err != nil {
		return nil, fmt.Errorf("failed to get service %s: %w", serviceName, err)
	}

	return app, nil
}

func New(v *viper.Viper) (adapterx.Server, func(), error) {
	panic(wire.Build(
		NewServer,
		wire.Struct(new(wirex.Injector), "*"),
		InitApplication,
		configx.NewConfiguration,

		// adapter layer
		grpcx.NewServer,
		NewInitServersFn,
		httpx.NewServer,

		// biz layer
		domain.ProviderAccountServiceSet,
		domain.ProviderOrderServiceSet,
		domain.ProviderNotificationServiceSet,
		domain.ProviderPaymentServiceSet,
		domain.ProviderLogisticsServiceSet,
		domain.ProviderRestaurantServiceSet,

		// repo layer
		postgresqlx.NewOrderRepo,
		mongodbx.NewRestaurantRepo,
		mongodbx.NewPaymentRepo,
		mongodbx.NewNotificationRepo,
		mongodbx.NewDeliveryRepo,

		// infra layer
		grpcx.NewClient,
		authx.New,
		otelx.NewSDK,
		snowflakex.NewNode,
		postgresqlx.NewClient,
		mongodbx.NewClientWithClean,
	))
}
