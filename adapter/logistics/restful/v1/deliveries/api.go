package deliveries

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/godine/adapter/logistics/wirex"
	"github.com/blackhorseya/godine/entity/domain/logistics/biz"
	"github.com/blackhorseya/godine/entity/domain/logistics/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"github.com/blackhorseya/godine/pkg/responsex"
	"github.com/gin-gonic/gin"
)

type impl struct {
	injector *wirex.Injector
}

// Handle is used to handle the deliveries restful api
func Handle(g *gin.RouterGroup, injector *wirex.Injector) {
	instance := &impl{injector: injector}

	deliveries := g.Group("/deliveries")
	{
		deliveries.GET("", instance.GetList)
		deliveries.GET("/:id", instance.GetByID)
		deliveries.POST("", instance.Post)
		deliveries.PATCH("/:id/status", instance.PatchWithStatus)
	}
}

// GetList is used to get the list of deliveries
// @Summary Get the list of deliveries
// @Description Get the list of deliveries
// @Tags deliveries
// @Accept json
// @Produce json
// @Param driver_id query string false "driver id: adcf23bc-cd32-4176-8d46-68f15ebdfa98"
// @Param params query biz.ListDeliveriesOptions false "search params"
// @Security Bearer
// @Success 200 {object} responsex.Response{data=[]model.Delivery}
// @Failure 400 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Header 200 {number} X-Total-Count "total count"
// @Router /v1/deliveries [get]
func (i *impl) GetList(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var options biz.ListDeliveriesOptions
	err = c.ShouldBindQuery(&options)
	if err != nil {
		_ = c.Error(errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	driverID := c.Param("driver_id")

	items, total, err := i.injector.LogisticsService.ListDeliveriesByDriver(ctx, driverID, options)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Header("X-Total-Count", strconv.Itoa(total))
	responsex.OK(c, items)
}

// GetByID is used to get the delivery by id
// @Summary Get the delivery by id
// @Description Get the delivery by id
// @Tags deliveries
// @Accept json
// @Produce json
// @Param id path string true "delivery id"
// @Security Bearer
// @Success 200 {object} responsex.Response{data=model.Delivery}
// @Failure 400 {object} responsex.Response
// @Failure 404 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Router /v1/deliveries/{id} [get]
func (i *impl) GetByID(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	id := c.Param("id")
	item, err := i.injector.LogisticsService.GetDelivery(ctx, id)
	if err != nil {
		_ = c.Error(err)
		return
	}

	responsex.OK(c, item)
}

// Post is used to create a new delivery
// @Summary Create a new delivery
// @Description Create a new delivery
// @Tags deliveries
// @Accept json
// @Produce json
// @Param request body model.Delivery true "delivery request"
// @Security Bearer
// @Success 200 {object} responsex.Response{data=model.Delivery}
// @Failure 400 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Router /v1/deliveries [post]
func (i *impl) Post(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var payload *model.Delivery
	err = c.ShouldBindJSON(&payload)
	if err != nil {
		_ = c.Error(errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	err = i.injector.LogisticsService.CreateDelivery(ctx, payload)
	if err != nil {
		_ = c.Error(err)
		return
	}

	responsex.OK(c, payload)
}

// PatchWithStatusPayload defines the payload for patch with status
type PatchWithStatusPayload struct {
	Status string `json:"status" binding:"required" example:"delivered"`
}

// PatchWithStatus is used to update the delivery status
// @Summary Update the delivery status
// @Description Update the delivery status
// @Tags deliveries
// @Accept json
// @Produce json
// @Param id path string true "delivery id"
// @Param request body PatchWithStatusPayload true "delivery status request"
// @Security Bearer
// @Success 200 {object} responsex.Response{data=model.Delivery}
// @Failure 400 {object} responsex.Response
// @Failure 404 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Router /v1/deliveries/{id}/status [patch]
func (i *impl) PatchWithStatus(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	id := c.Param("id")

	var payload PatchWithStatusPayload
	err = c.ShouldBindJSON(&payload)
	if err != nil {
		_ = c.Error(errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	err = i.injector.LogisticsService.UpdateDeliveryStatus(ctx, id, payload.Status)
	if err != nil {
		_ = c.Error(err)
		return
	}

	item, err := i.injector.LogisticsService.GetDelivery(ctx, id)
	if err != nil {
		_ = c.Error(err)
		return
	}

	responsex.OK(c, item)
}
