package user

import (
	"time"

	"github.com/goravel/framework/facades"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/http/responses"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
)

func (u userImpl) Create(usr requests.CreateUser) (models.User, error) {
	usrEmailVerifiedAt, _ := time.Parse(time.RFC3339, usr.EmailVerifiedAt)

	var user models.User

	user.Name = usr.Name
	user.Avatar = usr.Avatar
	user.Email = usr.Email
	user.EmailVerifiedAt = usrEmailVerifiedAt
	user.ID = usr.ID

	err := facades.Orm().Query().Create(&user)
	if err != nil {
		return models.User{}, handlePgError(err)
	}

	return user, nil
}

const (
	DUPLICATE_KEY = "23505"
)

func handlePgError(err error) error {
	switch e := err.(type) {
	case *pgconn.PgError:
		facades.Log().
			In("/app/http/services/user/create_user.go").
			Hint("database error").
			Code(e.SQLState()).
			Error(err)

		switch e.SQLState() {
		case DUPLICATE_KEY:
			return responses.ErrUserAlreadyRegistered
		default:
			return responses.ErrUnhandledPgError
		}
	default:
		facades.Log().Hint("database error").Error(err)
		return responses.ErrUnhandledPgError
	}
}
