package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
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

func (r *WeightController) Show(ctx http.Context) http.Response {
	var sw requests.SearchWeight

	err := ctx.Request().BindQuery(&sw)
	if err != nil {
		facades.Log().Errorf("failed to bind weight by animals filters: %v", err)
		ctx.Request().AbortWithStatusJson(http.StatusInternalServerError, http.Json{"error": err.Error()})
		return nil
	}

	total, weights, err := r.Weight.ListByAnimals(sw, ctx.Request().Route("id"))
	if err != nil {
		facades.Log().Errorf("failed to list weights by animals: %v", err)
		ctx.Request().AbortWithStatusJson(http.StatusInternalServerError, http.Json{"error": err.Error()})
		return nil
	}

	return ctx.Response().Success().Json(http.Json{
		"total":   total,
		"weights": weights,
	})

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
