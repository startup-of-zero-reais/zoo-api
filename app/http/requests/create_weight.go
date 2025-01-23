package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type CreateWeight struct {
	Weight   float64 `form:"weight" json:"weight"`
	AnimalId string  `form:"animal_id" json:"animal_id"`
}

func (r *CreateWeight) Authorize(ctx http.Context) error { return nil }

func (r *CreateWeight) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"weight": "required",
	}
}

func (r *CreateWeight) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"weight.required": "Weight history should have at least weight",
	}
}

func (r *CreateWeight) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateWeight) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
