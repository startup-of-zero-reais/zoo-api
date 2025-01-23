package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type SearchWeight struct {
	WeightsFrom  string `form:"weights_from" json:"weights_from"`
	WeightsUntil string `form:"weights_until" json:"weights_until"`
	Rel          string `form:"rel" json:"rel"`
}

func (r *SearchWeight) Authorize(ctx http.Context) error {
	return nil
}

func (r *SearchWeight) Rules(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *SearchWeight) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *SearchWeight) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *SearchWeight) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
