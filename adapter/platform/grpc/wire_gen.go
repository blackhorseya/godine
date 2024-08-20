// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package grpc

import (
	"fmt"
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
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
	initServers := NewInitServersFn()
	server, err := grpcx.NewServer(application, initServers)
	if err != nil {
		return nil, err
	}
	restful := NewServer(server)
	return restful, nil
}

// wire.go:

const serverName = "platform"

// NewInitServersFn creates and returns a new InitServers function.
func NewInitServersFn() grpcx.InitServers {
	return func(s *grpc.Server) {

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
