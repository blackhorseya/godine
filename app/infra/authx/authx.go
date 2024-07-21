package authx

import (
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

// Options is a struct that represents the options.
type Options struct {
	Domain       string `json:"domain" yaml:"domain"`
	ClientID     string `json:"client_id" yaml:"clientID"`
	ClientSecret string `json:"client_secret" yaml:"clientSecret"`
	CallbackURL  string `json:"callback_url" yaml:"callbackURL"`
}

// Authx is a struct that represents the authx.
type Authx struct {
	*oidc.Provider
	oauth2.Config
}
