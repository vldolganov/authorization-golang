package auth

import (
	"authorizationGolang/database/models"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"

	"authorizationGolang/config"
	"authorizationGolang/database"
)

func GoogleCallback(c *fiber.Ctx) error {

	var db = database.Connection
	if c.Query("state") != "random" {
		return c.Status(fiber.StatusBadRequest).SendString("State is not valid")
	}

	token, err := config.SetupConfig().Exchange(oauth2.NoContext, c.Query("code"))
	if err != nil {
		return err
	}

	url := "https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken

	resp, err := http.Get(url)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON("user not found")
	}

	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)

	var response GooglePayload
	json.Unmarshal(content, &response)

	var googleRes = Response{
		Login:       response.Email,
		AccessToken: token.AccessToken,
	}

	var user = models.Users{
		Login: response.Email,
	}

	result := db.Create(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON("db error")
	} else if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON("User already exist")
	}

	c.Redirect("/")
	return c.Status(fiber.StatusOK).JSON(googleRes)
}
