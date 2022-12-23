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

	result := db.Where("login", user.Login).Find(&user)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON("user not found")
	}

	checkPassHash := utilities.CheckPasswordHash(payload.Password, user.Password)

	if !checkPassHash {
		return c.Status(fiber.StatusUnauthorized).JSON("wrong password")
	}

	refreshSecret := os.Getenv("SECRET_REFRESH-")
	accessSecrer := os.Getenv("SECRET_ACCESS")
	refreshToken, err := utilities.CreateToken(user.ID, 240*time.Hour, refreshSecret)
	accessToken, err := utilities.CreateToken(user.ID, 15*time.Minute, accessSecrer)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("bad req")
	}

	var res = Response{
		user,
		accessToken,
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "refresh_token"
	cookie.Value = refreshToken
	cookie.Expires = time.Now().Add(240 * time.Hour)
	cookie.HTTPOnly = true
	c.Cookie(cookie)

	return c.Status(fiber.StatusCreated).JSON(res)
}
