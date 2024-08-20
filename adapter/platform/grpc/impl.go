package grpc

import (
	"encoding/gob"
	"net/http"

	"github.com/blackhorseya/godine/adapter/platform/grpc/web/templates"
	"github.com/blackhorseya/godine/adapter/platform/wirex"
	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
	"github.com/blackhorseya/godine/app/infra/transports/httpx"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	injector   *wirex.Injector
	grpcserver *grpcx.Server
	httpserver *httpx.Server
}

// NewServer creates and returns a new grpcserver.
func NewServer(injector *wirex.Injector, grpcserver *grpcx.Server, httpserver *httpx.Server) adapterx.Restful {
	return &impl{
		injector:   injector,
		grpcserver: grpcserver,
		httpserver: httpserver,
	}
}

func (i *impl) Start() error {
	ctx := contextx.Background()
	err := i.grpcserver.Start(ctx)
	if err != nil {
		ctx.Error("Failed to start grpc grpcserver", zap.Error(err))
		return err
	}

	err = i.InitRouting()
	if err != nil {
		ctx.Error("Failed to init routing", zap.Error(err))
		return err
	}

	err = i.httpserver.Start(ctx)
	if err != nil {
		ctx.Error("Failed to start http grpcserver", zap.Error(err))
		return err
	}

	return nil
}

func (i *impl) AwaitSignal() error {
	ctx := contextx.Background()
	ctx.Info("receive signal to stop grpcserver")

	if err := i.httpserver.Stop(ctx); err != nil {
		ctx.Error("Failed to stop httpserver", zap.Error(err))
		return err
	}

	if err := i.grpcserver.Stop(ctx); err != nil {
		ctx.Error("Failed to stop grpcserver", zap.Error(err))
		return err
	}

	return nil
}

func (i *impl) InitRouting() error {
	router := i.httpserver.Router

	gob.Register(map[string]interface{}{})
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	templates.SetHTMLTemplate(router)

	// web
	router.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})
	// router.GET("/login", i.login)
	// router.GET("/callback", i.callback)
	// router.GET("/user", IsAuthenticated, i.user)
	// router.GET("/logout", i.logout)

	return nil
}

func (i *impl) GetRouter() *gin.Engine {
	return i.httpserver.Router
}
