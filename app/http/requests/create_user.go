package requests

import (
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type CreateUser struct {
	ID              string `form:"id" json:"id"`
	Name            string `form:"name" json:"name"`
	Avatar          string `form:"avatar" json:"avatar"`
	Email           string `form:"email" json:"email"`
	EmailVerifiedAt string `form:"email_verified_at" json:"email_verified_at"`
}

func (r *CreateUser) Authorize(ctx http.Context) error { return nil }

func (r *CreateUser) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"id": "required|min_len:34",
	}
}

func (r *CreateUser) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"id.required": "user should have at least id",
		"id.min_len":  "user should have an valid id",
	}
}

func (r *CreateUser) Attributes(ctx http.Context) map[string]string {
	return map[string]string{
		"id": "user id",
	}
}

func (r *CreateUser) PrepareForValidation(ctx http.Context, data validation.Data) error {
	if emailVerified, exists := data.Get("email_verified_at"); exists {
		_, err := time.Parse(time.RFC3339, emailVerified.(string))
		if err != nil {
			return err
		}
	}

	return nil
}
