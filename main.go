package main

import (
	"authorizationGolang/apis"
	"authorizationGolang/database"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	database.InitConnection()
	apis.InitRouter(app)
	log.Fatal(app.Listen(":5050"))
}
