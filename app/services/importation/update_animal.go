package importation

import (
	"time"

	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/helpers"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/http/responses"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (i importationImpl) UpdateAnimal(ra requests.UpdateImportAnimal, ID string) error {
	missingFields := validateRequiredFields(ra)

	var ia models.ImportAnimals
	if len(missingFields) > 0 {
		if len(missingFields) == 1 {
			ia.Reason = "O campo " + missingFields[0] + " é obrigatório."
		} else {
			ia.Reason = "Os campos " + helpers.JoinWithAnd(missingFields) + " são obrigatórios."
		}
	} else {
		ia.Reason = ""
	}

	landingAt, _ := time.Parse(time.RFC3339, ra.LandingAt)
	bornDate, _ := time.Parse(time.RFC3339, ra.BornDate)

	ia.Name = ra.Name
	ia.WasherCode = ra.WasherCode
	ia.MicrochipCode = ra.MicrochipCode
	ia.LandingAt = landingAt
	ia.Origin = ra.Origin
	ia.BornDate = bornDate
	ia.Age = ra.Age
	ia.Observation = ra.Observation
	ia.Gender = ra.Gender
	ia.SpeciesID = helpers.ToNullableString(ra.SpeciesID)
	ia.EnclosureID = helpers.ToNullableString(ra.EnclosureID)

	result, err := facades.
		Orm().
		Query().
		Where("id", ID).
		Select(
			"reason",
			"name",
			"washer_code",
			"microchip_code",
			"landing_at",
			"origin",
			"born_date",
			"age",
			"observation",
			"gender",
			"species_id",
			"enclosure_id",
		).
		Update(&ia)
	if err != nil {
		facades.Log().Errorf("Error on update animal: %v", err)
		return responses.ErrUnhandledPgError
	}

	if result.RowsAffected == 0 {
		facades.Log().Infof("ID %v for import species not found", ID)
		return nil
	}

	return nil
}

func validateRequiredFields(ra requests.UpdateImportAnimal) []string {
	var missing []string

	if ra.Origin == "" {
		missing = append(missing, "'Origem'")
	}
	if ra.SpeciesID == "" {
		missing = append(missing, "'Espécie'")
	}
	if ra.EnclosureID == "" {
		missing = append(missing, "'Recinto'")
	}
	if ra.BornDate == "" && ra.Age == "" {
		missing = append(missing, "'Data de Nascimento' ou 'Idade'")
	}

	return missing
}
