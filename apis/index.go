package apis

import (
	"authorizationGolang/apis/auth"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {
	api := app.Group("/api")

	auth.UserRouter(api)
}
