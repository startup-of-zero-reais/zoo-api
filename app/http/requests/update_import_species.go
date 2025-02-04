package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UpdateImportSpecies struct {
	CommonName     string `form:"common_name" json:"common_name"`
	ScientificName string `form:"scientific_name" json:"scientific_name"`
	Kind           string `form:"kind" json:"kind"`
	TaxonomicOrder string `form:"taxonomic_order" json:"taxonomic_order"`
}

func (r *UpdateImportSpecies) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateImportSpecies) Rules(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateImportSpecies) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateImportSpecies) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateImportSpecies) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
