package user

import "github.com/startup-of-zero-reais/zoo-api/app/models"

type User interface {
	GetByID(string) (models.User, error)
}

type userImpl struct{}

var _ User = (*userImpl)(nil)

func NewUserService() *userImpl {
	return &userImpl{}
}
