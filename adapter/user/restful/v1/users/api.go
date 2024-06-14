package users

import (
	"github.com/blackhorseya/godine/adapter/user/wirex"
	"github.com/blackhorseya/godine/entity/user/model"
	"github.com/gin-gonic/gin"
)

type impl struct {
	injector *wirex.Injector
}

// Handle is used to handle the users restful API.
func Handle(g *gin.RouterGroup, injector *wirex.Injector) {
	i := &impl{injector: injector}

	group := g.Group("/users")
	{
		group.POST("", i.Post)
	}
}

// PostPayload represents the post user payload.
type PostPayload struct {
	Name     string        `json:"name" binding:"required" example:"guest"`
	Email    string        `json:"email" binding:"required" example:"guest@gmail.com"`
	Password string        `json:"password" binding:"required" example:"guest"`
	Address  model.Address `json:"address"`
}

// Post is used to create a user.
// @Summary Create a user
// @Description create a user
// @Tags users
// @Accept json
// @Produce json
// @Param payload body PostPayload true "user payload"
// @Success 200 {object} responsex.Response{data=model.User}
// @Failure 400 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Router /v1/users [post]
func (i *impl) Post(c *gin.Context) {
	// todo: 2024/6/14|sean|implement the post user logic
	panic("implement me")
}
