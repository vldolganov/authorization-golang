package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"authorizationGolang/apis/auth"
	"authorizationGolang/database"
)

func main() {

	var slt []string
	slt = append(slt, "as")
	err := godotenv.Load("./cmd/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	app := fiber.New()
	database.InitConnection(config)

	auth.UserRouter(app)

	log.Fatal(app.Listen(":5000"))
}
