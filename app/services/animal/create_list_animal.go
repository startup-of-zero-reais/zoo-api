package animal

import (
	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/http/responses"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (e animalImpl) CreateListAnimal(animals []models.Animal) error {
	err := facades.Orm().Query().Create(&animals)
	if err != nil {
		facades.Log().Errorf("failed to create list animals %v", err)
		return responses.ErrUnhandledPgError
	}

	return nil
}
