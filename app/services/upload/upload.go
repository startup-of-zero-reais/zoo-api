package upload

import (
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
	"github.com/startup-of-zero-reais/zoo-api/app/services/animal"
	"github.com/startup-of-zero-reais/zoo-api/app/services/enclosure"
	"github.com/startup-of-zero-reais/zoo-api/app/services/species"
)

type Upload interface {
	Process(string, requests.CreateFile)
	GetImportsStatus() ([]models.ImportStatus, error)
}

type uploadImpl struct {
	EnclosureService enclosure.Enclosure
	SpeciesService   species.Species
	AnimalService    animal.Animal
}

var _ Upload = (*uploadImpl)(nil)

func NewUploadService() *uploadImpl {
	return &uploadImpl{
		EnclosureService: enclosure.NewEnclosureService(),
		SpeciesService:   species.NewSpeciesService(),
		AnimalService:    animal.NewAnimalService(),
	}
}
