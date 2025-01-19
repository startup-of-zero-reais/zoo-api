package enclosure

import (
	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/http/responses"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (e enclosureImpl) List(identification string) (int64, []models.Enclosure, error) {
	var enclosures []models.Enclosure
	var total int64

	query := facades.Orm().Query()

	if identification != "" {
		query = query.Where(`search_vector @@ plainto_tsquery('portuguese', ?)`, identification)
	}

	err := query.Table("enclosures").Count(&total)
	if err != nil {
		facades.Log().Errorf("failed to count enclosure %v", err)
		return total, nil, responses.ErrUnhandledPgError
	}

	err = query.Find(&enclosures)
	if err != nil {
		facades.Log().Errorf("failed to list enclosure %v", err)
		return 0, nil, responses.ErrUnhandledPgError
	}

	return total, enclosures, nil
}
