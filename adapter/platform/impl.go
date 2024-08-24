package platform

import (
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/blackhorseya/godine/adapter/platform/web/templates"
	"github.com/blackhorseya/godine/api"
	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
	"github.com/blackhorseya/godine/app/infra/transports/httpx"
	restB "github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

type impl struct {
	injector   *Injector
	grpcserver *grpcx.Server
	httpserver *httpx.Server
}

// NewServer creates and returns a new grpcserver.
func NewServer(injector *Injector, grpcserver *grpcx.Server, httpserver *httpx.Server) adapterx.Restful {
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
	router.GET("/login", i.login)
	router.GET("/callback", i.callback)
	router.GET("/user", IsAuthenticated, i.user)
	router.GET("/logout", i.logout)

	// api
	gw := runtime.NewServeMux()
	err := restB.RegisterRestaurantServiceHandlerClient(contextx.Background(), gw, i.injector.RestaurantClient)
	if err != nil {
		return fmt.Errorf("failed to register restaurant service handler client: %w", err)
	}
	router.Any("/api/v1/*any", gin.WrapH(gw))

	// swagger
	router.GET("/swagger", func(c *gin.Context) {
		data, err2 := api.GatewayOpenAPI.ReadFile("gateway/apidocs.swagger.json")
		if err2 != nil {
			c.String(http.StatusInternalServerError, "failed to read swagger file")
			return
		}
		c.Data(http.StatusOK, "application/json; charset=utf-8", data)
	})
	router.GET("/api/docs/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.URL("/swagger"),
	))

	return nil
}

func (i *impl) GetRouter() *gin.Engine {
	return i.httpserver.Router
}
