package authx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/blackhorseya/godine/entity/domain/user/model"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/blackhorseya/godine/pkg/errorx"
	"github.com/blackhorseya/godine/pkg/responsex"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

// Options is a struct that represents the options.
type Options struct {
	Domain       string   `json:"domain" yaml:"domain"`
	ClientID     string   `json:"client_id" yaml:"clientID"`
	ClientSecret string   `json:"client_secret" yaml:"clientSecret"`
	CallbackURL  string   `json:"callback_url" yaml:"callbackURL"`
	Audiences    []string `json:"audiences" yaml:"audiences"`
}

// Authx is a struct that represents the authx.
type Authx struct {
	*oidc.Provider
	oauth2.Config
	middleware *jwtmiddleware.JWTMiddleware
}

// New returns a new Authx.
func New(options Options) (*Authx, error) {
	issuerURL, err := url.Parse("https://" + options.Domain + "/")
	if err != nil {
		return nil, err
	}

	provider, err := oidc.NewProvider(contextx.Background(), issuerURL.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get provider: %w", err)
	}

	config := oauth2.Config{
		ClientID:     options.ClientID,
		ClientSecret: options.ClientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  options.CallbackURL,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	// create middleware
	jwksProvider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)
	jwtValidator, err := validator.New(
		jwksProvider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		options.Audiences,
		validator.WithCustomClaims(func() validator.CustomClaims {
			return &CustomClaims{}
		}),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		return nil, err
	}

	return &Authx{
		Provider: provider,
		Config:   config,
		middleware: jwtmiddleware.New(
			jwtValidator.ValidateToken,
			jwtmiddleware.WithErrorHandler(func(w http.ResponseWriter, r *http.Request, err error) {
				contextx.Background().Error("error validating token", zap.Error(err))
			}),
		),
	}, nil
}

// VerifyIDToken is a method to verify the id token.
func (a *Authx) VerifyIDToken(ctx contextx.Contextx, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

// CustomClaims is the custom claims.
type CustomClaims struct {
	Email string `json:"email,omitempty"`
}

func (c *CustomClaims) Validate(_ context.Context) error {
	return nil
}

// ParseJWT is used to parse the jwt.
func (a *Authx) ParseJWT() gin.HandlerFunc {
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
			customClaims, ok := claims.CustomClaims.(*CustomClaims)
			if !ok {
				ctx.Error("custom claims is not valid")
				return
			}

			by := &model.User{
				ID:    claims.RegisteredClaims.Subject,
				Email: customClaims.Email,
				Roles: nil,
			}
			c.Set(contextx.KeyCtx, contextx.WithValue(ctx, contextx.KeyHandler, by))

			// continue to the next middleware
			c.Next()
		}

		a.middleware.CheckJWT(handler).ServeHTTP(c.Writer, c.Request)

		if encounteredError {
			responsex.Err(c, errorx.Wrap(http.StatusUnauthorized, 401, errors.New("unauthorized")))
			c.Abort()
			return
		}
	}
}
