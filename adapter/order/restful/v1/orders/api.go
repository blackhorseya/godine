package orders

import (
	"net/http"

	"github.com/blackhorseya/godine/adapter/order/wirex"
	"github.com/blackhorseya/godine/entity/order/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"github.com/blackhorseya/godine/pkg/responsex"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	}
}

// PostPayload is the post payload
type PostPayload struct {
	UserID       string `json:"user_id" binding:"required" example:"8db96c75-cb3c-4c00-8bab-fd088c986f51"`
	RestaurantID string `json:"restaurant_id" binding:"required" example:"a1dbb32b-05f0-4354-8253-60f4c6deae12"`
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

	userID, err := uuid.Parse(payload.UserID)
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	restaurantID, err := uuid.Parse(payload.RestaurantID)
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	order, err := i.injector.OrderService.CreateOrder(ctx, userID, restaurantID, nil, model.Address{}, 0)
	if err != nil {
		responsex.Err(c, err)
		return
	}

	responsex.OK(c, order)
}