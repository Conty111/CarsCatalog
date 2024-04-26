package serializers

import (
	"github.com/Conty111/CarsCatalog/internal/gateways/web/controllers/apiv1/user"
	"github.com/Conty111/CarsCatalog/internal/models"
)

type CarInfo struct {
	ID     string        `jsonapi:"primary,ID"`
	RegNum string        `jsonapi:"attr,regNum"`
	Mark   string        `jsonapi:"attr,mark"`
	Model  string        `jsonapi:"attr,model"`
	Year   int           `jsonapi:"attr,year"`
	Owner  user.UserInfo `jsonapi:"relation,owner,omitempty"`
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
		Owner: user.UserInfo{
			ID:         c.OwnerID.String(),
			Name:       c.Owner.Name,
			Surname:    c.Owner.Surname,
			Patronymic: *c.Owner.Patronymic,
		},
	}

	return carInfo
}
