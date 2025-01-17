package providers

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthServiceProvider struct{}

func (receiver *AuthServiceProvider) Register(app foundation.Application) {
	app.Singleton("google_provider", func(app foundation.Application) (any, error) {
		return &oauth2.Config{
			ClientID:     facades.Config().GetString("auth.oauth.google.client_id"),
			ClientSecret: facades.Config().GetString("auth.oauth.google.secret_key"),
			RedirectURL:  facades.Config().GetString("auth.oauth.google.callback_url"),
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.profile",
				"https://www.googleapis.com/auth/userinfo.email",
				"openid",
			},
			Endpoint: google.Endpoint,
		}, nil
	})
}

func (receiver *AuthServiceProvider) Boot(app foundation.Application) {
}
