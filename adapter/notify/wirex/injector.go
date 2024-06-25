package wirex

import (
	"github.com/blackhorseya/godine/app/infra/configx"
)

// Injector is a struct that contains all the dependencies needed by the order package.
type Injector struct {
	A *configx.Application

	// NotifyService biz.INotificationBiz
}
