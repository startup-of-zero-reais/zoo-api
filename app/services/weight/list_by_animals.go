package weight

import (
	"strings"
	"unicode"

	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/http/responses"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (s *weightImpl) ListByAnimals(sw requests.SearchWeight, animalID string) (int64, []models.Weight, error) {
	var weights []models.Weight
	var total int64

	query := facades.Orm().Query().Where("animal_id", animalID)

	if sw.WeightsFrom != "" {
		query = query.Where("updated_at >= ?", sw.WeightsFrom)
	}

	if sw.WeightsUntil != "" {
		query = query.Where("updated_at <= ?", sw.WeightsUntil)
	}

	err := query.Table("weight_history").Count(&total)
	if err != nil {
		facades.Log().Errorf("failed to count weights by animals %v", err)
		return 0, nil, responses.ErrUnhandledPgError
	}

	if sw.Rel != "" {
		relations := strings.Split(sw.Rel, ",")
		for _, rel := range relations {
			query = query.With(toPascalCase(rel))
		}
	}

	err = query.OrderByDesc("updated_at").Find(&weights)

	if err != nil {
		facades.Log().Errorf("failed to list weights by animals :%v", err)
		return 0, nil, responses.ErrUnhandledPgError
	}

	return total, weights, nil
}

func toPascalCase(input string) string {
	parts := strings.Split(input, ".")
	for i, part := range parts {
		if len(part) > 0 {
			runes := []rune(part)
			runes[0] = unicode.ToUpper(runes[0])
			parts[i] = string(runes)
		}
	}
	return strings.Join(parts, ".")
}
