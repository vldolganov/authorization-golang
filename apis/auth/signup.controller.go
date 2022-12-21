package auth

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"authorizationGolang/database"
	"authorizationGolang/database/models"
	"authorizationGolang/utilities/hash"
	"authorizationGolang/utilities/jwt"
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

	hashPassword, _ := hash.HashPassword(payload.Password)

	var user = models.Users{
		Login:    payload.Login,
		Password: hashPassword,
	}

	result := db.Create(&user)

	dbError := result.RowsAffected

	if dbError == 0 {
		return c.Status(fiber.StatusBadRequest).JSON("User already exist")
	}

	refreshToken, err := jwt.CreateToken(user.ID, 240*time.Hour)
	accessToken, err := jwt.CreateToken(user.ID, 240*time.Minute)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("bad req")
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "refresh_token"
	cookie.Value = refreshToken
	cookie.Expires = time.Now().Add(240 * time.Hour)
	c.Cookie(cookie)
	var res = Response{
		user,
		accessToken,
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
