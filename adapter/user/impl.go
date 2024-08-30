package user

import (
	"github.com/blackhorseya/godine/app/infra/transports/httpx"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/gin-gonic/gin"
)

type impl struct {
	server *httpx.Server
}

// NewRestful is to create a new restful adapter
func NewRestful(server *httpx.Server) adapterx.Restful {
	return &impl{server: server}
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
