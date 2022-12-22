package auth

import (
	"authorizationGolang/config"
	"github.com/gofiber/fiber/v2"
)

func GoogleLogin(c *fiber.Ctx) {
	googleConfig := config.SetupConfig()
	url := googleConfig.AuthCodeURL("randomstate")

	c.Redirect(url)
}
