package config

import (
	"strings"

	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config()

	allowedOrigins := strings.Split(
		config.Env("ALLOWED_ORIGINS", "http://localhost:3000,http://localhost:8080").(string),
		",",
	)

	config.Add("cors", map[string]any{
		// Cross-Origin Resource Sharing (CORS) Configuration
		//
		// Here you may configure your settings for cross-origin resource sharing
		// or "CORS". This determines what cross-origin operations may execute
		// in web browsers. You are free to adjust these settings as needed.
		//
		// To learn more: https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
		"paths":                []string{"*"},
		"allowed_methods":      []string{"*"},
		"allowed_origins":      allowedOrigins,
		"allowed_headers":      []string{"*"},
		"exposed_headers":      []string{""},
		"max_age":              0,
		"supports_credentials": true,
	})
}
