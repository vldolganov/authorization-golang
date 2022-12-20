package jwt

type ExtraString struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}
