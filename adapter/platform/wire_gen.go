// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package platform

import (
	"fmt"
	"github.com/blackhorseya/godine/adapter/platform/handlers"
	biz2 "github.com/blackhorseya/godine/app/domain/logistics/biz"
	"github.com/blackhorseya/godine/app/domain/logistics/repo/delivery"
	"github.com/blackhorseya/godine/app/domain/notification"
	biz3 "github.com/blackhorseya/godine/app/domain/order/biz"
	"github.com/blackhorseya/godine/app/domain/order/repo/order"
	"github.com/blackhorseya/godine/app/domain/payment"
	"github.com/blackhorseya/godine/app/domain/restaurant"
	"github.com/blackhorseya/godine/app/domain/user/biz"
	"github.com/blackhorseya/godine/app/infra/authx"
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/app/infra/snowflakex"
	"github.com/blackhorseya/godine/app/infra/storage/mongodbx"
	"github.com/blackhorseya/godine/app/infra/storage/postgresqlx"
	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
	biz9 "github.com/blackhorseya/godine/entity/domain/logistics/biz"
	biz7 "github.com/blackhorseya/godine/entity/domain/notification/biz"
	biz8 "github.com/blackhorseya/godine/entity/domain/order/biz"
	biz6 "github.com/blackhorseya/godine/entity/domain/payment/biz"
	biz5 "github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	biz4 "github.com/blackhorseya/godine/entity/domain/user/biz"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// Injectors from wire.go:

func New(v *viper.Viper) (adapterx.Restful, error) {
	configuration, err := configx.NewConfiguration(v)
	if err != nil {
		return nil, err
	}
	application, err := initApplication(configuration)
	if err != nil {
		return nil, err
	}
	authxAuthx, err := authx.New(application)
	if err != nil {
		return nil, err
	}
	client, err := grpcx.NewClient(configuration, authxAuthx)
	if err != nil {
		return nil, err
	}
	restaurantServiceClient, err := restaurant.NewRestaurantServiceClient(client)
	if err != nil {
		return nil, err
	}
	restaurantServiceHandler := handlers.NewRestaurantServiceHandler(restaurantServiceClient)
	injector := &Injector{
		C:                        configuration,
		A:                        application,
		Authx:                    authxAuthx,
		RestaurantServiceHandler: restaurantServiceHandler,
	}
	accountServiceServer := biz.NewAccountService()
	mongoClient, err := mongodbx.NewClient(application)
	if err != nil {
		return nil, err
	}
	iRestaurantRepo := mongodbx.NewRestaurantRepo(mongoClient)
	restaurantServiceServer := restaurant.NewRestaurantService(iRestaurantRepo)
	menuServiceServer := restaurant.NewMenuService(iRestaurantRepo)
	iPaymentRepo := mongodbx.NewPaymentRepo(mongoClient)
	paymentServiceServer := payment.NewPaymentService(iPaymentRepo)
	iNotificationRepo := mongodbx.NewNotificationRepo(mongoClient)
	notificationServiceServer := notification.NewNotificationService(iNotificationRepo)
	db, err := postgresqlx.NewClient(application)
	if err != nil {
		return nil, err
	}
	node, err := snowflakex.NewNode()
	if err != nil {
		return nil, err
	}
	iOrderRepo, err := order.NewGORM(db, node)
	if err != nil {
		return nil, err
	}
	menuServiceClient, err := restaurant.NewMenuServiceClient(client)
	if err != nil {
		return nil, err
	}
	accountServiceClient, err := biz.NewAccountServiceClient(client)
	if err != nil {
		return nil, err
	}
	notificationServiceClient, err := notification.NewNotificationServiceClient(client)
	if err != nil {
		return nil, err
	}
	paymentServiceClient, err := payment.NewPaymentServiceClient(client)
	if err != nil {
		return nil, err
	}
	logisticsServiceClient, err := biz2.NewLogisticsServiceClient(client)
	if err != nil {
		return nil, err
	}
	orderServiceServer := biz3.NewOrderService(iOrderRepo, restaurantServiceClient, menuServiceClient, accountServiceClient, notificationServiceClient, paymentServiceClient, logisticsServiceClient)
	iDeliveryRepo := delivery.NewMongodb(mongoClient)
	logisticsServiceServer := biz2.NewLogisticsService(iDeliveryRepo, notificationServiceClient)
	initServers := NewInitServersFn(accountServiceServer, restaurantServiceServer, menuServiceServer, paymentServiceServer, notificationServiceServer, orderServiceServer, logisticsServiceServer)
	server, err := grpcx.NewServer(application, initServers, authxAuthx)
	if err != nil {
		return nil, err
	}
	restful := NewServer(injector, server)
	return restful, nil
}

// wire.go:

const serverName = "platform"

// NewInitServersFn creates and returns a new InitServers function.
func NewInitServersFn(
	accountServer biz4.AccountServiceServer,
	restaurantServer biz5.RestaurantServiceServer,
	menuServer biz5.MenuServiceServer,
	paymentServer biz6.PaymentServiceServer,
	notifyServer biz7.NotificationServiceServer,
	orderServer biz8.OrderServiceServer,
	logisticsServer biz9.LogisticsServiceServer,
) grpcx.InitServers {
	return func(s *grpc.Server) {
		healthServer := health.NewServer()
		grpc_health_v1.RegisterHealthServer(s, healthServer)
		healthServer.SetServingStatus(serverName, grpc_health_v1.HealthCheckResponse_SERVING)
		biz4.RegisterAccountServiceServer(s, accountServer)
		biz5.RegisterRestaurantServiceServer(s, restaurantServer)
		biz5.RegisterMenuServiceServer(s, menuServer)
		biz6.RegisterPaymentServiceServer(s, paymentServer)
		biz7.RegisterNotificationServiceServer(s, notifyServer)
		biz8.RegisterOrderServiceServer(s, orderServer)
		biz9.RegisterLogisticsServiceServer(s, logisticsServer)
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
