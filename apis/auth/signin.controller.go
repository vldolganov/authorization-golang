package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"authorizationGolang/database"
	"authorizationGolang/database/models"
	"authorizationGolang/utilities/hash"
	"authorizationGolang/utilities/jwt"
)

func SignIn(c *fiber.Ctx) error {

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

	db.Where("login", user.Login).Find(&user)

	checkPassHash := hash.CheckPasswordHash(payload.Password, user.Password)

	if !checkPassHash {
		return c.Status(fiber.StatusUnauthorized).JSON("wrong password")
	}

	token, err := jwt.CreateToken(user.ID, user.Login)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("bad req")
	}

	var res = Response{
		user,
		token,
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
