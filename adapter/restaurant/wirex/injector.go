package wirex

import (
	"github.com/blackhorseya/godine/app/infra/configx"
)

// Injector is used to inject restaurant service.
type Injector struct {
	A *configx.Application
}