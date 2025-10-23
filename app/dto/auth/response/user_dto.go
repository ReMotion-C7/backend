package response

type UserDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Role int    `json:"role"`
}
