package grpc

import (
	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
	"github.com/blackhorseya/godine/app/infra/transports/httpx"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	grpcserver *grpcx.Server
	httpserver *httpx.Server
}

// NewServer creates and returns a new grpcserver.
func NewServer(grpcserver *grpcx.Server, httpserver *httpx.Server) adapterx.Restful {
	return &impl{
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
	return nil
}

func (i *impl) GetRouter() *gin.Engine {
	return nil
}
