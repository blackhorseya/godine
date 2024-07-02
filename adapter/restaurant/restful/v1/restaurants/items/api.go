package items

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/godine/adapter/restaurant/wirex"
	"github.com/blackhorseya/godine/app/infra/otelx"
	"github.com/blackhorseya/godine/entity/domain/restaurant/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"github.com/blackhorseya/godine/pkg/responsex"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type impl struct {
	injector *wirex.Injector
}

// Handle is used to handle the items restful API.
func Handle(g *gin.RouterGroup, injector *wirex.Injector) {
	i := &impl{injector: injector}

	group := g.Group("/items")
	{
		group.GET("", i.GetList)
		group.POST("", i.Post)
		group.GET("/:item_id", i.GetByID)
		group.PUT("/:item_id", i.PutByID)
		group.DELETE("/:item_id", i.DeleteByID)
	}
}

// GetListQuery is the get list query.
type GetListQuery struct {
	Page int `form:"page" default:"1" minimum:"1"`
	Size int `form:"size" default:"10" minimum:"1" maximum:"100"`
}

// GetList is used to get the items list.
// @Summary Get the items list.
// @Description Get the items list.
// @Tags items
// @Accept json
// @Produce json
// @Param restaurant_id path string true "restaurant id"
// @Param query query GetListQuery false "get list query"
// @Success 200 {object} responsex.Response{data=[]model.MenuItem}
// @Failure 500 {object} responsex.Response
// @Header 200 {number} X-Total-Count "total count"
// @Router /v1/restaurants/{restaurant_id}/items [get]
func (i *impl) GetList(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	ctx, span := otelx.Span(ctx, "api.items.get_list")
	defer span.End()

	var query GetListQuery
	err = c.ShouldBindQuery(&query)
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	restaurantID, err := uuid.Parse(c.Param("restaurant_id"))
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	items, total, err := i.injector.MenuService.ListMenuItems(ctx, restaurantID.String())
	if err != nil {
		responsex.Err(c, err)
		return
	}

	c.Header("X-Total-Count", strconv.Itoa(total))
	responsex.OK(c, items)
}

// PostPayload is the post payload.
type PostPayload struct {
	Name        string  `json:"name" binding:"required" example:"items item name"`
	Description string  `json:"description" example:""`
	Price       float64 `json:"price" binding:"required" example:"10"`
}

// Post is used to add a items item.
// @Summary Add a items item.
// @Description Add a items item.
// @Tags items
// @Accept json
// @Produce json
// @Param restaurant_id path string true "restaurant id"
// @Param payload body PostPayload true "items item payload"
// @Success 200 {object} responsex.Response{data=model.MenuItem}
// @Failure 500 {object} responsex.Response
// @Router /v1/restaurants/{restaurant_id}/items [post]
func (i *impl) Post(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	ctx, span := otelx.Span(ctx, "api.items.post")
	defer span.End()

	restaurantID, err := uuid.Parse(c.Param("restaurant_id"))
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	var payload PostPayload
	err = c.ShouldBindJSON(&payload)
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	item, err := i.injector.MenuService.AddMenuItem(
		ctx,
		restaurantID.String(),
		payload.Name,
		payload.Description,
		payload.Price,
	)
	if err != nil {
		responsex.Err(c, err)
		return
	}

	responsex.OK(c, item)
}

// GetByID is used to get the items item by id.
// @Summary Get the items item by id.
// @Description Get the items item by id.
// @Tags items
// @Accept json
// @Produce json
// @Param restaurant_id path string true "restaurant id"
// @Param item_id path string true "item id"
// @Success 200 {object} responsex.Response{data=model.MenuItem}
// @Failure 500 {object} responsex.Response
// @Router /v1/restaurants/{restaurant_id}/items/{item_id} [get]
func (i *impl) GetByID(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	ctx, span := otelx.Span(ctx, "api.items.get_by_id")
	defer span.End()

	restaurantID, err := uuid.Parse(c.Param("restaurant_id"))
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	itemID, err := uuid.Parse(c.Param("item_id"))
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	item, err := i.injector.MenuService.GetMenuItem(ctx, restaurantID.String(), itemID.String())
	if err != nil {
		responsex.Err(c, err)
		return
	}

	responsex.OK(c, item)
}

// PutByID is used to update the items item by id.
// @Summary Update the items item by id.
// @Description Update the items item by id.
// @Tags items
// @Accept json
// @Produce json
// @Param restaurant_id path string true "restaurant id"
// @Param item_id path string true "item id"
// @Param payload body model.MenuItem true "items item payload"
// @Success 200 {object} responsex.Response{data=model.MenuItem}
// @Failure 500 {object} responsex.Response
// @Router /v1/restaurants/{restaurant_id}/items/{item_id} [put]
func (i *impl) PutByID(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	ctx, span := otelx.Span(ctx, "api.items.put_by_id")
	defer span.End()

	restaurantID, err := uuid.Parse(c.Param("restaurant_id"))
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	itemID, err := uuid.Parse(c.Param("item_id"))
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	var payload model.MenuItem
	err = c.ShouldBindJSON(&payload)
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	err = i.injector.MenuService.UpdateMenuItem(
		ctx,
		restaurantID.String(),
		itemID.String(),
		payload.Name,
		payload.Description,
		payload.Price,
		payload.IsAvailable,
	)
	if err != nil {
		responsex.Err(c, err)
		return
	}

	item, err := i.injector.MenuService.GetMenuItem(ctx, restaurantID.String(), itemID.String())
	if err != nil {
		responsex.Err(c, err)
		return
	}

	responsex.OK(c, item)
}

// DeleteByID is used to delete the items item by id.
// @Summary Delete the items item by id.
// @Description Delete the items item by id.
// @Tags items
// @Accept json
// @Produce json
// @Param restaurant_id path string true "restaurant id"
// @Param item_id path string true "item id"
// @Success 204 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Router /v1/restaurants/{restaurant_id}/items/{item_id} [delete]
func (i *impl) DeleteByID(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	ctx, span := otelx.Span(ctx, "api.items.delete_by_id")
	defer span.End()

	restaurantID, err := uuid.Parse(c.Param("restaurant_id"))
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	itemID, err := uuid.Parse(c.Param("item_id"))
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusBadRequest, 400, err))
		return
	}

	err = i.injector.MenuService.RemoveMenuItem(ctx, restaurantID.String(), itemID.String())
	if err != nil {
		responsex.Err(c, err)
		return
	}

	responsex.OK(c, nil)
}
