package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"authorizationGolang/database/models"
)

var Connection *gorm.DB

func InitConnection() {

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("PORT")

	dsn := "host=" + host + " user=" + user +
			" password=" + password + " dbname=" + dbname + " port=" + port

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Users{})
	Connection = db
}
