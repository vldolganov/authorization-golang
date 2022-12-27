package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"authorizationGolang/apis/auth"
	"authorizationGolang/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	database.InitConnection()

	auth.UserRouter(app)

	log.Fatal(app.Listen(":5000"))
}
