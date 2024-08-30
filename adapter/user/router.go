package user

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"net/url"

	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func (i *impl) login(c *gin.Context) {
	state, err := generateRandomState()
	if err != nil {
		_ = c.Error(err)
		return
	}

	session := sessions.Default(c)
	session.Set("state", state)
	err = session.Save()
	if err != nil {
		_ = c.Error(err)
		return
	}

	var options []oauth2.AuthCodeOption
	for _, audience := range i.injector.A.Auth0.Audiences {
		options = append(options, oauth2.SetAuthURLParam("audience", audience))
	}

	c.Redirect(http.StatusTemporaryRedirect, i.injector.Authx.AuthCodeURL(state, options...))
}

func (i *impl) callback(c *gin.Context) {
	ctx := contextx.Background()

	session := sessions.Default(c)
	if c.Query("state") != session.Get("state") {
		_ = c.Error(errorx.New(http.StatusBadRequest, 400, "invalid state"))
		return
	}

	token, err := i.injector.Authx.Exchange(ctx, c.Query("code"))
	if err != nil {
		_ = c.Error(errorx.Wrap(http.StatusUnauthorized, 401, err))
		return
	}

	idToken, err := i.injector.Authx.VerifyIDToken(ctx, token)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var profile map[string]interface{}
	err = idToken.Claims(&profile)
	if err != nil {
		_ = c.Error(err)
		return
	}

	session.Set("access_token", token.AccessToken)
	session.Set("profile", profile)
	err = session.Save()
	if err != nil {
		_ = c.Error(err)
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

func (i *impl) logout(c *gin.Context) {
	logoutURL, err := url.Parse("https://" + i.injector.A.Auth0.Domain + "/v2/logout")
	if err != nil {
		_ = c.Error(err)
		return
	}

	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + c.Request.Host)
	if err != nil {
		_ = c.Error(err)
		return
	}

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", i.injector.Authx.ClientID)
	logoutURL.RawQuery = parameters.Encode()

	c.Redirect(http.StatusTemporaryRedirect, logoutURL.String())
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
