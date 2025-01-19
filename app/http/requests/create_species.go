package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type CreateSpecies struct {
	CommonName    string `form:"common_name" json:"common_name"`
	CientificName string `form:"cientific_name" json:"cientific_name"`
}

func (r *CreateSpecies) Authorize(ctx http.Context) error {
	return nil
}

func (r *CreateSpecies) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"common_name":    "required|min_len:3",
		"cientific_name": "required|min_len:3",
	}
}

func (r *CreateSpecies) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"common_name.required":    "species should have at least common_name",
		"cientific_name.required": "species should have at least cientific_name",
	}
}

func (r *CreateSpecies) Attributes(ctx http.Context) map[string]string {
	return map[string]string{
		"common_name":    "enclosure common name",
		"cientific_name": "enclosure cientific name",
	}
}

func (r *CreateSpecies) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
