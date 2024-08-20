package wirex

import (
	"github.com/blackhorseya/godine/app/infra/configx"
)

// Injector is used to inject the dependencies.
type Injector struct {
	C *configx.Configuration
	A *configx.Application
}
