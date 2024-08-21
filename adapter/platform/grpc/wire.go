//go:build wireinject

//go:generate wire

package grpc

import (
	"fmt"

	"github.com/blackhorseya/godine/adapter/platform/wirex"
	biz2 "github.com/blackhorseya/godine/app/domain/restaurant/biz"
	"github.com/blackhorseya/godine/app/domain/user/biz"
	"github.com/blackhorseya/godine/app/infra/authx"
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
	"github.com/blackhorseya/godine/app/infra/transports/httpx"
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
) grpcx.InitServers {
	return func(s *grpc.Server) {
		healthServer := health.NewServer()
		grpc_health_v1.RegisterHealthServer(s, healthServer)
		healthServer.SetServingStatus(serverName, grpc_health_v1.HealthCheckResponse_SERVING)

		userB.RegisterAccountServiceServer(s, accountServer)
		restB.RegisterRestaurantServiceServer(s, restaurantServer)
		restB.RegisterMenuServiceServer(s, menuServer)

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

func New(v *viper.Viper) (adapterx.Restful, error) {
	panic(wire.Build(
		NewServer,
		wire.Struct(new(wirex.Injector), "*"),
		grpcx.NewServer,
		httpx.NewServer,
		initApplication,
		configx.NewConfiguration,
		NewInitServersFn,
		initAuthx,

		biz.NewAccountService,
		biz2.NewRestaurantService,
		biz2.NewMenuService,
	))
}
