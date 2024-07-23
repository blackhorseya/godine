package authz

import (
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

// Authz is the authorization struct.
type Authz struct {
	*casbin.Enforcer
}

// New is used to create a new authorization.
func New(app *configx.Application) (*Authz, error) {
	adapter, err := gormadapter.NewAdapter(app.Casbin.PolicyDriver, app.Casbin.PolicyDSN, true)

	enforcer, err := casbin.NewEnforcer(app.Casbin.ModelPath, adapter)
	if err != nil {
		return nil, err
	}

	return &Authz{
		Enforcer: enforcer,
	}, nil
}
