package structs

type UsersPost struct {
	Username string `json:"username,required"`
	Password string `json:"password,required"`
}