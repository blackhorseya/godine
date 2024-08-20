//go:build wireinject

//go:generate wire

package grpc

import (
	"fmt"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const serverName = "platform"

// NewInitServersFn creates and returns a new InitServers function.
func NewInitServersFn() grpcx.InitServers {
	return func(s *grpc.Server) {
		// TODO: 2024/8/21|sean|init servers
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
		grpcx.NewServer,
		initApplication,
		configx.NewConfiguration,
		NewInitServersFn,
	))
}
