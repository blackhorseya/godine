package restaurants

import (
	"github.com/blackhorseya/godine/adapter/restaurant/restful/v1/restaurants/menu"
	"github.com/blackhorseya/godine/adapter/restaurant/wirex"
	_ "github.com/blackhorseya/godine/entity/restaurant/model" // swagger docs
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

		menu.Handle(group, injector)
	}
}

// GetList is used to get the restaurant list.
// @Summary Get the restaurant list.
// @Description Get the restaurant list.
// @Tags restaurants
// @Accept json
// @Produce json
// @Success 200 {object} responsex.Response{data=[]model.Restaurant}
// @Failure 500 {object} responsex.Response
// @Header 200 {int} X-Total-Count "total count"
// @Router /v1/restaurants [get]
func (i *impl) GetList(c *gin.Context) {
	// todo: 2024/6/12|sean|implement me
	panic("implement me")
}

// PostPayload is the post payload.
type PostPayload struct {
	Name        string `json:"name" binding:"required"`
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
	// todo: 2024/6/12|sean|implement me
	panic("implement me")
}
