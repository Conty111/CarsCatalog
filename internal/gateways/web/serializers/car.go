package serializers

import (
	"github.com/Conty111/CarsCatalog/internal/gateways/web/controllers/apiv1/user"
	"github.com/Conty111/CarsCatalog/internal/models"
)

type CarInfo struct {
	ID     string        `jsonapi:"ID"`
	RegNum string        `jsonapi:"regNum"`
	Mark   string        `jsonapi:"mark"`
	Model  string        `jsonapi:"model"`
	Year   int           `jsonapi:"year"`
	Owner  user.UserInfo `jsonapi:"owner"`
}

func SerializeCarInfo(carModel *models.Car) *CarInfo {
	return &CarInfo{
		ID:     carModel.ID.String(),
		Model:  carModel.Model,
		Mark:   carModel.Mark,
		RegNum: carModel.RegNum,
		Year:   int(carModel.Year),
		Owner: user.UserInfo{
			ID:         carModel.OwnerID.String(),
			Name:       carModel.Owner.Name,
			Surname:    carModel.Owner.Surname,
			Patronymic: *carModel.Owner.Patronymic,
		},
	}
}
