package notifications

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/godine/adapter/notify/wirex"
	"github.com/blackhorseya/godine/entity/notification/biz"
	_ "github.com/blackhorseya/godine/entity/notification/biz" // import biz
	"github.com/blackhorseya/godine/entity/notification/model"
	_ "github.com/blackhorseya/godine/entity/notification/model" // import model
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"github.com/blackhorseya/godine/pkg/responsex"
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
	To      string `json:"to" binding:"required"`
	OrderID string `json:"order_id" binding:"required"`
	Message string `json:"message" binding:"required"`
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
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var payload PostPayload
	err = c.ShouldBindJSON(&payload)
	if err != nil {
		_ = c.Error(errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	notify := model.NewNotify("", payload.To, payload.OrderID, payload.Message)
	err = i.injector.NotifyService.CreateNotification(ctx, notify)
	if err != nil {
		_ = c.Error(err)
		return
	}

	responsex.OK(c, notify)
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
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var options biz.ListNotificationsOptions
	err = c.ShouldBindQuery(&options)
	if err != nil {
		_ = c.Error(errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	items, total, err := i.injector.NotifyService.ListNotificationsByUser(ctx, "", options)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Header("X-Total-Count", strconv.Itoa(total))
	responsex.OK(c, items)
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
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	id := c.Param("id")
	item, err := i.injector.NotifyService.GetNotification(ctx, id)
	if err != nil {
		_ = c.Error(err)
		return
	}

	responsex.OK(c, item)
}
