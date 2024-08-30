//go:build wireinject

//go:generate wire

package user

import (
	"fmt"

	"github.com/blackhorseya/godine/app/infra/authx"
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/app/infra/transports/httpx"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func initApplication(config *configx.Configuration) (*configx.Application, error) {
	app, err := config.GetService("user")
	if err != nil {
		return nil, fmt.Errorf("failed to get service %s: %w", "platform", err)
	}

	err = otelx.SetupOTelSDK(contextx.Background(), app)
	if err != nil {
		return nil, fmt.Errorf("failed to setup otel sdk: %w", err)
	}

	return app, nil
}

func New(v *viper.Viper) (adapterx.Restful, error) {
	panic(wire.Build(
		NewRestful,
		wire.Struct(new(Injector), "*"),
		httpx.NewServer,
		initApplication,
		configx.NewConfiguration,
		authx.New,
	))
}
