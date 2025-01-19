package controllers

import (
	"fmt"

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
	var req requests.CreateEnclosure

	errors, err := ctx.Request().ValidateRequest(&req)
	if err != nil {
		facades.Log().Errorf("failed to get enclosure by id %v", err)
		ctx.Request().AbortWithStatus(http.StatusInternalServerError)
		return nil
	}

	if errors != nil {
		ctx.Response().Json(http.StatusBadRequest, http.Json{
			"errors": errors,
		})
	}

	fmt.Println("errors", errors)

	fmt.Printf("%+v\n", req)

	enclosure, err := r.EnclosureService.Create(req)
	if err != nil {
		facades.Log().Errorf("failed to get enclosure by id %v", err)
		ctx.Request().AbortWithStatus(http.StatusInternalServerError)
		return nil
	}

	return ctx.Response().Success().Json(http.Json{
		"data": enclosure,
	})
}
