package auth

import (
	"authorizationGolang/database"
	"authorizationGolang/database/models"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func SignUp(c *fiber.Ctx) error {
	var payload RequestPayload

	var db = database.Connection

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("bad request")
	} else if strings.TrimSpace(payload.Login) == "" &&
		strings.TrimSpace(payload.Password) == "" {
		return c.Status(fiber.StatusBadRequest).JSON("pass login or password")
	}

	var user = models.Users{
		Login:    payload.Login,
		Password: payload.Password,
	}

	db.Create(&user)

	return c.Status(fiber.StatusCreated).JSON(user)
}
