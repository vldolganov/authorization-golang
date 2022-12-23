package auth

import (
	"authorizationGolang/config"
	"github.com/gofiber/fiber/v2"
)

func GoogleLogin(c *fiber.Ctx) error {
	googleConfig := config.SetupConfig()
	url := googleConfig.AuthCodeURL("random")
	return c.Status(fiber.StatusTemporaryRedirect).Redirect(url)
}
