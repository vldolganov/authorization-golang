package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"authorizationGolang/database/models"
)

var Connection *gorm.DB

func InitConnection() {

	dsn := "host=localhost user=dolganoffadmin password=dolganoffadmin dbname=users port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Users{})
	Connection = db
}
