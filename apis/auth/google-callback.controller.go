package auth

import (
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"

	"authorizationGolang/config"
)

func GoogleCallback(c *fiber.Ctx) error {

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
	return c.Status(fiber.StatusOK).Send(content)
}
