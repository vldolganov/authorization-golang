package utilities

import (
	"time"

	"github.com/pascaldekloe/jwt"

	"encoding/json"
)

func CreateToken(userId uint, lifeTime time.Duration, secret string) string {

	var claims jwt.Claims
	claims.Issued = jwt.NewNumericTime(time.Now().Round(time.Second))
	claims.Set = map[string]interface{}{"id": userId}
	claims.Expires = jwt.NewNumericTime(time.Now().Add(lifeTime))
	var extraString = ExtraString{
		"HS256",
		"JWT",
	}

	jsonExtra, jsonErr := json.Marshal(extraString)

	if jsonErr != nil {
		return "Extra string error"
	}
	token, err := claims.HMACSign(jwt.HS256, []byte(secret), jsonExtra)

	if err != nil {
		return "Token error"
	}
	
	return string(token)
}
