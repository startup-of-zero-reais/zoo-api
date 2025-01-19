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

func (r *SpeciesController) Index(ctx http.Context) http.Response {
	var se requests.SearchSpecies

	err := ctx.Request().BindQuery(&se)
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"error": err.Error()})
	}

	total, species, err := r.Species.List(se)
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"error": err.Error()})
	}

	return ctx.Response().Success().Json(http.Json{
		"total":   total,
		"species": species,
	})
}
