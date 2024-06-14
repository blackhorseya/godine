package items

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

// Handle is used to handle the items restful API.
func Handle(g *gin.RouterGroup, injector *wirex.Injector) {
	i := &impl{injector: injector}

	group := g.Group("/items")
	{
		group.GET("", i.GetList)
		group.POST("", i.Post)
		group.GET("/:item_id", i.GetByID)
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
// @Header 200 {int} X-Total-Count "total count"
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

	item, err := i.injector.MenuService.AddMenuItem(ctx, restaurantID, payload.Name, payload.Description, payload.Price)
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
	// todo: 2024/6/14|sean|implement get item by id
	panic("implement me")
}
