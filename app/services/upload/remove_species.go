package upload

import (
	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (e *uploadImpl) RemoveSpecies(ids []string) error {
	result, err := facades.Orm().Query().Where("id IN ?", ids).Delete(&models.ImportSpecies{})
	if err != nil {
		facades.Log().Errorf("failed to get import species %v", err)
		return err
	}

	if result.RowsAffected == 0 {
		facades.Log().Warningf("no rows affected on ids %v", ids)
	}

	return nil
}
