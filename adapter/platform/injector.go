package platform

import (
	"github.com/blackhorseya/godine/app/infra/authx"
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/restaurant/biz/bizconnect"
)

// Injector is used to inject the dependencies.
type Injector struct {
	C     *configx.Configuration
	A     *configx.Application
	OTelx *otelx.SDK

	Authx *authx.Authx

	RestaurantServiceHandler bizconnect.RestaurantServiceHandler
}
