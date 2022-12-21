package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"authorizationGolang/apis/auth"
	"authorizationGolang/database"
)

func main() {
	app := fiber.New()
	database.InitConnection()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	auth.UserRouter(app)

	log.Fatal(app.Listen(":5050"))
}
