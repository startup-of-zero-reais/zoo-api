package importation

import (
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
)

type Importation interface {
	UpdateEnclosure(re requests.UpdateImportEnclosure, ID string) error
	UpdateSpecies(re requests.UpdateImportSpecies, ID string) error
}

type importationImpl struct {
}

var _ Importation = (*importationImpl)(nil)

func NewImportationService() *importationImpl {
	return &importationImpl{}
}
