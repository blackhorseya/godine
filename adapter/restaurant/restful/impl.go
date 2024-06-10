package restful

import (
	"github.com/blackhorseya/godine/adapter/restaurant/wirex"
	"github.com/blackhorseya/godine/app/infra/transports/httpx"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/gin-gonic/gin"
)

type impl struct {
	injector *wirex.Injector
	server   *httpx.Server
}

func newRestful(injector *wirex.Injector, server *httpx.Server) adapterx.Restful {
	return &impl{injector: injector, server: server}
}

func (i *impl) Start() error {
	// TODO implement me
	panic("implement me")
}

func (i *impl) AwaitSignal() error {
	// TODO implement me
	panic("implement me")
}

func (i *impl) InitRouting() error {
	// TODO implement me
	panic("implement me")
}

func (i *impl) GetRouter() *gin.Engine {
	// TODO implement me
	panic("implement me")
}
