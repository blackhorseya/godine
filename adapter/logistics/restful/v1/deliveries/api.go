package deliveries

import (
	"github.com/blackhorseya/godine/adapter/logistics/wirex"
	_ "github.com/blackhorseya/godine/entity/logistics/biz"   // import swagger docs
	_ "github.com/blackhorseya/godine/entity/logistics/model" // import swagger docs
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
	}
}

// GetList is used to get the list of deliveries
// @Summary Get the list of deliveries
// @Description Get the list of deliveries
// @Tags deliveries
// @Accept json
// @Produce json
// @Param params query biz.ListDeliveriesOptions false "search params"
// @Security Bearer
// @Success 200 {object} responsex.Response{data=[]model.Delivery}
// @Failure 400 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Header 200 {int} X-Total-Count "total count"
// @Router /v1/deliveries [get]
func (i *impl) GetList(c *gin.Context) {
	// todo: 2024/6/25|sean|implement get list
}
