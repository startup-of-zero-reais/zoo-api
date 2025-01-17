package config

import (
	"os"

	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config()
	config.Add("auth", map[string]any{
		// Authentication Defaults
		//
		// This option controls the default authentication "guard"
		// reset options for your application. You may change these defaults
		// as required, but they're a perfect start for most applications.
		"defaults": map[string]any{
			"guard": "user",
		},

		// Authentication Guards
		//
		// Next, you may define every authentication guard for your application.
		// Of course, a great default configuration has been defined for you
		// here which uses session storage and the Eloquent user provider.
		//
		// All authentication drivers have a user provider. This defines how the
		// users are actually retrieved out of your database or other storage
		// mechanisms used by this application to persist your user's data.
		//
		// Supported: "jwt"
		"guards": map[string]any{
			"user": map[string]any{
				"driver": "jwt",
			},
		},

		"oauth": map[string]any{
			"google": map[string]any{
				"client_id":    config.Env("GOOGLE_CLIENT_ID", ""),
				"secret_key":   config.Env("GOOGLE_SECRET_KEY", ""),
				"callback_url": config.Env("GOOGLE_REDIRECT_URL", "http://localhost:8080/api/v1/auth/callback"),
			},
		},

		"token_key": getSessionToken(),
	})
}

const (
	SESSION_TOKEN        = "session_token"
	SECURE_SESSION_TOKEN = "__Secure-session_token"
)

func getSessionToken() string {
	if env := os.Getenv("ENVIRONMENT"); env == "production" {
		return SECURE_SESSION_TOKEN
	}

	return SESSION_TOKEN
}
