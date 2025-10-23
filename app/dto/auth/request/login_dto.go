package request

type LoginDto struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}
