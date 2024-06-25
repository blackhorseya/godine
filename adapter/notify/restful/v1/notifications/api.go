package notifications

import (
	"github.com/blackhorseya/godine/adapter/notify/wirex"
	_ "github.com/blackhorseya/godine/entity/notification/biz"   // import biz
	_ "github.com/blackhorseya/godine/entity/notification/model" // import model
	"github.com/gin-gonic/gin"
)

type impl struct {
	injector *wirex.Injector
}

// Handle defines the API routes for the notification service.
func Handle(g *gin.RouterGroup, injector *wirex.Injector) {
	instance := &impl{injector: injector}

	group := g.Group("/notifications")
	{
		group.POST("", instance.Post)
		group.GET("", instance.GetList)
		group.GET("/:id", instance.GetByID)
	}
}

// PostPayload defines the request payload for creating a new notification.
type PostPayload struct {
}

// Post creates a new notification.
// @Summary Create a new notification
// @Description Create a new notification.
// @Tags notifications
// @Accept json
// @Produce json
// @Security Bearer
// @Param payload body PostPayload true "The request payload"
// @Success 201 {object} responsex.Response{data=model.Notification}
// @Failure 400 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Router /v1/notifications [post]
func (i *impl) Post(c *gin.Context) {
	// todo: 2024/6/26|sean|implement the post notification handler
}

// GetList retrieves a list of notifications.
// @Summary Retrieve a list of notifications
// @Description Retrieve a list of notifications.
// @Tags notifications
// @Accept json
// @Produce json
// @Security Bearer
// @Param params query biz.ListNotificationsOptions false "The list options"
// @Success 200 {object} responsex.Response{data=[]model.Notification}
// @Failure 400 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Header 200 {int} X-Total-Count "Total number of items"
// @Router /v1/notifications [get]
func (i *impl) GetList(c *gin.Context) {
	// todo: 2024/6/26|sean|implement the get list notification handler
}

// GetByID retrieves a notification by its ID.
// @Summary Retrieve a notification by ID
// @Description Retrieve a notification by ID.
// @Tags notifications
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "The notification ID"
// @Success 200 {object} responsex.Response{data=model.Notification}
// @Failure 400 {object} responsex.Response
// @Failure 404 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Router /v1/notifications/{id} [get]
func (i *impl) GetByID(c *gin.Context) {
	// todo: 2024/6/26|sean|implement the get by id notification handler
}
