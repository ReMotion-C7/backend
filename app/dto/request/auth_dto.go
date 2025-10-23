package request

type LoginDto struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type RegisterDto struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	DateOfBirth string `json:"dateOfBirth"`
	GenderId    int    `json:"genderId"`
}
