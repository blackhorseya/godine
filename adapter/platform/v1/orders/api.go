package orders

import (
	"github.com/blackhorseya/godine/adapter/platform/wirex"
	"github.com/gin-gonic/gin"
)

// Handler will handle the platform v1 orders api.
func Handler(g *gin.RouterGroup, injector *wirex.Injector) {
	orders := g.Group("orders")
	{
		orders.GET("")
		orders.GET("/:id")
		orders.POST("")
	}
}
