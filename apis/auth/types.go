package auth

type RequestPayload struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
