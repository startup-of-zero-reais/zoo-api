package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
	"github.com/startup-of-zero-reais/zoo-api/app/services/importation"
)

type ImportController struct {
	importation.Importation
}

func NewImportController() *ImportController {
	return &ImportController{
		importation.NewImportationService(),
	}
}

func (r *ImportController) UpdateEnclosure(ctx http.Context) http.Response {
	var ie requests.UpdateImportEnclosure

	id := ctx.Request().Route("id")

	err := ctx.Request().Bind(&ie)
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"error": err.Error()})
	}

	err = r.Importation.UpdateEnclosure(ie, id)
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"error": err.Error()})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{"message": "Update Import Enclosure success!!"})
}

func (r *ImportController) UpdateSpecies(ctx http.Context) http.Response {
	var is requests.UpdateImportSpecies

	id := ctx.Request().Route("id")

	err := ctx.Request().Bind(&is)
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"error": err.Error()})
	}

	err = r.Importation.UpdateSpecies(is, id)
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"error": err.Error()})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{"message": "Update Import Species success!!"})
}

func (r *ImportController) UpdateAnimal(ctx http.Context) http.Response {
	var ia requests.UpdateImportAnimal

	id := ctx.Request().Route("id")

	err := ctx.Request().Bind(&ia)
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"error": err.Error()})
	}

	err = r.Importation.UpdateAnimal(ia, id)
	if err != nil {
		return ctx.Response().Status(http.StatusInternalServerError).Json(http.Json{"error": err.Error()})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{"message": "Update Import Animal success!!"})
}
