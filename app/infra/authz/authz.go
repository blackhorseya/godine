package authz

import (
	"fmt"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

// Authz is the authorization struct.
type Authz struct {
	*casbin.Enforcer
}

// New is used to create a new authorization.
func New(app *configx.Application) (*Authz, error) {
	if !app.Casbin.Enabled {
		contextx.Background().Warn("casbin is disabled")
		return nil, nil //nolint:nilnil // return nil to indicate casbin is disabled
	}

	var adapter, err = gormadapter.NewAdapter(app.Casbin.PolicyDriver, app.Storage.Mysql.DSN, true)
	if err != nil {
		return nil, fmt.Errorf("failed to create casbin adapter: %w", err)
	}

	enforcer, err := casbin.NewEnforcer(app.Casbin.ModelPath, adapter)
	if err != nil {
		return nil, fmt.Errorf("failed to create casbin enforcer: %w", err)
	}

	return &Authz{
		Enforcer: enforcer,
	}, nil
}
