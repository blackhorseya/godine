package menu

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/godine/adapter/restaurant/wirex"
	"github.com/blackhorseya/godine/app/infra/otelx"
	_ "github.com/blackhorseya/godine/entity/restaurant/model" // swagger docs
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"github.com/blackhorseya/godine/pkg/responsex"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type impl struct {
	injector *wirex.Injector
}

// Handle is used to handle the menu restful API.
func Handle(g *gin.RouterGroup, injector *wirex.Injector) {
	i := &impl{injector: injector}

	group := g.Group("/menu")
	{
		group.GET("", i.GetList)
		group.POST("", i.Post)
	}
}

// GetListQuery is the get list query.
type GetListQuery struct {
	Page int `form:"page" default:"1" minimum:"1"`
	Size int `form:"size" default:"10" minimum:"1" maximum:"100"`
}

// GetList is used to get the menu list.
// @Summary Get the menu list.
// @Description Get the menu list.
// @Tags menu
// @Accept json
// @Produce json
// @Param restaurant_id path string true "restaurant id"
// @Param query query GetListQuery false "get list query"
// @Success 200 {object} responsex.Response{data=[]model.MenuItem}
// @Failure 500 {object} responsex.Response
// @Header 200 {int} X-Total-Count "total count"
// @Router /v1/restaurants/{restaurant_id}/menu [get]
func (i *impl) GetList(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	ctx, span := otelx.Span(ctx, "api.menu.get_list")
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

	items, total, err := i.injector.MenuService.ListMenuItems(ctx, restaurantID)
	if err != nil {
		responsex.Err(c, err)
		return
	}

	c.Header("X-Total-Count", strconv.Itoa(total))
	responsex.OK(c, items)
}

// PostPayload is the post payload.
type PostPayload struct {
	Name        string  `json:"name" binding:"required" example:"menu item name"`
	Description string  `json:"description" example:""`
	Price       float64 `json:"price" binding:"required" example:"10"`
}

// Post is used to add a menu item.
// @Summary Add a menu item.
// @Description Add a menu item.
// @Tags menu
// @Accept json
// @Produce json
// @Param restaurant_id path string true "restaurant id"
// @Param payload body PostPayload true "menu item payload"
// @Success 200 {object} responsex.Response{data=model.MenuItem}
// @Failure 500 {object} responsex.Response
// @Router /v1/restaurants/{restaurant_id}/menu [post]
func (i *impl) Post(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	ctx, span := otelx.Span(ctx, "api.menu.post")
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

	item, err := i.injector.MenuService.AddMenuItem(ctx, restaurantID, payload.Name, payload.Description, payload.Price)
	if err != nil {
		responsex.Err(c, err)
		return
	}

	responsex.OK(c, item)
}