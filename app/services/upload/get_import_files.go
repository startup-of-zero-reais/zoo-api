package upload

import (
	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

// GetImportFiles implements Upload.
func (e *uploadImpl) GetImportFiles(stateID string) ([]models.ImportEnclosure, []models.ImportSpecies, []models.ImportAnimals, error) {
	var enclosures []models.ImportEnclosure
	var species []models.ImportSpecies
	var animals []models.ImportAnimals

	query := facades.Orm().Query()

	err := query.Where("state_id = ?", stateID).Find(&enclosures)
	if err != nil {
		return nil, nil, nil, err
	}

	err = query.Where("state_id = ?", stateID).Find(&species)
	if err != nil {
		return nil, nil, nil, err
	}

	err = query.Where("state_id = ?", stateID).Find(&animals)
	if err != nil {
		return nil, nil, nil, err
	}

	return enclosures, species, animals, nil
}
