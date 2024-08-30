package platform

import (
	"encoding/gob"
	"errors"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	"github.com/blackhorseya/godine/adapter/platform/web/templates"
	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
	"github.com/blackhorseya/godine/app/infra/transports/httpx"
	"github.com/blackhorseya/godine/entity/domain/restaurant/biz/bizconnect"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
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
	compress1KB := connect.WithCompressMinBytes(1024)
	api := http.NewServeMux()
	api.Handle(bizconnect.NewRestaurantServiceHandler(i.injector.RestaurantServiceHandler, compress1KB))

	mux := http.NewServeMux()
	mux.Handle("/grpc/", http.StripPrefix("/grpc", api))
	mux.Handle(grpchealth.NewHandler(grpchealth.NewStaticChecker(bizconnect.RestaurantServiceName), compress1KB))
	mux.Handle(grpcreflect.NewHandlerV1(grpcreflect.NewStaticReflector(bizconnect.RestaurantServiceName), compress1KB))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(
		grpcreflect.NewStaticReflector(bizconnect.RestaurantServiceName),
		compress1KB,
	))

	i.httpserver.Handler = h2c.NewHandler(newCORS().Handler(mux), &http2.Server{})

	return nil
}

func (i *impl) GetRouter() *gin.Engine {
	return i.web.Router
}

func newCORS() *cors.Cors {
	// To let web developers play with the demo service from browsers, we need a
	// very permissive CORS setup.
	return cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowOriginFunc: func(origin string) bool {
			// Allow all origins, which effectively disables CORS.
			return true
		},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{
			// Content-Type is in the default safelist.
			"Accept",
			"Accept-Encoding",
			"Accept-Post",
			"Connect-Accept-Encoding",
			"Connect-Content-Encoding",
			"Content-Encoding",
			"Grpc-Accept-Encoding",
			"Grpc-Encoding",
			"Grpc-Message",
			"Grpc-Status",
			"Grpc-Status-Details-Bin",
		},
		// Let browsers cache CORS information for longer, which reduces the number
		// of preflight requests. Any changes to ExposedHeaders won't take effect
		// until the cached data expires. FF caps this value at 24h, and modern
		// Chrome caps it at 2h.
		MaxAge: int(2 * time.Hour / time.Second),
	})
}
