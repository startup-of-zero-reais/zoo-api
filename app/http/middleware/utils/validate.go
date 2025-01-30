package utils

import (
	"github.com/goravel/framework/contracts/http"
)

func Validate(input http.FormRequest) http.Middleware {
	return func(ctx http.Context) {
		if BindAndValidate(ctx, input) {
			return
		}

		ctx.Request().Next()
	}
}
