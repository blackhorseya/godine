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
	opsB "github.com/blackhorseya/godine/entity/domain/logistics/biz"
	notifyB "github.com/blackhorseya/godine/entity/domain/notification/biz"
	orderB "github.com/blackhorseya/godine/entity/domain/order/biz"
	payB "github.com/blackhorseya/godine/entity/domain/payment/biz"
	restB "github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	userB "github.com/blackhorseya/godine/entity/domain/user/biz"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

const serverName = "platform"

// NewInitServersFn creates and returns a new InitServers function.
func NewInitServersFn(
	accountServer userB.AccountServiceServer,
	restaurantServer restB.RestaurantServiceServer,
	menuServer restB.MenuServiceServer,
	paymentServer payB.PaymentServiceServer,
	notifyServer notifyB.NotificationServiceServer,
	orderServer orderB.OrderServiceServer,
	logisticsServer opsB.LogisticsServiceServer,
) grpcx.InitServers {
	return func(s *grpc.Server) {
		healthServer := health.NewServer()
		grpc_health_v1.RegisterHealthServer(s, healthServer)
		healthServer.SetServingStatus(serverName, grpc_health_v1.HealthCheckResponse_SERVING)

		userB.RegisterAccountServiceServer(s, accountServer)
		restB.RegisterRestaurantServiceServer(s, restaurantServer)
		restB.RegisterMenuServiceServer(s, menuServer)
		payB.RegisterPaymentServiceServer(s, paymentServer)
		notifyB.RegisterNotificationServiceServer(s, notifyServer)
		orderB.RegisterOrderServiceServer(s, orderServer)
		opsB.RegisterLogisticsServiceServer(s, logisticsServer)

		reflection.Register(s)
	}
}

func initApplication(config *configx.Configuration) (*configx.Application, error) {
	app, err := config.GetService(serverName)
	if err != nil {
		return nil, fmt.Errorf("failed to get service %s: %w", serverName, err)
	}

	err = otelx.SetupOTelSDK(contextx.Background(), app)
	if err != nil {
		return nil, fmt.Errorf("failed to setup otel sdk: %w", err)
	}

	return app, nil
}

func New(v *viper.Viper) (adapterx.Restful, error) {
	panic(wire.Build(
		NewServer,
		wire.Struct(new(Injector), "*"),
		grpcx.NewServer,
		initApplication,
		configx.NewConfiguration,
		NewInitServersFn,
		authx.New,
		grpcx.NewClient,

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
		mongodbx.NewClient,

		handlers.NewRestaurantServiceHandler,
	))
}

func NewV2(v *viper.Viper) (adapterx.Restful, func(), error) {
	panic(wire.Build(
		NewServer,
		wire.Struct(new(Injector), "*"),
		grpcx.NewServer,
		initApplication,
		configx.NewConfiguration,
		NewInitServersFn,
		authx.New,
		grpcx.NewClient,

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
