package animal

import (
	"strings"

	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/http/responses"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

// List implements Animal.
func (e *animalImpl) List(sa requests.SearchAnimals) (int64, []models.Animal, error) {
	var animals []models.Animal
	var total int64

	query := facades.Orm().Query()

	if sa.Search != "" {
		query = query.Where(`search_vector @@ plainto_tsquery('portuguese', ?)`, sa.Search)
	}

	err := query.Table("animals").Count(&total)
	if err != nil {
		facades.Log().Errorf("failed to count animals: %v", err)
		return 0, nil, responses.ErrUnhandledPgError
	}

	if sa.Rel != "" {
		relations := strings.Split(sa.Rel, ",")
		for _, rel := range relations {
			query = query.With(UcFirst(rel))
		}
	}

	err = query.Find(&animals)
	if err != nil {
		facades.Log().Errorf("failed to list animals: %v", err)
		return 0, nil, responses.ErrUnhandledPgError
	}

	return total, animals, nil
}

func UcFirst(s string) string {
	first := string(s[0])
	rest := s[1:]

	return strings.ToUpper(first) + rest
}
