package weight

import (
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

type Weight interface {
	Create(requests.CreateWeight, string) (models.Weight, error)
}

type weightImpl struct{}

var _ Weight = (*weightImpl)(nil)

func NewWeightService() *weightImpl {
	return &weightImpl{}
}
