package user

type UserInfo struct {
	ID         string `json:"ID"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}
