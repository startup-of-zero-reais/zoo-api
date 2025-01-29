package upload

import (
	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

// GetImportsStatus implements Upload.
func (e *uploadImpl) GetImportsStatus() ([]models.ImportStatus, error) {
	var imports []models.ImportStatus
	err := facades.Orm().Query().Find(&imports)
	return imports, err
}
