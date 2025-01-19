package enclosure

import (
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

type Enclosure interface {
	Create(requests.CreateEnclosure) (models.Enclosure, error)

	List(identification string) (int64, []models.Enclosure, error)
}

type enclosureImpl struct{}

var _ Enclosure = (*enclosureImpl)(nil)

func NewEnclosureService() *enclosureImpl {
	return &enclosureImpl{}
}
