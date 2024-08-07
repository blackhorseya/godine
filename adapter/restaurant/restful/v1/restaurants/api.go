package restaurants

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/godine/adapter/restaurant/restful/v1/restaurants/items"
	"github.com/blackhorseya/godine/adapter/restaurant/wirex"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/restaurant/biz"
	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"github.com/blackhorseya/godine/pkg/responsex"
	"github.com/gin-gonic/gin"
)

type impl struct {
	injector *wirex.Injector
}

// Handle is used to handle the restaurant restful API.
func Handle(g *gin.RouterGroup, injector *wirex.Injector) {
	i := &impl{injector: injector}

	group := g.Group("/restaurants")
	{
		group.GET("", i.GetList)
		group.POST("", i.Post)
		group.GET("/:restaurant_id", i.GetByID)
		group.PUT("/:restaurant_id", i.PutByID)
		group.PATCH("/:restaurant_id/status", i.PatchWithStatus)
		group.DELETE("/:restaurant_id", i.DeleteByID)

		items.Handle(group.Group("/:restaurant_id"), injector)
	}
}

// GetListQuery is the get list query.
type GetListQuery struct {
	Page int `form:"page" default:"1" minimum:"1"`
	Size int `form:"size" default:"10" minimum:"1" maximum:"100"`
}

// GetList is used to get the restaurant list.
// @Summary Get the restaurant list.
// @Description Get the restaurant list.
// @Tags restaurants
// @Accept json
// @Produce json
// @Param query query GetListQuery false "get list query"
// @Success 200 {object} responsex.Response{data=[]model.Restaurant}
// @Failure 500 {object} responsex.Response
// @Header 200 {number} X-Total-Count "total count"
// @Router /v1/restaurants [get]
func (i *impl) GetList(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	ctx, span := otelx.Span(ctx, "api.restaurants.get_list")
	defer span.End()

	var query GetListQuery
	err = c.ShouldBindQuery(&query)
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	ret, total, err := i.injector.RestaurantService.ListRestaurants(ctx, biz.ListRestaurantsOptions{
		Page: query.Page,
		Size: query.Size,
	})
	if err != nil {
		responsex.Err(c, err)
		return
	}

	c.Header("X-Total-Count", strconv.Itoa(total))
	responsex.OK(c, ret)
}

// PostPayload is the post payload.
type PostPayload struct {
	Name        string `json:"name" binding:"required" example:"restaurant name"`
	Description string `json:"description"`
}

// Post is used to create the restaurant.
// @Summary Create the restaurant.
// @Description Create the restaurant.
// @Tags restaurants
// @Accept json
// @Produce json
// @Param payload body PostPayload true "restaurant payload"
// @Success 200 {object} responsex.Response{data=model.Restaurant}
// @Failure 500 {object} responsex.Response
// @Router /v1/restaurants [post]
func (i *impl) Post(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	ctx, span := otelx.Span(ctx, "api.restaurants.post")
	defer span.End()

	var payload PostPayload
	err = c.ShouldBindJSON(&payload)
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	item, err := i.injector.RestaurantService.CreateRestaurant(ctx, payload.Name, payload.Description)
	if err != nil {
		responsex.Err(c, err)
		return
	}

	responsex.OK(c, item)
}

// GetByID is used to get the restaurant by id.
// @Summary Get the restaurant by id.
// @Description Get the restaurant by id.
// @Tags restaurants
// @Accept json
// @Produce json
// @Param restaurant_id path string true "restaurant id"
// @Success 200 {object} responsex.Response{data=model.Restaurant}
// @Failure 500 {object} responsex.Response
// @Router /v1/restaurants/{restaurant_id} [get]
func (i *impl) GetByID(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	ctx, span := otelx.Span(ctx, "api.restaurants.get_by_id")
	defer span.End()

	item, err := i.injector.RestaurantService.GetRestaurant(ctx, c.Param("restaurant_id"))
	if err != nil {
		responsex.Err(c, err)
		return
	}

	responsex.OK(c, item)
}

// PutByID is used to update the restaurant by id.
// @Summary Update the restaurant by id.
// @Description Update the restaurant by id.
// @Tags restaurants
// @Accept json
// @Produce json
// @Param restaurant_id path string true "restaurant id"
// @Param payload body model.Restaurant true "restaurant payload"
// @Success 200 {object} responsex.Response{data=model.Restaurant}
// @Failure 500 {object} responsex.Response
// @Router /v1/restaurants/{restaurant_id} [put]
func (i *impl) PutByID(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var payload model.Restaurant
	err = c.ShouldBindJSON(&payload)
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	err = i.injector.RestaurantService.UpdateRestaurant(ctx, c.Param("restaurant_id"), payload.Name, payload.Address)
	if err != nil {
		responsex.Err(c, err)
		return
	}

	item, err := i.injector.RestaurantService.GetRestaurant(ctx, c.Param("restaurant_id"))
	if err != nil {
		responsex.Err(c, err)
		return
	}

	responsex.OK(c, item)
}

// PatchWithStatusPayload is the patch with status payload.
type PatchWithStatusPayload struct {
	IsOpen bool `json:"is_open" example:"true"`
}

// PatchWithStatus is used to update the restaurant status by id.
// @Summary Update the restaurant status by id.
// @Description Update the restaurant status by id.
// @Tags restaurants
// @Accept json
// @Produce json
// @Param restaurant_id path string true "restaurant id"
// @Param payload body PatchWithStatusPayload true "restaurant status payload"
// @Success 200 {object} responsex.Response{data=model.Restaurant}
// @Failure 500 {object} responsex.Response
// @Router /v1/restaurants/{restaurant_id}/status [patch]
func (i *impl) PatchWithStatus(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var payload PatchWithStatusPayload
	err = c.ShouldBindJSON(&payload)
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	err = i.injector.RestaurantService.ChangeRestaurantStatus(ctx, c.Param("restaurant_id"), payload.IsOpen)
	if err != nil {
		responsex.Err(c, err)
		return
	}

	item, err := i.injector.RestaurantService.GetRestaurant(ctx, c.Param("restaurant_id"))
	if err != nil {
		responsex.Err(c, err)
		return
	}

	responsex.OK(c, item)
}

// DeleteByID is used to delete the restaurant by id.
// @Summary Delete the restaurant by id.
// @Description Delete the restaurant by id.
// @Tags restaurants
// @Accept json
// @Produce json
// @Param restaurant_id path string true "restaurant id"
// @Success 204 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Router /v1/restaurants/{restaurant_id} [delete]
func (i *impl) DeleteByID(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = i.injector.RestaurantService.DeleteRestaurant(ctx, c.Param("restaurant_id"))
	if err != nil {
		responsex.Err(c, err)
		return
	}

	responsex.OK(c, nil)
}
