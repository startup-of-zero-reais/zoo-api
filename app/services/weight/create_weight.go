package weight

import (
	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/http/responses"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (e weightImpl) Create(wr requests.CreateWeight, userID string) (models.Weight, error) {
	var weight models.Weight
	weight.Weight = wr.Weight
	weight.UserID = userID
	weight.AnimalID = wr.AnimalID

	err := facades.Orm().Query().Create(&weight)
	if err != nil {
		facades.Log().Errorf("failed to create weight history %v", err)
		return models.Weight{}, responses.ErrUnhandledPgError
	}

	return weight, nil
}
