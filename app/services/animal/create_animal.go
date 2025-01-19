package animal

import (
	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/http/responses"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (e animalImpl) Create(ac requests.CreateAnimal) (models.Animal, error) {
	var animal models.Animal

	animal.Name = ac.Name
	animal.MarkType = models.MarkTypeStatus(ac.MarkType)
	animal.MarkNumber = ac.MarkNumber
	animal.LandingAt = ac.LandingAt
	animal.Origin = ac.Origin

	err := facades.Orm().Query().Create(&animal)
	if err != nil {
		facades.Log().Errorf("failed to create enclosure %v", err)
		return models.Animal{}, responses.ErrUnhandledPgError
	}

	return animal, nil
}
