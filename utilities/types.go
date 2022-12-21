package utilities

type ExtraString struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type HashParams struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}
