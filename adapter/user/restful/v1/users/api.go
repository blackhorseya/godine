package users

import (
	"net/http"

	"github.com/blackhorseya/godine/adapter/user/wirex"
	"github.com/blackhorseya/godine/entity/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"github.com/blackhorseya/godine/pkg/responsex"
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
		group.GET("/:id", i.GetByID)
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

	item, err := i.injector.UserService.CreateUser(ctx, payload.Name, payload.Email, payload.Password, payload.Address)
	if err != nil {
		responsex.Err(c, err)
		return
	}

	responsex.OK(c, item)
}

// GetByID is used to get a user by id.
// @Summary Get a user by id
// @Description get a user by id
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} responsex.Response{data=model.User}
// @Failure 400 {object} responsex.Response
// @Failure 404 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Router /v1/users/{id} [get]
func (i *impl) GetByID(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	id := c.Param("id")

	item, err := i.injector.UserService.GetUser(ctx, id)
	if err != nil {
		responsex.Err(c, err)
		return
	}

	responsex.OK(c, item)
}
