package upload

import (
	"errors"

	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (e *uploadImpl) RemoveAnimals(ids []string) error {

	result, err := facades.Orm().Query().Where("id IN ?", ids).Delete(&models.ImportAnimals{})
	if err != nil {
		facades.Log().Errorf("failed to get import animals %v", err)
		return err
	}

	if result.RowsAffected == 0 {
		return errors.New("no records deleted")
	}

	return nil
}
