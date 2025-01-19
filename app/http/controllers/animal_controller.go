package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/services/animal"
)

type AnimalController struct {
	animal.Animal
}

func NewAnimalController() *AnimalController {
	return &AnimalController{
		animal.NewAnimalService(),
	}
}

func (r *AnimalController) Store(ctx http.Context) http.Response {
	var createAnimal requests.CreateAnimal

	err := ctx.Request().Bind(&createAnimal)
	if err != nil {
		facades.Log().Errorf("failed to bind %v", err)
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"error": err.Error()})
	}

	animal, err := r.Create(createAnimal)
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"error": err.Error()})
	}

	return ctx.Response().Success().Json(animal)
}
