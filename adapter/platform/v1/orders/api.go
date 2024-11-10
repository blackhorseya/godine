package orders

import (
	"github.com/gin-gonic/gin"
)

// Handler will handle the platform v1 orders api.
func Handler(g *gin.RouterGroup) {
	orders := g.Group("orders")
	{
		orders.GET("")
		orders.GET("/:id")
		orders.POST("")
	}
}
