package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"time"
)

func CreateToken(userId uint) (string, error) {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["id"] = userId
	atClaims["exp"] = time.Now().Add(time.Hour * 100).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString(myEnv["SECRET_WORD"])
	if err != nil {
		return "", err
	}
	return token, nil
}
