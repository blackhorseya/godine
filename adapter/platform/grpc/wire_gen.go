// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package grpc

import (
	"fmt"
	"github.com/blackhorseya/godine/adapter/platform/wirex"
	biz2 "github.com/blackhorseya/godine/app/domain/restaurant/biz"
	"github.com/blackhorseya/godine/app/domain/restaurant/repo/restaurant"
	"github.com/blackhorseya/godine/app/domain/user/biz"
	"github.com/blackhorseya/godine/app/infra/authx"
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/app/infra/storage/mongodbx"
	"github.com/blackhorseya/godine/app/infra/storage/redix"
	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
	"github.com/blackhorseya/godine/app/infra/transports/httpx"
	biz4 "github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	biz3 "github.com/blackhorseya/godine/entity/domain/user/biz"
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
	injector := &wirex.Injector{
		C:     configuration,
		A:     application,
		Authx: authxAuthx,
	}
	accountServiceServer := biz.NewAccountService()
	client, err := mongodbx.NewClient(application)
	if err != nil {
		return nil, err
	}
	redisClient, err := redix.NewClient(application)
	if err != nil {
		return nil, err
	}
	iRestaurantRepo := restaurant.NewMongodb(client, redisClient)
	restaurantServiceServer := biz2.NewRestaurantService(iRestaurantRepo)
	menuServiceServer := biz2.NewMenuService(iRestaurantRepo)
	initServers := NewInitServersFn(accountServiceServer, restaurantServiceServer, menuServiceServer)
	server, err := grpcx.NewServer(application, initServers, authxAuthx)
	if err != nil {
		return nil, err
	}
	httpxServer, err := httpx.NewServer(application)
	if err != nil {
		return nil, err
	}
	restful := NewServer(injector, server, httpxServer)
	return restful, nil
}

// wire.go:

const serverName = "platform"

// NewInitServersFn creates and returns a new InitServers function.
func NewInitServersFn(
	accountServer biz3.AccountServiceServer,
	restaurantServer biz4.RestaurantServiceServer,
	menuServer biz4.MenuServiceServer,
) grpcx.InitServers {
	return func(s *grpc.Server) {
		healthServer := health.NewServer()
		grpc_health_v1.RegisterHealthServer(s, healthServer)
		healthServer.SetServingStatus(serverName, grpc_health_v1.HealthCheckResponse_SERVING)
		biz3.RegisterAccountServiceServer(s, accountServer)
		biz4.RegisterRestaurantServiceServer(s, restaurantServer)
		biz4.RegisterMenuServiceServer(s, menuServer)
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
