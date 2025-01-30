package upload

import (
	"errors"
	"fmt"

	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/helpers"
	"github.com/startup-of-zero-reais/zoo-api/app/http/responses"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (e *uploadImpl) GetImportEnclosures(ids []string) ([]models.Enclosure, error) {
	var importEnclosures []models.ImportEnclosure

	err := facades.Orm().Query().Where("id IN ?", ids).Find(&importEnclosures)
	if err != nil {
		facades.Log().Errorf("failed to get import enclosures %v", err)
		return nil, responses.ErrUnhandledPgError
	}

	missingIDs := helpers.FindMissingIDs(ids, importEnclosures, func(ie models.ImportEnclosure) string {
		return ie.ID
	})

	if len(missingIDs) > 0 {
		errMsg := fmt.Sprintf("Some import enclosures were not found: %v", missingIDs)
		facades.Log().Error(errMsg)
		return nil, errors.New(errMsg)
	}

	return convertImportEnclosureToEnclosure(importEnclosures)
}

func convertImportEnclosureToEnclosure(importEnclosures []models.ImportEnclosure) ([]models.Enclosure, error) {
	var enclosures []models.Enclosure

	for _, importEnclosure := range importEnclosures {
		enclosures = append(enclosures, models.Enclosure{
			Identification: importEnclosure.Identification,
		})
	}

	return enclosures, nil
}
