package sessions

import (
	"github.com/blackhorseya/godine/adapter/user/wirex"
	_ "github.com/blackhorseya/godine/entity/domain/user/model" // import model
	"github.com/gin-gonic/gin"
)

type impl struct {
	injector *wirex.Injector
}

// Handle is used to handle the v1 restful API.
func Handle(g *gin.RouterGroup, injector *wirex.Injector) {
	instance := &impl{injector: injector}

	group := g.Group("/sessions")
	{
		group.POST("", instance.Post)
	}
}

// PostPayload represents the post session payload.
type PostPayload struct {
	Name string `json:"name"`
}

// Post is used to create a session.
// @Summary Create a session
// @Description create a session
// @Tags sessions
// @Accept json
// @Produce json
// @Security Bearer
// @Param payload body PostPayload true "session payload"
// @Success 200 {object} responsex.Response{data=model.Account}
// @Failure 400 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Router /v1/sessions [post]
func (i *impl) Post(c *gin.Context) {
}
