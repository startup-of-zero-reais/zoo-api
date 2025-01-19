package animal

import (
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

type Animal interface {
	Create(requests.CreateAnimal) (models.Animal, error)
}

type animalImpl struct{}

var _ Animal = (*animalImpl)(nil)

func NewAnimalService() *animalImpl {
	return &animalImpl{}
}
