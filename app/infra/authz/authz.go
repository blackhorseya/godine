package authz

import (
	"fmt"
	"net/http"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"github.com/blackhorseya/godine/pkg/responsex"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Authz is the authorization struct.
type Authz struct {
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

// ProtectRouter is used to protect the router.
func (a *Authz) ProtectRouter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !a.enabled {
			c.Next()
			return
		}

		ctx, err := contextx.FromGin(c)
		if err != nil {
			_ = c.Error(err)
			return
		}

		by, err := model.FromContext(ctx)
		if err != nil {
			responsex.Err(c, errorx.Wrap(http.StatusUnauthorized, 401, err))
			c.Abort()
			return
		}

		subject := by.GetSubject()
		path := c.Request.URL.Path // example: /api/v1/restaurants
		method := c.Request.Method // example: GET

		allowed, err := a.Enforcer.Enforce(subject, path, method)
		if err != nil {
			_ = c.Error(err)
			c.Abort()
			return
		}

		ctx.Debug("authz",
			zap.String("subject", subject),
			zap.String("path", path),
			zap.String("method", method),
			zap.Bool("allowed", allowed))

		if !allowed {
			responsex.Err(c, errorx.New(http.StatusForbidden, 403, "forbidden"))
			c.Abort()
			return
		}

		c.Next()
	}
}
