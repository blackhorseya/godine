package platform

import (
	"github.com/blackhorseya/godine/app/infra/authx"
	"github.com/blackhorseya/godine/app/infra/configx"
)

// Injector is used to inject the dependencies.
type Injector struct {
	C *configx.Configuration
	A *configx.Application

	Authx *authx.Authx
}
