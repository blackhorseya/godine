package v1

import (
	"github.com/blackhorseya/godine/adapter/logistics/restful/v1/deliveries"
	"github.com/blackhorseya/godine/adapter/logistics/wirex"
	"github.com/gin-gonic/gin"
)

// Handle is used to handle the v1 restful api
func Handle(g *gin.RouterGroup, injector *wirex.Injector) {
	v1 := g.Group("/v1")
	{
		deliveries.Handle(v1, injector)
	}
}
