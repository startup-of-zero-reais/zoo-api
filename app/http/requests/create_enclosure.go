package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type CreateEnclosure struct {
	Identification string `form:"identification" json:"identification"`
}

func (r *CreateEnclosure) Authorize(ctx http.Context) error { return nil }

func (r *CreateEnclosure) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"identification": "required",
	}
}

func (r *CreateEnclosure) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"identification.required": "enclosure should have at least identification",
	}
}

func (r *CreateEnclosure) Attributes(ctx http.Context) map[string]string {
	return map[string]string{
		"identification": "enclosure identification",
	}
}

func (r *CreateEnclosure) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
