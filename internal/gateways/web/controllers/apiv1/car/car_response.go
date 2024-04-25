package car

import "github.com/Conty111/CarsCatalog/internal/gateways/web/controllers/apiv1/user"

type CarInfo struct {
	ID     string        `jsonapi:"ID"`
	RegNum string        `jsonapi:"regNum"`
	Mark   string        `jsonapi:"mark"`
	Model  string        `jsonapi:"model"`
	Year   int           `jsonapi:"year"`
	Owner  user.UserInfo `jsonapi:"owner"`
}
