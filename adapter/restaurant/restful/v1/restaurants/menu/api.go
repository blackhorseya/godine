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
	}
}

// GetList is used to get the menu list.
// @Summary Get the menu list.
// @Description Get the menu list.
// @Tags menu
// @Accept json
// @Produce json
// @Param restaurant_id path string true "restaurant id"
// @Success 200 {object} responsex.Response{data=[]model.MenuItem}
// @Failure 500 {object} responsex.Response
// @Header 200 {int} X-Total-Count "total count"
// @Router /v1/restaurants/{restaurant_id}/menu [get]
func (i *impl) GetList(c *gin.Context) {
	// todo: 2024/6/12|sean|implement me
	panic("implement me")
}
