//go:build wireinject

//go:generate wire

package restful

import (
	"github.com/blackhorseya/godine/adapter/logistics/wirex"
	"github.com/blackhorseya/godine/app/domain/logistics/biz"
	biz2 "github.com/blackhorseya/godine/app/domain/notification/biz"
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/app/infra/storage/mongodbx"
	"github.com/blackhorseya/godine/app/infra/transports/httpx"
	"github.com/blackhorseya/godine/app/infra/transports/kafkax"
	"github.com/blackhorseya/godine/app/infra/transports/mqx"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/logging"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func initApplication() (*configx.Application, error) {
	app, err := configx.LoadApplication(&configx.C.LogisticsRestful)
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

	wire.Struct(new(wirex.Injector), "*"),
	initApplication,
	httpx.NewServer,

	biz.ProviderLogisticsSet,
	mongodbx.NewClient,
	kafkax.NewWriter,
	mqx.NewKafkaEventBus,

	biz2.NewNotificationHTTPClient,
)

func New(v *viper.Viper) (adapterx.Restful, error) {
	panic(wire.Build(providerSet))
}
