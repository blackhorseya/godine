package authz

import (
	"fmt"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/stretchr/testify/mock"
)

// Authz is the authorization struct.
type Authz struct {
	mock.Mock

	enabled bool
	*casbin.Enforcer
}

// New is used to create a new authorization.
func New(app *configx.Application) (*Authz, error) {
	if !app.Casbin.Enabled {
		contextx.Background().Warn("casbin is disabled")
		return &Authz{
			enabled: false,
		}, nil
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
		enabled:  true,
		Enforcer: enforcer,
	}, nil
}

// InitPolicy is used to initialize the policy.
func (a *Authz) InitPolicy() (err error) {
	if !a.enabled {
		return nil
	}

	a.ClearPolicy()

	_, err = a.AddPolicy("owner", "restaurant", "manage")
	if err != nil {
		return err
	}

	_, err = a.AddPolicy("admin", "restaurant", "manage")
	if err != nil {
		return err
	}

	_, err = a.AddPolicy("editor", "restaurant", "edit")
	if err != nil {
		return err
	}

	_, err = a.AddPolicy("viewer", "restaurant", "view")
	if err != nil {
		return err
	}

	return nil
}
