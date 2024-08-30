package platform

import (
	"encoding/gob"
	"errors"
	"net/http"
	"time"

	"github.com/blackhorseya/godine/adapter/platform/web/templates"
	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
	"github.com/blackhorseya/godine/app/infra/transports/httpx"
	"github.com/blackhorseya/godine/entity/domain/restaurant/biz/bizconnect"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	injector   *Injector
	grpcserver *grpcx.Server
	web        *httpx.Server
	httpserver *http.Server
}

// NewServer creates and returns a new grpcserver.
func NewServer(injector *Injector, grpcserver *grpcx.Server, web *httpx.Server) adapterx.Restful {
	return &impl{
		injector:   injector,
		grpcserver: grpcserver,
		web:        web,
		httpserver: &http.Server{
			Addr:              ":8080",
			ReadHeaderTimeout: time.Second,
			ReadTimeout:       10 * time.Second,
			WriteTimeout:      10 * time.Second,
			MaxHeaderBytes:    8 * 1024,
		},
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

	err = i.web.Start(ctx)
	if err != nil {
		ctx.Error("Failed to start http grpcserver", zap.Error(err))
		return err
	}

	go func() {
		if err = i.httpserver.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			ctx.Error("Failed to start http server", zap.Error(err))
		}
	}()

	return nil
}

func (i *impl) AwaitSignal() error {
	ctx := contextx.Background()
	ctx.Info("receive signal to stop grpcserver")

	if err := i.web.Stop(ctx); err != nil {
		ctx.Error("Failed to stop web", zap.Error(err))
		return err
	}

	if err := i.grpcserver.Stop(ctx); err != nil {
		ctx.Error("Failed to stop grpcserver", zap.Error(err))
		return err
	}

	if err := i.httpserver.Close(); err != nil {
		ctx.Error("Failed to close http server", zap.Error(err))
		return err
	}

	return nil
}

func (i *impl) InitRouting() error {
	router := i.web.Router

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

	// grpc
	api := http.NewServeMux()
	api.Handle(bizconnect.NewRestaurantServiceHandler(i.injector.RestaurantServiceHandler))

	mux := http.NewServeMux()
	mux.Handle("/grpc/", http.StripPrefix("/grpc", api))

	i.httpserver.Handler = mux

	return nil
}

func (i *impl) GetRouter() *gin.Engine {
	return i.web.Router
}
