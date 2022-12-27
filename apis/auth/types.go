package auth

type RequestPayload struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Response struct {
	ID          uint   `json:"id"`
	Login       string `json:"login"`
	AccessToken string `json:"access_token"`
}

type GooglePayload struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
