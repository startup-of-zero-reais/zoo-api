package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type CreateSpecies struct {
	CommonName     string `form:"common_name" json:"common_name"`
	ScientificName string `form:"scientific_name" json:"scientific_name"`
	Kind           string `form:"kind" json:"kind"`
	TaxonomicOrder string `form:"taxonomic_order" json:"taxonomic_order"`
}

func (r *CreateSpecies) Authorize(ctx http.Context) error {
	return nil
}

func (r *CreateSpecies) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"common_name":     "required|min_len:3",
		"scientific_name": "required|min_len:3",
		"kind":            "required|min_len:3",
		"taxonomic_order": "required|min_len:3",
	}
}

func (r *CreateSpecies) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"common_name.required":     "species should have at least common_name",
		"scientific_name.required": "species should have at least scientific_name",
		"kind.required":            "species should have at least kind",
		"taxonomic_order.required": "species should have at least order",
	}
}

func (r *CreateSpecies) Attributes(ctx http.Context) map[string]string {
	return map[string]string{
		"common_name":     "enclosure common name",
		"scientific_name": "enclosure scientific name",
	}
}

func (r *CreateSpecies) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
