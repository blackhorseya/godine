package restful

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"net/http"
	"strings"

	v1 "github.com/blackhorseya/godine/adapter/user/restful/v1"
	"github.com/blackhorseya/godine/adapter/user/wirex"
	_ "github.com/blackhorseya/godine/api/user/restful" // swagger docs
	"github.com/blackhorseya/godine/app/infra/transports/httpx"
	"github.com/blackhorseya/godine/pkg/adapterx"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"github.com/blackhorseya/godine/pkg/responsex"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// @title Godine User Restful API
// @version 0.1.0
// @description Godine User Restful API document.
//
// @contact.name Sean Zheng
// @contact.email blackhorseya@gmail.com
// @contact.url https://blog.seancheng.space
//
// @license.name GPL-3.0
// @license.url https://spdx.org/licenses/GPL-3.0-only.html
//
// @BasePath /api
//
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
type impl struct {
	injector *wirex.Injector
	server   *httpx.Server
}

func newRestful(injector *wirex.Injector, server *httpx.Server) adapterx.Restful {
	return &impl{injector: injector, server: server}
}

func (i *impl) Start() error {
	ctx := contextx.Background()

	err := i.InitRouting()
	if err != nil {
		return err
	}

	err = i.server.Start(ctx)
	if err != nil {
		return err
	}

	ctx.Info("start restful server", zap.String("swagger_url", fmt.Sprintf(
		"http://%s/api/docs/index.html",
		strings.ReplaceAll(i.injector.A.HTTP.GetAddr(), "0.0.0.0", "localhost"),
	)))

	return nil
}

func (i *impl) AwaitSignal() error {
	ctx := contextx.Background()
	ctx.Info("receive signal to stop server")

	if err := i.server.Stop(ctx); err != nil {
		ctx.Error("Failed to stop server", zap.Error(err))
		return fmt.Errorf("failed to stop server: %w", err)
	}

	return nil
}

func (i *impl) InitRouting() error {
	router := i.server.Router

	gob.Register(map[string]interface{}{})
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	router.LoadHTMLGlob("web/template/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})
	router.GET("/login", i.login)
	router.GET("/callback", i.callback)
	router.GET("/user", IsAuthenticated, i.user)

	// api
	api := router.Group("/api")
	{
		api.GET("/docs/*any", ginSwagger.WrapHandler(
			swaggerFiles.Handler,
			ginSwagger.InstanceName("user_restful"),
		))
		api.GET("/healthz", i.Healthz)

		v1.Handle(api, i.injector)
	}

	return nil
}

func (i *impl) GetRouter() *gin.Engine {
	return i.server.Router
}

// Healthz is used to check the health of the service.
// @Summary Check the health of the service.
// @Description Check the health of the service.
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Router /healthz [get]
func (i *impl) Healthz(c *gin.Context) {
	responsex.OK(c, nil)
}

func (i *impl) login(c *gin.Context) {
	state, err := generateRandomState()
	if err != nil {
		responsex.Err(c, err)
		return
	}

	session := sessions.Default(c)
	session.Set("state", state)
	err = session.Save()
	if err != nil {
		responsex.Err(c, err)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, i.injector.Authx.AuthCodeURL(state))
}

func (i *impl) callback(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	session := sessions.Default(c)
	if c.Query("state") != session.Get("state") {
		responsex.Err(c, errorx.New(http.StatusBadRequest, 400, "invalid state parameter"))
		return
	}

	token, err := i.injector.Authx.Exchange(ctx, c.Query("code"))
	if err != nil {
		responsex.Err(c, errorx.Wrap(http.StatusUnauthorized, 401, err))
		return
	}

	idToken, err := i.injector.Authx.VerifyIDToken(ctx, token)
	if err != nil {
		responsex.Err(c, err)
		return
	}

	var profile map[string]interface{}
	err = idToken.Claims(&profile)
	if err != nil {
		responsex.Err(c, err)
		return
	}

	session.Set("access_token", token.AccessToken)
	session.Set("profile", profile)
	err = session.Save()
	if err != nil {
		responsex.Err(c, err)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "/user")
}

func (i *impl) user(c *gin.Context) {
	session := sessions.Default(c)
	profile := session.Get("profile")
	accessToken := session.Get("access_token")

	c.HTML(http.StatusOK, "user.html", map[string]interface{}{
		"profile":      profile,
		"access_token": accessToken,
	})
}

// IsAuthenticated is a middleware that checks if
// the user has already been authenticated previously.
func IsAuthenticated(ctx *gin.Context) {
	if sessions.Default(ctx).Get("profile") == nil {
		ctx.Redirect(http.StatusSeeOther, "/")
	} else {
		ctx.Next()
	}
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
