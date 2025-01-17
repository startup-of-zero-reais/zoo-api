package user

import (
	"fmt"

	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

// GetByID implements User.
func (u *userImpl) GetByID(userID string) (models.User, error) {
	var user models.User

	err := facades.Orm().
		Query().
		Where("id", userID).
		First(&user)
	if err != nil {
		facades.Log().
			Hint("failed to get user by id").
			Error(err)
		return user, fmt.Errorf("failed to get user by id")
	}

	if user.ID == "" {
		return user, fmt.Errorf("user not found")
	}

	return user, nil
}
