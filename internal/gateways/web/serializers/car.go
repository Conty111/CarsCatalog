package serializers

import (
	"github.com/Conty111/CarsCatalog/internal/gateways/web/controllers/apiv1/user"
	"github.com/Conty111/CarsCatalog/internal/models"
)

type CarInfo struct {
	ID         string `jsonapi:"ID"`
	Attributes struct {
		RegNum string `jsonapi:"regNum"`
		Mark   string `jsonapi:"mark"`
		Model  string `jsonapi:"model"`
		Year   int    `jsonapi:"year"`
	} `jsonapi:"attributes"`
	Relationships struct {
		Owner struct {
			Data user.UserInfo `jsonapi:"data"`
		} `jsonapi:"owner"`
	} `jsonapi:"relationships"`
}

func SerializeCarInfo(carModel *models.Car) *CarInfo {
	if carModel == nil {
		return nil
	}

	carInfo := &CarInfo{
		ID: carModel.ID.String(),
		Relationships: struct {
			Owner struct {
				Data user.UserInfo `jsonapi:"data"`
			} `jsonapi:"owner"`
		}{
			Owner: struct {
				Data user.UserInfo `jsonapi:"data"`
			}{
				Data: user.UserInfo{
					ID:         carModel.Owner.ID.String(),
					Name:       carModel.Owner.Name,
					Surname:    carModel.Owner.Surname,
					Patronymic: *carModel.Owner.Patronymic,
				},
			},
		},
	}

	carInfo.Attributes.RegNum = carModel.RegNum
	carInfo.Attributes.Mark = carModel.Mark
	carInfo.Attributes.Model = carModel.Model
	carInfo.Attributes.Year = int(carModel.Year)

	return carInfo
}
