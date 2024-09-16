//go:build wireinject

//go:generate wire

package platform

import (
	"fmt"

	"github.com/blackhorseya/godine/adapter/platform/handlers"
	"github.com/blackhorseya/godine/app/domain/logistics"
	"github.com/blackhorseya/godine/app/domain/notification"
	"github.com/blackhorseya/godine/app/domain/order"
	"github.com/blackhorseya/godine/app/domain/payment"
	"github.com/blackhorseya/godine/app/domain/restaurant"
	"github.com/blackhorseya/godine/app/domain/user"
	"github.com/blackhorseya/godine/app/infra/authx"
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/app/infra/snowflakex"
	"github.com/blackhorseya/godine/app/infra/storage/mongodbx"
	"github.com/blackhorseya/godine/app/infra/storage/postgresqlx"
	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
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
		wire.Struct(new(Injector), "*"),
		grpcx.NewServer,
		InitApplication,
		configx.NewConfiguration,
		NewInitServersFn,
		authx.New,
		grpcx.NewClient,
		otelx.NewSDK,
		user.ProviderUserBizSet,
		user.NewAccountServiceClient,
		restaurant.ProviderRestaurantBizSet,
		restaurant.NewRestaurantServiceClient,
		restaurant.NewMenuServiceClient,
		notification.ProviderNotificationBizSet,
		notification.NewNotificationServiceClient,
		payment.ProviderPaymentBizSet,
		payment.NewPaymentServiceClient,
		order.ProviderOrderBizSet,
		logistics.ProviderLogisticsBizSet,
		logistics.NewLogisticsServiceClient,
		snowflakex.NewNode,
		postgresqlx.NewClient,
		mongodbx.NewClientWithClean,
		handlers.NewRestaurantServiceHandler,
	))
}
