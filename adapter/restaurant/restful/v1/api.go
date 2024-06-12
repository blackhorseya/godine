package v1

import (
	"github.com/blackhorseya/godine/adapter/restaurant/restful/v1/restaurants"
	"github.com/blackhorseya/godine/adapter/restaurant/wirex"
	"github.com/gin-gonic/gin"
)

// Handle is used to handle the v1 restful API.
func Handle(g *gin.RouterGroup, injector *wirex.Injector) {
	v1 := g.Group("/v1")
	{
		restaurants.Handle(v1, injector)
	}
}
