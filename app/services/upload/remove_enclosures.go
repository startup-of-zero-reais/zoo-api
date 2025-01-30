package upload

import (
	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (e *uploadImpl) RemoveEnclosures(ids []string) error {
	result, err := facades.Orm().Query().Where("id IN ?", ids).Delete(&models.ImportEnclosure{})
	if err != nil {
		facades.Log().Errorf("failed to get import enclosures %v", err)
		return err
	}

	if result.RowsAffected == 0 {
		facades.Log().Warningf("no rows affected on ids %v", ids)
	}

	return nil
}
