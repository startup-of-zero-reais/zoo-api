package user

import (
	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/http/responses"
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
		return user, responses.ErrUnhandledPgError
	}

	if user.ID == "" {
		return user, responses.ErrUserNotFound
	}

	return user, nil
}
