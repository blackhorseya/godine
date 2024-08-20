// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package grpc

import (
	"fmt"
	"github.com/blackhorseya/godine/adapter/platform/wirex"
	"github.com/blackhorseya/godine/app/domain/user/biz"
	"github.com/blackhorseya/godine/app/infra/authx"
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
	"github.com/blackhorseya/godine/app/infra/transports/httpx"
	biz2 "github.com/blackhorseya/godine/entity/domain/user/biz"
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
	authx, err := initAuthx(application)
	if err != nil {
		return nil, err
	}
	injector := &wirex.Injector{
		C:     configuration,
		A:     application,
		Authx: authx,
	}
	accountServiceServer := biz.NewAccountService()
	initServers := NewInitServersFn(accountServiceServer)
	server, err := grpcx.NewServer(application, initServers)
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
	accountServer biz2.AccountServiceServer,
) grpcx.InitServers {
	return func(s *grpc.Server) {
		healthServer := health.NewServer()
		grpc_health_v1.RegisterHealthServer(s, healthServer)
		healthServer.SetServingStatus(serverName, grpc_health_v1.HealthCheckResponse_SERVING)
		biz2.RegisterAccountServiceServer(s, accountServer)
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

func initAuthx(app *configx.Application) (*authx.Authx, error) {
	return authx.New(app.Auth0)
}
