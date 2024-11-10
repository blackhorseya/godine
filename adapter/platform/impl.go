package platform

import (
	"context"
	"errors"
	"net/http"

	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
	"github.com/blackhorseya/godine/app/infra/transports/httpx"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	injector   *Injector
	grpcserver *grpcx.Server
	httpserver *httpx.Server
}

// NewServer creates and returns a new grpcserver.
func NewServer(injector *Injector, grpcserver *grpcx.Server, httpserver *httpx.Server) adapterx.Server {
	return &impl{
		injector:   injector,
		grpcserver: grpcserver,
		httpserver: httpserver,
	}
}

func (i *impl) Start(c context.Context) error {
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

	go func() {
		if err = i.httpserver.Start(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			ctx.Error("Failed to start http server", zap.Error(err))
		}
	}()

	return nil
}

func (i *impl) Shutdown(c context.Context) error {
	ctx := contextx.Background()
	ctx.Info("receive signal to stop grpcserver")

	if err := i.grpcserver.Stop(ctx); err != nil {
		ctx.Error("Failed to stop grpcserver", zap.Error(err))
		return err
	}

	if err := i.httpserver.Stop(ctx); err != nil {
		ctx.Error("Failed to close http server", zap.Error(err))
		return err
	}

	return nil
}

func (i *impl) InitRouting() error {
	// TODO: 2024/11/10|sean|init routing
	return nil
}

func (i *impl) GetRouter() *gin.Engine {
	return i.httpserver.Router
}
