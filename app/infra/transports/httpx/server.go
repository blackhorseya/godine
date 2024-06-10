package httpx

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/responsex"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/zap"
)

// Server is an HTTP server.
type Server struct {
	httpserver *http.Server
	Router     *gin.Engine
}

// NewServer is used to create a new HTTP server.
func NewServer(app *configx.Application) (*Server, error) {
	ctx := contextx.Background()

	gin.SetMode(app.HTTP.Mode)

	router := gin.New()
	router.Use(ginzap.GinzapWithConfig(ctx.Logger, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		SkipPaths:  nil,
		Context:    nil,
	}))
	router.Use(otelgin.Middleware(app.Name))
	router.Use(contextx.AddContextxMiddleware())
	router.Use(responsex.AddErrorHandlingMiddleware())
	router.Use(ginzap.CustomRecoveryWithZap(ctx.Logger, true, func(c *gin.Context, err any) {
		responsex.Err(c, fmt.Errorf("%v", err))
		c.Abort()
	}))

	httpserver := &http.Server{
		Addr:              app.HTTP.GetAddr(),
		Handler:           router,
		ReadHeaderTimeout: time.Second,
	}

	return &Server{
		httpserver: httpserver,
		Router:     router,
	}, nil
}

// Start is used to start the server.
func (s *Server) Start(ctx contextx.Contextx) error {
	go func() {
		err := s.httpserver.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			ctx.Fatal("start http server error", zap.Error(err))
		}
	}()

	return nil
}

// Stop is used to stop the server.
func (s *Server) Stop(ctx contextx.Contextx) error {
	timeout, cancelFunc := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()

	err := s.httpserver.Shutdown(timeout)
	if err != nil {
		return err
	}

	return nil
}
