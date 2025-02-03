package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UpdateImportEnclosure struct {
	Id             string `form:"id" json:"id"`
	Identification string `form:"identification" json:"identification"`
}

func (r *UpdateImportEnclosure) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateImportEnclosure) Rules(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateImportEnclosure) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateImportEnclosure) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateImportEnclosure) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
