package utils

import "github.com/goravel/framework/contracts/http"

func BindAndValidate(ctx http.Context, input http.FormRequest) bool {
	errors, err := ctx.Request().ValidateRequest(input)
	if err != nil {
		ctx.Request().AbortWithStatusJson(http.StatusBadRequest, http.Json{
			"error": err.Error(),
		})
		return true
	}

	if errors != nil {
		ctx.Request().AbortWithStatusJson(http.StatusBadRequest, errors.All())
		return true
	}

	return false
}
