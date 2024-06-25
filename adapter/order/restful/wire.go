//go:build wireinject

//go:generate wire

package restful

import (
	"github.com/blackhorseya/godine/adapter/order/wirex"
	biz4 "github.com/blackhorseya/godine/app/domain/logistics/biz"
	"github.com/blackhorseya/godine/app/domain/order/biz"
	"github.com/blackhorseya/godine/app/domain/order/repo/order"
	biz2 "github.com/blackhorseya/godine/app/domain/restaurant/biz"
	biz3 "github.com/blackhorseya/godine/app/domain/user/biz"
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/app/infra/storage/mongodbx"
	"github.com/blackhorseya/godine/app/infra/transports/httpx"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/logging"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func initApplication() (*configx.Application, error) {
	app, err := configx.LoadApplication(&configx.C.OrderRestful)
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

	biz.NewOrderBiz,
	biz2.NewRestaurantHTTPClient,
	biz2.NewMenuHTTPClient,
	biz3.NewUserHTTPClient,
	order.NewMongodb,
	mongodbx.NewClient,
	biz4.NewLogisticsHTTPClient,
)

func New(v *viper.Viper) (adapterx.Restful, error) {
	panic(wire.Build(providerSet))
}
