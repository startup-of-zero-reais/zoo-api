package importation

import (
	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/helpers"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/http/responses"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (i importationImpl) UpdateSpecies(re requests.UpdateImportSpecies, ID string) error {
	var sp models.ImportSpecies

	missingFields := validateRequiredSpeciesFields(re)
	if len(missingFields) > 0 {
		if len(missingFields) == 1 {
			sp.Reason = "O campo " + missingFields[0] + " é obrigatório."
		} else {
			sp.Reason = "Os campos " + helpers.JoinWithAnd(missingFields) + " são obrigatórios."
		}
	} else {
		sp.Reason = ""
	}

	sp.CommonName = re.CommonName
	sp.ScientificName = re.ScientificName
	sp.Kind = re.Kind
	sp.Order = re.TaxonomicOrder

	result, err := facades.
		Orm().
		Query().
		Where("id", ID).
		Select("reason", "common_name", "scientific_name", "kind", "taxonomic_order").
		Update(&sp)
	if err != nil {
		facades.Log().Errorf("Falha ao atualizar species: %v", err)
		return responses.ErrUnhandledPgError
	}

	if result.RowsAffected == 0 {
		facades.Log().Infof("ID %v for import species not found", ID)
		return nil
	}

	return nil
}

func validateRequiredSpeciesFields(re requests.UpdateImportSpecies) []string {
	var missing []string

	if re.CommonName == "" {
		missing = append(missing, "Nome comum")
	}
	if re.ScientificName == "" {
		missing = append(missing, "Nome científico")
	}
	if re.Kind == "" {
		missing = append(missing, "Tipo")
	}
	if re.TaxonomicOrder == "" {
		missing = append(missing, "Ordem taxonômica")
	}

	return missing
}
