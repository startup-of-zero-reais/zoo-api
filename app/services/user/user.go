package user

import (
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

type User interface {
	Create(requests.CreateUser) (models.User, error)
	GetByID(string) (models.User, error)
}

type userImpl struct{}

var _ User = (*userImpl)(nil)

func NewUserService() *userImpl {
	return &userImpl{}
}
