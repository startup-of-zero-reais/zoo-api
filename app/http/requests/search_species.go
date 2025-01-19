package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type SearchSpecies struct {
	Search string `form:"search"`
}

func (r *SearchSpecies) Authorize(ctx http.Context) error {
	return nil
}

func (r *SearchSpecies) Rules(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *SearchSpecies) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *SearchSpecies) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *SearchSpecies) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
