package orders

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/godine/adapter/order/wirex"
	"github.com/blackhorseya/godine/entity/domain/order/biz"
	model2 "github.com/blackhorseya/godine/entity/domain/order/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"github.com/blackhorseya/godine/pkg/responsex"
	"github.com/gin-gonic/gin"
)

type impl struct {
	injector *wirex.Injector
}

// Handle is used to handle the orders restful api
func Handle(g *gin.RouterGroup, injector *wirex.Injector) {
	i := &impl{injector: injector}

	group := g.Group("/orders")
	{
		group.POST("", i.Post)
		group.GET("", i.GetList)
		group.GET("/:order_id", i.GetByID)
		group.PATCH("/:order_id/status", i.PatchWithStatus)
	}
}

// PostPayload is the post payload
type PostPayload struct {
	UserID       string             `json:"user_id" binding:"required" example:"adcf23bc-cd32-4176-8d46-68f15ebdfa98"`
	RestaurantID string             `json:"restaurant_id" binding:"required"`
	Items        []model2.OrderItem `json:"items" binding:"required"`
}

// Post is the post method
// @Summary Create a new order
// @Description Create a new order
// @Tags orders
// @Accept json
// @Produce json
// @Param payload body PostPayload true "order payload"
// @Security Bearer
// @Success 200 {object} responsex.Response{data=model.Order}
// @Failure 400 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Router /v1/orders [post]
func (i *impl) Post(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var payload PostPayload
	err = c.ShouldBindJSON(&payload)
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	order, err := i.injector.OrderService.CreateOrder(
		ctx,
		payload.UserID,
		payload.RestaurantID,
		payload.Items,
		model2.Address{},
		0,
	)
	if err != nil {
		responsex.Err(c, err)
		return
	}

	responsex.OK(c, order)
}

// GetList is the get list method
// @Summary Get order list
// @Description Get order list
// @Tags orders
// @Accept json
// @Produce json
// @Param query query biz.ListOrdersOptions false "list order options"
// @Security Bearer
// @Success 200 {object} responsex.Response{data=[]model.Order}
// @Failure 500 {object} responsex.Response
// @Header 200 {int} X-Total-Count "Total number of items"
// @Router /v1/orders [get]
func (i *impl) GetList(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var query biz.ListOrdersOptions
	err = c.ShouldBindQuery(&query)
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	orders, total, err := i.injector.OrderService.ListOrders(ctx, query)
	if err != nil {
		responsex.Err(c, err)
		return
	}

	c.Header("X-Total-Count", strconv.Itoa(total))
	responsex.OK(c, orders)
}

// GetByID is the get by id method
// @Summary Get order by id
// @Description Get order by id
// @Tags orders
// @Accept json
// @Produce json
// @Param order_id path string true "order id"
// @Security Bearer
// @Success 200 {object} responsex.Response{data=model.Order}
// @Failure 404 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Router /v1/orders/{order_id} [get]
func (i *impl) GetByID(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	orderID := c.Param("order_id")
	order, err := i.injector.OrderService.GetOrder(ctx, orderID)
	if err != nil {
		responsex.Err(c, err)
		return
	}

	responsex.OK(c, order)
}

// PatchWithStatusPayload is the patch with status payload
type PatchWithStatusPayload struct {
	Status string `json:"status" binding:"required" example:"confirmed"`
}

// PatchWithStatus is the patch with status method
// @Summary Update order status
// @Description Update order status
// @Tags orders
// @Accept json
// @Produce json
// @Param order_id path string true "order id"
// @Param payload body PatchWithStatusPayload true "order status payload"
// @Security Bearer
// @Success 200 {object} responsex.Response{data=model.Order}
// @Failure 400 {object} responsex.Response
// @Failure 404 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Router /v1/orders/{order_id}/status [patch]
func (i *impl) PatchWithStatus(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	orderID := c.Param("order_id")

	var payload PatchWithStatusPayload
	err = c.ShouldBindJSON(&payload)
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	err = i.injector.OrderService.UpdateOrderStatus(ctx, orderID, payload.Status)
	if err != nil {
		responsex.Err(c, err)
		return
	}

	order, err := i.injector.OrderService.GetOrder(ctx, orderID)
	if err != nil {
		responsex.Err(c, err)
		return
	}

	responsex.OK(c, order)
}
