package importation

import (
	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/http/responses"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (i importationImpl) UpdateEnclosure(re requests.UpdateImportEnclosure, ID string) error {
	var ie models.ImportEnclosure

	if re.Identification == "" {
		ie.Reason = "A identificação do recinto deve ser preenchida."
	} else {
		ie.Reason = ""
	}
	ie.Identification = re.Identification

	_, err := facades.
		Orm().
		Query().
		Where("id", ID).
		Select("reason", "identification").
		Update(&ie)
	if err != nil {
		facades.Log().Errorf("failed to update enclosure %v", err)
		return responses.ErrUnhandledPgError
	}

	return nil
}
