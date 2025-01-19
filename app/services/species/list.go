package species

import (
	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/http/responses"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (s *SpeciesImpl) List(se requests.SearchSpecies) (int64, []models.Species, error) {
	var species []models.Species

	query := facades.Orm().Query()

	if se.Search != "" {
		query = query.Where(`search_vector @@ plainto_tsquery('portuguese',?)`, se.Search)
	}

	var total int64

	err := query.Table("species").Count(&total)
	if err != nil {
		facades.Log().Errorf("failed to count species %v", err)
		return 0, nil, responses.ErrUnhandledPgError
	}

	err = query.Find(&species)
	if err != nil {
		facades.Log().Errorf("failed to list species :%v", err)
		return 0, nil, responses.ErrUnhandledPgError
	}

	return total, species, nil
}
