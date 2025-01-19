package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/session/middleware"

	"github.com/startup-of-zero-reais/zoo-api/app/http/controllers"
	"github.com/startup-of-zero-reais/zoo-api/app/http/middleware/enclosure"
	"github.com/startup-of-zero-reais/zoo-api/app/http/middleware/utils"
)

func Api() {
	facades.Route().
		Prefix("/api/v1").
		Middleware(middleware.StartSession()).
		Group(apiv1)
}

func apiv1(base route.Router) {
	authController := controllers.NewAuthController()
	enclosureController := controllers.NewEnclosureController()

	// Groups declaration

	base.Get("/ping", func(ctx http.Context) http.Response {
		return ctx.Response().Success().Json(http.Json{"message": "pong"})
	})

	// authentication routes
	base.
		Prefix("/auth").
		Group(func(auth route.Router) {
			auth.Get("/google", authController.RedirectToGoogle)
			auth.Get("/callback", authController.HandleCallback)
			auth.Post("/logout", authController.Logout)
		})

	// routes who should be auth

	base.
		Middleware(utils.GrantAuth(authController.UserService.GetByID)).
		Group(authRoutes(authController))

	base.Group(enclosureRoutes(enclosureController))
}

func authRoutes(authController *controllers.AuthController) func(route.Router) {
	return func(router route.Router) {
		router.Get("/auth/me", authController.Me)
	}
}

func enclosureRoutes(enclosureController *controllers.EnclosureController) func(route.Router) {
	return func(router route.Router) {
		router.Middleware(enclosure.Validate()).Post("/enclosures", enclosureController.Store)
	}
}
