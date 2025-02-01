package upload

import (
	"errors"
	"fmt"

	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/helpers"
	"github.com/startup-of-zero-reais/zoo-api/app/http/responses"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (e *uploadImpl) GetImportSpecies(ids []string) ([]models.Species, error) {
	var importSpecies []models.ImportSpecies

	err := facades.Orm().Query().Where("id IN ?", ids).Find(&importSpecies)
	if err != nil {
		facades.Log().Errorf("failed to get import species %v", err)
		return nil, responses.ErrUnhandledPgError
	}

	filteredIDs := helpers.Filter(importSpecies, func(i int, is models.ImportSpecies) bool {
		for _, id := range ids {
			if id == is.ID {
				return false
			}
		}

		return true
	})

	if len(filteredIDs) == 0 {
		errMsg := fmt.Sprintf("No reported species IDs found: %v", ids)
		facades.Log().Error(errMsg)
		return nil, errors.New(errMsg)
	}

	return convertImportSpeciesToSpecies(importSpecies)
}

func convertImportSpeciesToSpecies(importSpecies []models.ImportSpecies) ([]models.Species, error) {
	var species []models.Species

	for _, is := range importSpecies {
		species = append(species, models.Species{
			CommonName:     is.CommonName,
			ScientificName: is.ScientificName,
			Kind:           is.Kind,
			TaxonomicOrder: is.Order,
		})
	}

	return species, nil
}
