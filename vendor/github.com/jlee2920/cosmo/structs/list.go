package structs

type ListModel struct {
	Skip int `json:"skip"`
	Limit int `json:"limit"`
	Resources []string `json:"resources"`
}