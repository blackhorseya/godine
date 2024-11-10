package v1

import (
	"github.com/blackhorseya/godine/adapter/platform/v1/orders"
	"github.com/gin-gonic/gin"
)

// Handler will handle the platform v1 api.
func Handler(g *gin.RouterGroup) {
	v1 := g.Group("v1")
	{
		orders.Handler(v1)
	}
}
