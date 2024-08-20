package grpc

import (
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/gin-gonic/gin"
)

type impl struct {
}

// NewServer creates and returns a new server.
func NewServer() adapterx.Restful {
	return &impl{}
}

func (i *impl) Start() error {
	// TODO: 2024/8/21|sean|implement grpc server start
	return nil
}

func (i *impl) AwaitSignal() error {
	ctx := contextx.Background()
	ctx.Info("receive signal to stop server")

	// TODO: 2024/8/21|sean|implement grpc server stop

	return nil
}

func (i *impl) InitRouting() error {
	return nil
}

func (i *impl) GetRouter() *gin.Engine {
	return nil
}
