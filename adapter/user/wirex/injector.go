package wirex

import (
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/entity/domain/user/biz"
)

// Injector defines the injector struct.
type Injector struct {
	A *configx.Application

	UserService biz.IUserBiz
}
