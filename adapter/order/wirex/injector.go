package wirex

import (
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/entity/domain/order/biz"
)

// Injector is a struct that contains all the dependencies needed by the order package.
type Injector struct {
	C *configx.Configuration
	A *configx.Application

	OrderService biz.IOrderBiz
}
