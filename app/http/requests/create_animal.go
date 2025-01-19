package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type CreateAnimal struct {
	Name        string `form:"name" json:"name"`
	MarkType    string `form:"mark_type" json:"mark_type" validate:"required,enum:washer,microchip"`
	MarkNumber  string `form:"mark_number" json:"mark_number"`
	LandingAt   string `form:"landing_at" json:"landing_at"`
	Origin      string `form:"origin" json:"origin"`
	Age         int64  `form:"age" json:"age"`
	SpeciesId   string `form:"species_id" json:"species_id"`
	EnclosureId string `form:"enclosure_id" json:"enclosure_id"`
}

func (r *CreateAnimal) Authorize(ctx http.Context) error { return nil }

func (r *CreateAnimal) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":         "required",
		"mark_type":    "required|in:washer,microchip",
		"mark_number":  "required",
		"landing_at":   "required",
		"origin":       "required",
		"age":          "required",
		"species_id":   "uuid",
		"enclosure_id": "uuid",
	}
}

func (r *CreateAnimal) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"name.required":        "Animals should have at least name",
		"mark_type.required":   "The mark type field is required.",
		"mark_type.in":         "The mark type must be one of 'wahser' or 'microchip'",
		"mark_number.required": "The mark number field is required",
		"landing_at.required":  "The landing at field is required",
		"origin.required":      "The origin field is required",
		"species_id.uuid":      "Species identifier should be valid uuid",
		"enclosure_id.uuid":    "Enclosure identifier should be valid uuid",
	}
}

func (r *CreateAnimal) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateAnimal) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
