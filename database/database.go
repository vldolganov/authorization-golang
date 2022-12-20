package database

import (
	"authorizationGolang/database/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Connection *gorm.DB

func InitConnection() {

	var myEnv map[string]string
	myEnv, err := godotenv.Read()

	host := myEnv["HOST"]
	user := myEnv["USER"]
	password := myEnv["PASSWORD"]
	dbname := myEnv["DBNAME"]
	port := myEnv["PORT"]

	dsn := "host=" + host + " user=" + user +
			" password=" + password + " dbname=" + dbname + " port=" + port

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Users{})
	Connection = db
}
