package user

type UserInfo struct {
	ID         string `jsonapi:"ID"`
	Name       string `jsonapi:"name"`
	Surname    string `jsonapi:"surname"`
	Patronymic string `jsonapi:"patronymic"`
}
