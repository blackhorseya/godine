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

// ProtectRouter is used to protect the router.
func (a *Authz) ProtectRouter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, err := contextx.FromGin(c)
		if err != nil {
			_ = c.Error(err)
			return
		}

		by, err := model.FromContext(ctx)
		if err != nil {
			responsex.Err(c, errorx.Wrap(http.StatusUnauthorized, 401, err))
			c.Abort()
		}

		subject := by.GetSubject()
		method := c.Request.Method
		path := c.Request.URL.Path

		allowed, err := a.Enforcer.Enforce(subject, path, method)
		if err != nil {
			_ = c.Error(err)
			c.Abort()
		}

		ctx.Debug("authz",
			zap.String("subject", subject),
			zap.String("method", method),
			zap.String("path", path),
			zap.Bool("allowed", allowed))

		if !allowed {
			responsex.Err(c, errorx.New(http.StatusForbidden, 403, "forbidden"))
			c.Abort()
		}

		c.Next()
	}
}