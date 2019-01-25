package structs

type RolePost struct {
	Rolename string `json:"rolename"`
	Permissions []string `json:"permissions"`
}