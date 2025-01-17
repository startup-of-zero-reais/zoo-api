package providers

import (
	"github.com/goravel/framework/contracts/database/seeder"
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"

	"github.com/startup-of-zero-reais/zoo-api/database/seeders"
)

type DatabaseServiceProvider struct{}

func (receiver *DatabaseServiceProvider) Register(app foundation.Application) {
}

func (receiver *DatabaseServiceProvider) Boot(app foundation.Application) {
	facades.Seeder().Register([]seeder.Seeder{
		&seeders.DatabaseSeeder{},
	})
}
