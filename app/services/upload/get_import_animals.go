package upload

import (
	"errors"
	"fmt"

	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/helpers"
	"github.com/startup-of-zero-reais/zoo-api/app/http/responses"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (e *uploadImpl) GetImportAnimals(ids []string) ([]models.Animal, error) {
	var importAnimals []models.ImportAnimals

	err := facades.Orm().Query().Where("id IN ?", ids).Find(&importAnimals)
	if err != nil {
		facades.Log().Errorf("failed to get import animals %v", err)
		return nil, responses.ErrUnhandledPgError
	}

	filteredIDs := helpers.Filter(importAnimals, func(i int, ia models.ImportAnimals) bool {
		for _, id := range ids {
			if id == ia.ID {
				return false
			}
		}

		return true
	})

	if len(filteredIDs) == 0 {
		errMsg := fmt.Sprintf("No reported animals IDs found: %v", ids)
		facades.Log().Error(errMsg)
		return nil, errors.New(errMsg)
	}

	return convertImportAnimalsToAnimals(importAnimals)
}

func convertImportAnimalsToAnimals(importAnimals []models.ImportAnimals) ([]models.Animal, error) {
	var animals []models.Animal

	for _, importAnimal := range importAnimals {
		age, err := helpers.GetAge(importAnimal.Age)
		if err != nil {
			facades.Log().Errorf("failed to get correctly age: %v", err)
			return nil, responses.ErrCannotImportAnimalAge
		}

		gender, err := helpers.GetGender(importAnimal.Gender)
		if err != nil {
			facades.Log().Errorf("failed to get correctly gender: %v", err)
			return nil, responses.ErrCannotImportAnimalGender
		}

		animals = append(animals, models.Animal{
			Name:          importAnimal.Name,
			WasherCode:    importAnimal.WasherCode,
			MicrochipCode: importAnimal.MicrochipCode,
			LandingAt:     importAnimal.LandingAt,
			Origin:        importAnimal.Origin,
			Observation:   importAnimal.Observation,
			BornDate:      importAnimal.BornDate,
			Age:           age,
			Gender:        gender,
			SpeciesID:     importAnimal.SpeciesID,
			EnclosureID:   importAnimal.EnclosureID,
		})
	}

	return animals, nil
}
