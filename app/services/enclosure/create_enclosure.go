package enclosure

import (
	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/http/responses"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (e enclosureImpl) Create(er requests.CreateEnclosure) (models.Enclosure, error) {
	var enclosure models.Enclosure

	enclosure.Identification = er.Identification

	err := facades.Orm().Query().Create(&enclosure)
	if err != nil {
		facades.Log().Errorf("failed to create enclosure %v", err)
		return models.Enclosure{}, responses.ErrUnhandledPgError
	}

	return enclosure, nil
}
