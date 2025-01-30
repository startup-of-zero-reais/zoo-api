package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type ConfirmUpload struct {
	Type string   `form:"type" json:"type"`
	IDs  []string `form:"ids" json:"ids"`
}

func (r *ConfirmUpload) Authorize(ctx http.Context) error {
	return nil
}

func (r *ConfirmUpload) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"ids":   "required|array",
		"ids.*": "required|uuid",
		"type":  "required|in:animal,species,enclosure",
	}
}

func (r *ConfirmUpload) Messages(ctx http.Context) map[string]string {
	return map[string]string{
		"type.required": "The type field is required.",
		"type.in":       "The type field must be one of: animal, species or enclosure.",
	}
}

func (r *ConfirmUpload) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *ConfirmUpload) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
