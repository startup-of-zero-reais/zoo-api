package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/startup-of-zero-reais/zoo-api/app/services/user"
)

type AuthController struct {
	UserService user.User
}

func NewAuthController() *AuthController {
	return &AuthController{
		UserService: user.NewUserService(),
	}
}

func (a AuthController) Me(ctx http.Context) http.Response {
	usr := ctx.Request().Session().Get("user")
	return ctx.Response().Success().Json(usr)
}

func (a AuthController) RedirectToGoogle(ctx http.Context) http.Response {
	return nil
}

func (a AuthController) HandleCallback(ctx http.Context) http.Response {
	return nil
}

func (a AuthController) Logout(ctx http.Context) http.Response {
	return nil
}
