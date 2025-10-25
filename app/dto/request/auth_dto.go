package request

type LoginDto struct {
	Identifier string `json:"identifier" validate:"required"`
	Password   string `json:"password" validate:"required"`
}

type RegisterDto struct {
	Email       string `json:"email" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Password    string `json:"password" validate:"required,min=8"`
	PhoneNumber string `json:"phoneNumber" validate:"required"` 
	DateOfBirth string `json:"dateOfBirth" validate:"required"`
	GenderId    int    `json:"genderId" validate:"required"`
}
