package v1

import (
	"github.com/blackhorseya/godine/adapter/order/restful/v1/orders"
	"github.com/blackhorseya/godine/adapter/order/wirex"
	"github.com/gin-gonic/gin"
)

// Handle is used to handle the v1 restful api
func Handle(g *gin.RouterGroup, injector *wirex.Injector) {
	group := g.Group("/v1")
	{
		orders.Handle(group, injector)
	}
}
