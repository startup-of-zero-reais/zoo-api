package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type SearchAnimals struct {
	Search string `form:"search"`
	Rel    string `form:"rel"`
}

func (r *SearchAnimals) Authorize(ctx http.Context) error {
	return nil
}

func (r *SearchAnimals) Rules(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *SearchAnimals) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *SearchAnimals) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *SearchAnimals) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
