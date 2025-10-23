package request

type RegisterDto struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	DateOfBirth string `json:"dateOfBirth"`
	GenderId    int    `json:"genderId"`
}
