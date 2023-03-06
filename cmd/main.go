package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"

	"authorizationGolang/apis/auth"
	"authorizationGolang/database"
)

func main() {

	var slt []string
	slt = append(slt, "as")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	database.InitConnection()

	auth.UserRouter(app)

	log.Fatal(app.Listen(":8080"))
}
