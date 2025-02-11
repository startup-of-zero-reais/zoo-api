package upload

import (
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
	"github.com/startup-of-zero-reais/zoo-api/app/services/animal"
	"github.com/startup-of-zero-reais/zoo-api/app/services/enclosure"
	"github.com/startup-of-zero-reais/zoo-api/app/services/species"
)

type Upload interface {
	Process(string, requests.CreateFile, models.ImportStatus)
	GetImportsStatus() ([]models.ImportStatus, error)
	GetImportAnimals(ids []string) ([]models.Animal, error)
	GetImportEnclosures(ids []string) ([]models.Enclosure, error)
	GetImportSpecies(ids []string) ([]models.Species, error)
	RemoveAnimals(ids []string) error
	RemoveEnclosures(ids []string) error
	RemoveSpecies(ids []string) error
	GetImportFiles(stateID string) ([]models.ImportEnclosure, []models.ImportSpecies, []models.ImportAnimals, error)
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
