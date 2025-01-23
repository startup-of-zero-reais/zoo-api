package species

import (
	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/http/responses"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (c *SpeciesImpl) Create(cs requests.CreateSpecies) (models.Species, error) {
	var species models.Species

	species.CommonName = cs.CommonName
	species.ScientificName = cs.ScientificName
	species.Kind = cs.Kind
	species.TaxonomicOrder = cs.TaxonomicOrder

	err := facades.Orm().Query().Create(&species)
	if err != nil {
		facades.Log().Errorf("failed to create species %v", err)
		return models.Species{}, responses.ErrUnhandledPgError
	}

	return species, nil
}
