package menu

import (
	"github.com/blackhorseya/godine/adapter/restaurant/wirex"
	_ "github.com/blackhorseya/godine/entity/restaurant/model" // swagger docs
	"github.com/gin-gonic/gin"
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
	// todo: 2024/6/12|sean|implement me
	panic("implement me")
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
	// todo: 2024/6/12|sean|implement me
	panic("implement me")
}
