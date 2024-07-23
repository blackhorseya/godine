package v1

import (
	"github.com/blackhorseya/godine/adapter/user/restful/v1/users"
	"github.com/blackhorseya/godine/adapter/user/wirex"
	"github.com/gin-gonic/gin"
)

// Handle is used to handle the v1 restful API.
func Handle(g *gin.RouterGroup, injector *wirex.Injector) {
	v1 := g.Group("/v1", injector.Authz.ProtectRouter())
	{
		users.Handle(v1, injector)
	}
}
