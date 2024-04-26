package serializers

import (
	"github.com/Conty111/CarsCatalog/internal/models"
)

type UserInfo struct {
	ID         string `json:"ID"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}

type CarInfo struct {
	ID     string   `json:"ID"`
	RegNum string   `json:"regNum"`
	Mark   string   `json:"mark"`
	Model  string   `json:"model"`
	Year   int      `json:"year"`
	Owner  UserInfo `json:"owner,omitempty"`
}

func SerializeCarInfo(c *models.Car) *CarInfo {
	if c == nil {
		return nil
	}

	carInfo := &CarInfo{
		ID:     c.ID.String(),
		RegNum: c.RegNum,
		Model:  c.Model,
		Mark:   c.Mark,
		Year:   int(c.Year),
		Owner: UserInfo{
			ID:         c.OwnerID.String(),
			Name:       c.Owner.Name,
			Surname:    c.Owner.Surname,
			Patronymic: *c.Owner.Patronymic,
		},
	}

	return carInfo
}
