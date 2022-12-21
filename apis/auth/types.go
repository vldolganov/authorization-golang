package auth

import "authorizationGolang/database/models"

type RequestPayload struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Response struct {
	User        models.Users `json:"user"`
	AccessToken string       `json:"access_token"`
}
