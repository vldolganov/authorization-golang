package auth

import (
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"authorizationGolang/database"
	"authorizationGolang/database/models"
	"authorizationGolang/utilities"
)

func SignUp(c *fiber.Ctx) error {
	var payload RequestPayload
	var db = database.Connection

	login := strings.TrimSpace(payload.Login)
	password := strings.TrimSpace(payload.Password)

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("bad request")
	} else if login != "" && password != "" {
		return c.Status(fiber.StatusBadRequest).JSON("pass login or password")
	}

	hashPassword := utilities.HashPassword(payload.Password)

	var user = models.Users{
		Login:    payload.Login,
		Password: hashPassword,
	}

	result := db.Create(&user)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON("User already exist")
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON("db error")
	}

	refreshSecret := os.Getenv("SECRET_REFRESH")
	accessSecret := os.Getenv("SECRET_ACCESS")
	refreshToken := utilities.CreateToken(user.ID, 240*time.Hour, refreshSecret)
	accessToken := utilities.CreateToken(user.ID, 15*time.Minute, accessSecret)

	cookie := new(fiber.Cookie)
	cookie.Name = "refresh_token"
	cookie.Value = refreshToken
	cookie.Expires = time.Now().Add(240 * time.Hour)
	c.Cookie(cookie)
	var res = Response{
		user.ID,
		user.Login,
		accessToken,
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
