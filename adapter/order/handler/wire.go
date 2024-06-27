//go:build wireinject

//go:generate wire

package handler

import (
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/app/infra/transports/httpx"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/logging"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func initApplication() (*configx.Application, error) {
	app, err := configx.LoadApplication(&configx.C.OrderHandler)
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
	newRestful,

	wire.Struct(new(Injector), "*"),
	initApplication,
	httpx.NewServer,
)

func New(v *viper.Viper) (adapterx.Restful, error) {
	panic(wire.Build(providerSet))
}
