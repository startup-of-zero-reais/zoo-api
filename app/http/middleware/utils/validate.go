package utils

import (
	"github.com/goravel/framework/contracts/http"
)

func Validate(input http.FormRequest) http.Middleware {
	return func(ctx http.Context) {
		errors, err := ctx.Request().ValidateRequest(input)
		if err != nil {
			ctx.Request().AbortWithStatusJson(http.StatusBadRequest, http.Json{
				"error": err.Error(),
			})
		}

		if errors != nil {
			ctx.Request().AbortWithStatusJson(http.StatusBadRequest, errors.All())
		}

		ctx.Request().Next()
	}
}
