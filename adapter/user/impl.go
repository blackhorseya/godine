package user

import (
	"encoding/gob"
	"net/http"

	"github.com/blackhorseya/godine/adapter/user/web/templates"
	"github.com/blackhorseya/godine/app/infra/transports/httpx"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	injector *Injector
	server   *httpx.Server
}

// NewRestful is to create a new restful adapter
func NewRestful(injector *Injector, server *httpx.Server) adapterx.Restful {
	return &impl{
		injector: injector,
		server:   server,
	}
}

func (i *impl) Start() error {
	ctx := contextx.Background()

	err := i.InitRouting()
	if err != nil {
		ctx.Error("Failed to init routing", zap.Error(err))
		return err
	}

	err = i.server.Start(ctx)
	if err != nil {
		ctx.Error("Failed to start http server", zap.Error(err))
		return err
	}

	return nil
}

func (i *impl) AwaitSignal() error {
	ctx := contextx.Background()
	ctx.Info("receive signal to stop grpcserver")

	if err := i.server.Stop(ctx); err != nil {
		ctx.Error("Failed to close http server", zap.Error(err))
		return err
	}

	return nil
}

func (i *impl) InitRouting() error {
	router := i.server.Router

	gob.Register(map[string]interface{}{})
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	templates.SetHTMLTemplate(router)

	// web
	router.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})
	router.GET("/login", i.login)
	router.GET("/callback", i.callback)
	router.GET("/user", IsAuthenticated, i.user)
	router.GET("/logout", i.logout)

	return nil
}

func (i *impl) GetRouter() *gin.Engine {
	return i.server.Router
}
