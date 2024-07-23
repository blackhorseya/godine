package wirex

import (
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/entity/domain/payment/biz"
)

// Injector is used to inject restaurant service.
type Injector struct {
	C *configx.Configuration
	A *configx.Application

	PaymentService biz.IPaymentBiz
}
