package wirex

import (
	"github.com/blackhorseya/godine/app/infra/authx"
	"github.com/blackhorseya/godine/app/infra/authz"
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/entity/domain/user/biz"
)

// Injector defines the injector struct.
type Injector struct {
	C *configx.Configuration
	A *configx.Application

	Authx *authx.Authx
	Authz *authz.Authz

	UserService biz.IUserBiz
}
