package weight

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/startup-of-zero-reais/zoo-api/app/http/middleware/utils"
	"github.com/startup-of-zero-reais/zoo-api/app/http/requests"
)

func Validate() http.Middleware {
	var createWeight requests.CreateWeight

	return utils.Validate(&createWeight)
}
