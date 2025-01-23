package weight

import (
	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/http/responses"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (e weightImpl) Create(wr requests.CreateWeight, userId string) (models.Weight, error) {
	var weight_history models.Weight
	weight_history.Weight = wr.Weight
	weight_history.UserId = userId
	weight_history.AnimalId = wr.AnimalId

	err := facades.Orm().Query().Create(&weight_history)
	if err != nil {
		facades.Log().Errorf("failed to create weight history %v", err)
		return models.Weight{}, responses.ErrUnhandledPgError
	}

	return weight_history, nil
}
