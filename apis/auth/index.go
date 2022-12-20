package auth

import (
	"github.com/gofiber/fiber/v2"
)

func UserRouter(api fiber.Router) {
	UserApi := api.Group("/auth")

	UserApi.Post("/sign-up", SignUp)
	UserApi.Post("/sign-in", SignIn)
}
