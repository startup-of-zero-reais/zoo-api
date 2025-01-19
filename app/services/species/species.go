package species

import (
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

type Species interface {
	Create(cs requests.CreateSpecies) (models.Species, error)
}

type SpeciesImpl struct{}

var _ Species = (*SpeciesImpl)(nil)

func NewSpeciesService() *SpeciesImpl {
	return &SpeciesImpl{}
}