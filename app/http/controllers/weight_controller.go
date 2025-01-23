package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/models"
	"github.com/startup-of-zero-reais/zoo-api/app/services/weight"
)

type WeightController struct {
	weight.Weight
}

func NewWeightController() *WeightController {
	return &WeightController{
		weight.NewWeightService(),
	}
}

func (r *WeightController) Store(ctx http.Context) http.Response {
	var weight requests.CreateWeight

	err := ctx.Request().Bind(&weight)

	weight.AnimalID = ctx.Request().Route("id")

	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"error": err.Error()})
	}

	user := ctx.Request().Session().Get("user").(models.User)

	result, err := r.Weight.Create(weight, user.ID)

	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"error": err.Error()})
	}

	return ctx.Response().Success().Json(result)
}
