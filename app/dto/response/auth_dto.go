package response

type AuthDataDto struct {
	AccessToken string  `json:"accessToken"`
	TokenType   string  `json:"tokenType"`
	ExpiresIn   int     `json:"expiresIn"`
	User        UserDto `json:"user"`
}

type UserDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Role int    `json:"role"`
}
