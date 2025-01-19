package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/services/enclosure"
)

type EnclosureController struct {
	EnclosureService enclosure.Enclosure
}

func NewEnclosureController() *EnclosureController {
	return &EnclosureController{
		EnclosureService: enclosure.NewEnclosureService(),
	}
}

func (r *EnclosureController) Store(ctx http.Context) http.Response {
	var createEnclosure requests.CreateEnclosure

	err := ctx.Request().Bind(&createEnclosure)
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"error": err.Error()})
	}

	enclosure, err := r.EnclosureService.Create(createEnclosure)
	if err != nil {
		facades.Log().Errorf("failed to get enclosure by id %v", err)
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"error": err.Error()})
	}

	return ctx.Response().Success().Json(enclosure)
}
