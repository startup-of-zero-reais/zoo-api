package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/services/species"
)

type SpeciesController struct {
	species.Species
}

func NewSpeciesController() *SpeciesController {
	return &SpeciesController{
		species.NewSpeciesService(),
	}
}

func (r *SpeciesController) Create(ctx http.Context) http.Response {
	var species requests.CreateSpecies

	err := ctx.Request().Bind(&species)
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"error": err.Error()})
	}

	result, err := r.Species.Create(species)
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"error": err.Error()})
	}

	return ctx.Response().Success().Json(result)
}
