package authx

import (
	"fmt"

	"github.com/blackhorseya/godine/pkg/contextx"
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

// New returns a new Authx.
func New(options Options) (*Authx, error) {
	provider, err := oidc.NewProvider(contextx.Background(), "https://"+options.Domain+"/")
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

	return &Authx{
		Provider: provider,
		Config:   config,
	}, nil
}
