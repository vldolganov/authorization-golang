package auth

import (
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App) {
	api := app.Group("/api")
	UserApi := api.Group("/auth")

	UserApi.Post("/sign-up", SignUp)
	UserApi.Post("/sign-in", SignIn)
	UserApi.Get("/google/login", GoogleLogin)
	UserApi.Get("/google-callback", GoogleCallback)
}
