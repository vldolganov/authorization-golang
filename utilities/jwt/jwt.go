package jwt

import (
	"encoding/json"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/pascaldekloe/jwt"
)

func CreateToken(userId uint, lifeTime time.Duration) (string, error) {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}

	var claims jwt.Claims
	claims.Issued = jwt.NewNumericTime(time.Now().Round(time.Second))
	claims.Set = map[string]interface{}{"id": userId}
	claims.Expires = jwt.NewNumericTime(time.Now().Add(lifeTime))
	var extraString = ExtraString{
		"HS256",
		"JWT",
	}

	jsonExtra, _ := json.Marshal(extraString)
	token, err := claims.HMACSign(jwt.HS256, []byte(myEnv["SECRET_WORD"]), jsonExtra)
	return string(token), nil
}

//func CheckValid(token string )
