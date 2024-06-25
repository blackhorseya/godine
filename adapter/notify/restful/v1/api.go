package v1

import (
	"github.com/blackhorseya/godine/adapter/notify/restful/v1/notifications"
	"github.com/blackhorseya/godine/adapter/notify/wirex"
	"github.com/gin-gonic/gin"
)

// Handle defines the API routes for the notification service.
func Handle(g *gin.RouterGroup, injector *wirex.Injector) {
	v1 := g.Group("/v1")
	{
		notifications.Handle(v1, injector)
	}
}
