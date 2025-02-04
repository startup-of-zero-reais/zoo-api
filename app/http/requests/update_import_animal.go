package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type UpdateImportAnimal struct {
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

func (r *UpdateImportAnimal) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateImportAnimal) Rules(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateImportAnimal) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateImportAnimal) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *UpdateImportAnimal) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
