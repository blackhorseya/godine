package user

import (
	"github.com/blackhorseya/godine/app/infra/authx"
	"github.com/blackhorseya/godine/app/infra/configx"
)

// Injector is the user injector.
type Injector struct {
	C *configx.Configuration
	A *configx.Application

	Authx *authx.Authx
}
