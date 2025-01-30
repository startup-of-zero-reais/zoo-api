package enclosure

import (
	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/http/responses"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (e enclosureImpl) CreateListEnclosure(encloures []models.Enclosure) error {
	err := facades.Orm().Query().Create(&encloures)
	if err != nil {
		facades.Log().Errorf("failed to create list enclosures %v", err)
		return responses.ErrUnhandledPgError
	}

	return nil
}
