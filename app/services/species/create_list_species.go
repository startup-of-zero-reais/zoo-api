package species

import (
	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/http/responses"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (e SpeciesImpl) CreateListSpecies(species []models.Species) error {
	err := facades.Orm().Query().Create(&species)
	if err != nil {
		facades.Log().Errorf("failed to create list species %v", err)
		return responses.ErrUnhandledPgError
	}

	return nil
}
