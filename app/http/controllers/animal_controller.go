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

func (r AnimalController) Index(ctx http.Context) http.Response {
	var searchAnimals requests.SearchAnimals
	err := ctx.Request().BindQuery(&searchAnimals)
	if err != nil {
		facades.Log().Errorf("failed to bind animals filters: %v", err)
		ctx.Request().AbortWithStatusJson(http.StatusInternalServerError, http.Json{"error": err.Error()})
		return nil
	}

	total, animals, err := r.Animal.List(searchAnimals)
	if err != nil {
		facades.Log().Errorf("failed to list animals: %v", err)
		ctx.Request().AbortWithStatusJson(http.StatusInternalServerError, http.Json{"error": err.Error()})
		return nil
	}

	return ctx.Response().Success().Json(http.Json{
		"total":   total,
		"animals": animals,
	})
}
