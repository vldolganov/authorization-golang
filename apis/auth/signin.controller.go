package auth

import "github.com/gofiber/fiber/v2"

func SignIn(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON("sad")
}
