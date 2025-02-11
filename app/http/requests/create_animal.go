package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type CreateAnimal struct {
	Name          string `form:"name" json:"name"`
	WasherCode    string `form:"washer_code" json:"washer_code"`
	MicrochipCode string `form:"microchip_code" json:"microchip_code"`
	LandingAt     string `form:"landing_at" json:"landing_at"`
	Origin        string `form:"origin" json:"origin"`
	BornDate      string `form:"born_date" json:"born_date"`
	Age           string `form:"age" json:"age"`
	Observation   string `form:"observation" json:"observation"`
	Gender        string `form:"gender" json:"gender"`
	SpeciesID     string `form:"species_id" json:"species_id"`
	EnclosureID   string `form:"enclosure_id" json:"enclosure_id"`
}

func (r *CreateAnimal) Authorize(ctx http.Context) error { return nil }

func (r *CreateAnimal) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"landing_at":   "required|date",
		"origin":       "required",
		"age":          "required_without:born_date|in:neonate,cub,young,adult,senile",
		"gender":       "in:male,female,undefined",
		"species_id":   "required|uuid",
		"enclosure_id": "required|uuid",
		"born_date":    "required_without:age|date",
	}
}

func (r *CreateAnimal) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"age.in":                "The age must be one of 'neonate', 'cub', 'young', 'adult' or 'senile'",
		"gender.in":             "The gender must be one of male, female or undefined",
		"landing_at.required":   "The landing at field is required",
		"landing_at.date":       "The landing at field is should be an date string (RFC3339)",
		"born_date.date":        "The born date at field is should be an date string (RFC3339)",
		"origin.required":       "The origin field is required",
		"species_id.required":   "The species id field is required",
		"enclosure_id.required": "The enclosure id field is required",
		"species_id.uuid":       "Species identifier should be valid uuid",
		"enclosure_id.uuid":     "Enclosure identifier should be valid uuid",
	}
}

func (r *CreateAnimal) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateAnimal) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
