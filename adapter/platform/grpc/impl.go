package grpc

import (
	"github.com/blackhorseya/godine/app/infra/transports/grpcx"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	server *grpcx.Server
}

// NewServer creates and returns a new server.
func NewServer(server *grpcx.Server) adapterx.Restful {
	return &impl{
		server: server,
	}
}

func (i *impl) Start() error {
	ctx := contextx.Background()
	err := i.server.Start(ctx)
	if err != nil {
		ctx.Error("Failed to start grpc server", zap.Error(err))
		return err
	}

	return nil
}

func (i *impl) AwaitSignal() error {
	ctx := contextx.Background()
	ctx.Info("receive signal to stop server")

	if err := i.server.Stop(ctx); err != nil {
		ctx.Error("Failed to stop server", zap.Error(err))
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
