package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/session/middleware"

	"github.com/startup-of-zero-reais/zoo-api/app/http/controllers"
	"github.com/startup-of-zero-reais/zoo-api/app/http/middleware/animal"
	"github.com/startup-of-zero-reais/zoo-api/app/http/middleware/enclosure"
	"github.com/startup-of-zero-reais/zoo-api/app/http/middleware/species"
	"github.com/startup-of-zero-reais/zoo-api/app/http/middleware/weight"

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
}

func authRoutes(authController *controllers.AuthController) func(route.Router) {
	return func(router route.Router) {
		router.Get("/auth/me", authController.Me)

		router.Group(enclosureRoutes())

		router.Group(speciesRoutes())

		router.Group(animalRoutes())

		router.Group(uploadRoutes())
	}
}

func enclosureRoutes() func(route.Router) {
	enclosureController := controllers.NewEnclosureController()

	return func(router route.Router) {
		router.Middleware(enclosure.Validate()).Post("/enclosures", enclosureController.Store)
		router.Get("/enclosures", enclosureController.Index)
	}
}

func speciesRoutes() func(route.Router) {
	speciesController := controllers.NewSpeciesController()

	return func(router route.Router) {
		router.Middleware(species.Validate()).Post("/species", speciesController.Create)
		router.Get("/species", speciesController.Index)
	}
}

func animalRoutes() func(route.Router) {
	animalController := controllers.NewAnimalController()
	weightController := controllers.NewWeightController()

	return func(router route.Router) {
		router.Middleware(animal.Validate()).Post("/animals", animalController.Store)
		router.Get("/animals", animalController.Index)
		router.Middleware(weight.Validate()).Post("/animals/{id}/weights", weightController.Store)
		router.Get("/animals/{id}/weights", weightController.Show)
	}
}

func uploadRoutes() func(route.Router) {
	uploadController := controllers.NewUploadController()

	return func(router route.Router) {
		router.Post("/upload", uploadController.Upload)
		router.Get("/upload/imports", uploadController.Index)
	}
}
