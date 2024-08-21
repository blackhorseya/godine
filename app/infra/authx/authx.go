package authx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/blackhorseya/godine/app/infra/configx"
	"github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"github.com/blackhorseya/godine/pkg/responsex"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

// Authx is a struct that represents the authx.
type Authx struct {
	*oidc.Provider
	oauth2.Config
	*validator.Validator
	middleware *jwtmiddleware.JWTMiddleware

	SkipPaths []string
}

// New returns a new Authx.
func New(app *configx.Application) (*Authx, error) {
	issuerURL, err := url.Parse("https://" + app.Auth0.Domain + "/")
	if err != nil {
		return nil, err
	}

	provider, err := oidc.NewProvider(contextx.Background(), issuerURL.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get provider: %w", err)
	}

	config := oauth2.Config{
		ClientID:     app.Auth0.ClientID,
		ClientSecret: app.Auth0.ClientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  app.Auth0.CallbackURL,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	// create middleware
	jwksProvider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)
	jwtValidator, err := validator.New(
		jwksProvider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		app.Auth0.Audiences,
		validator.WithCustomClaims(func() validator.CustomClaims {
			return &CustomClaims{}
		}),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		return nil, err
	}

	return &Authx{
		Provider:  provider,
		Config:    config,
		Validator: jwtValidator,
		middleware: jwtmiddleware.New(
			jwtValidator.ValidateToken,
			jwtmiddleware.WithErrorHandler(func(w http.ResponseWriter, r *http.Request, err error) {
				contextx.Background().Error("error validating token", zap.Error(err))
			}),
		),
		SkipPaths: []string{"/grpc.health.v1.Health", "/grpc.reflection.v1alpha.ServerReflection"},
	}, nil
}

// VerifyIDToken is a method to verify the id token.
func (x *Authx) VerifyIDToken(ctx contextx.Contextx, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: x.ClientID,
	}

	return x.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

// CustomClaims is the custom claims.
type CustomClaims struct {
	Roles []string `json:"https://seancheng.space/roles"`
}

func (c *CustomClaims) Validate(_ context.Context) error {
	return nil
}

// ParseJWT is used to parse the jwt.
func (x *Authx) ParseJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		encounteredError := true
		var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
			encounteredError = false
			c.Request = r

			ctx, err := contextx.FromGin(c)
			if err != nil {
				_ = c.Error(err)
				return
			}

			claims, ok := c.Request.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
			if !ok {
				ctx.Error("claims is not valid")
				responsex.Err(c, errorx.Wrap(http.StatusUnauthorized, 401, errors.New("claims is not valid")))
				return
			}

			by := &model.Account{
				Id:       claims.RegisteredClaims.Subject,
				Password: "",
				Address:  nil,
				IsActive: false,
				Level:    0,
			}
			c.Set(contextx.KeyCtx, contextx.WithValue(ctx, contextx.KeyHandler, by))

			// continue to the next middleware
			c.Next()
		}

		x.middleware.CheckJWT(handler).ServeHTTP(c.Writer, c.Request)

		if encounteredError {
			responsex.Err(c, errorx.Wrap(http.StatusUnauthorized, 401, errors.New("unauthorized")))
			c.Abort()
			return
		}
	}
}

// ExtractAccountFromToken is used to decode the access token.
func (x *Authx) ExtractAccountFromToken(accessToken string) (*model.Account, error) {
	validateToken, err := x.Validator.ValidateToken(context.Background(), accessToken)
	if err != nil {
		return nil, err
	}

	claims, ok := validateToken.(*validator.ValidatedClaims)
	if !ok {
		return nil, errors.New("claims is not valid")
	}

	return &model.Account{
		Id:          claims.RegisteredClaims.Subject,
		AccessToken: accessToken,
	}, nil
}

// SkipPath is used to skip the path.
func (x *Authx) SkipPath(path string) bool {
	for _, p := range x.SkipPaths {
		if strings.Contains(path, p) {
			return true
		}
	}

	return false
}
