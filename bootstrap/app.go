package bootstrap

import (
	"github.com/goravel/framework/foundation"

	"github.com/startup-of-zero-reais/zoo-api/config"
)

func Boot() {
	app := foundation.NewApplication()

	// Bootstrap the application
	app.Boot()

	// Bootstrap the config.
	config.Boot()
}
