package auth

import (
	"os"
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
	}

	var user = models.Users{
		Login:    payload.Login,
		Password: payload.Password,
	}

	result := db.Where("login", user.Login).Find(&user)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON("user not found")
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON("db error")
	}

	checkPassHash := utilities.CheckPasswordHash(payload.Password, user.Password)

	if !checkPassHash {
		return c.Status(fiber.StatusUnauthorized).JSON("wrong password")
	}

	refreshSecret := os.Getenv("SECRET_REFRESH-")
	accessSecrer := os.Getenv("SECRET_ACCESS")
	refreshToken := utilities.CreateToken(user.ID, 240*time.Hour, refreshSecret)
	accessToken := utilities.CreateToken(user.ID, 15*time.Minute, accessSecrer)

	var res = Response{
		user.ID,
		user.Login,
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
