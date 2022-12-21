package main

import (
	"log"

	"authorizationGolang/apis/auth"
	"authorizationGolang/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.InitConnection()

	auth.UserRouter(app)

	log.Fatal(app.Listen(":5050"))
}
