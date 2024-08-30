package user

import (
	"github.com/blackhorseya/godine/app/infra/transports/httpx"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	server *httpx.Server
}

// NewRestful is to create a new restful adapter
func NewRestful(server *httpx.Server) adapterx.Restful {
	return &impl{server: server}
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
	// TODO: 2024/8/30|sean|implement me
	return nil
}

func (i *impl) GetRouter() *gin.Engine {
	return i.server.Router
}
