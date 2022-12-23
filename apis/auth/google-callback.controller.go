package auth

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"os"

	"authorizationGolang/config"
)

func GoogleCallback(c *fiber.Ctx) error {
	if c.FormValue("state") != "random" {
		return c.Status(fiber.StatusBadRequest).SendString("state is not valid")
	}
	token, err := config.SetupConfig().Exchange(oauth2.NoContext, c.FormValue("code"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("state is not valid")
	}

	resp := c.Get(os.Getenv("GOOGLE_APIS") + token.AccessToken)

	return c.Status(fiber.StatusOK).JSON(resp)
}
