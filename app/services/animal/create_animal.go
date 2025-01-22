package animal

import (
	"time"

	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/http/responses"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (e animalImpl) Create(ac requests.CreateAnimal) (models.Animal, error) {
	var animal models.Animal

	landingAt, _ := time.Parse(time.RFC3339, ac.LandingAt)

	animal.EnclosureID = ac.EnclosureID
	animal.SpeciesID = ac.SpeciesID
	animal.Name = ac.Name
	animal.WasherCode = ac.WasherCode
	animal.MicrochipCode = ac.MicrochipCode
	animal.LandingAt = landingAt
	animal.BornDate = ac.BornDate
	animal.Age = ac.Age
	animal.Origin = ac.Origin
	animal.Gender = ac.Gender

	err := facades.Orm().Query().Create(&animal)
	if err != nil {
		facades.Log().Errorf("failed to create enclosure %v", err)
		return models.Animal{}, responses.ErrUnhandledPgError
	}

	return animal, nil
}
