package animal

import (
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

type Animal interface {
	CreateListAnimal(animals []models.Animal) error
	Create(requests.CreateAnimal) (models.Animal, error)
	List(requests.SearchAnimals) (int64, []models.Animal, error)
}

type animalImpl struct{}

var _ Animal = (*animalImpl)(nil)

func NewAnimalService() *animalImpl {
	return &animalImpl{}
}
