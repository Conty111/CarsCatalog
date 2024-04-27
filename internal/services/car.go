package services

import (
	"errors"
	"github.com/Conty111/CarsCatalog/internal/errs"
	"github.com/Conty111/CarsCatalog/internal/external_api"
	"github.com/Conty111/CarsCatalog/internal/gateways/web/controllers/apiv1/car"
	"github.com/Conty111/CarsCatalog/internal/gateways/web/helpers"
	"github.com/Conty111/CarsCatalog/internal/interfaces"
	"github.com/Conty111/CarsCatalog/internal/models"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"regexp"
)

const pattern = "^[ABEKMHOPCTYXАВЕКМНОРСТУХ]{1}\\d{3}[ABEKMHOPCTYXАВЕКМНОРСТУХ]{2}\\d{2,3}$"

type CarService struct {
	CarRepo      interfaces.CarManager
	UserProvider interfaces.UserProvider
	CarAPI       external_api.ExternalAPIClient
	re           *regexp.Regexp
}

func NewCarService(
	repo interfaces.CarManager,
	apiClient external_api.ExternalAPIClient,
	userProvider interfaces.UserProvider) car.Service {

	re := regexp.MustCompile(pattern)
	return &CarService{
		CarRepo:      repo,
		UserProvider: userProvider,
		re:           re,
		CarAPI:       apiClient,
	}
}

func (s *CarService) CreateCars(regNums []string) error {
	cars := make([]*models.Car, len(regNums))
	for i, regNum := range regNums {
		if !s.re.MatchString(regNum) {
			log.Error().Str("regNum", regNum).Msg("reg num is not in valid")
			return errs.NewInvalidRegNumError(regNum)
		}
		info, err := s.CarAPI.GetCarInfo(regNum)
		if err != nil {
			log.Error().Err(err).Str("regNum", regNum).Msg("error while getting info from external API")
			return external_api.NewExternalAPIError(err)
		}
		user, err := s.UserProvider.GetByFullName(
			info.Owner.Name,
			info.Owner.Surname,
			info.Owner.Patronymic,
		)
		if errors.Is(err, errs.UserNotFound) || errors.Is(err, gorm.ErrRecordNotFound) {
			log.Info().
				Str("name", info.Owner.Name).
				Str("surname", info.Owner.Surname).
				Str("patronymic", info.Owner.Patronymic).
				Msg("user not found in database, creating new")

			newUser := models.User{
				Name:       info.Owner.Name,
				Surname:    info.Owner.Surname,
				Patronymic: &info.Owner.Patronymic,
			}
			err = s.UserProvider.CreateUser(&newUser)
			if err != nil {
				log.Error().Err(err).Msg("error while creating user")
				return err
			}
			user = &newUser

		} else if err != nil {
			log.Error().Err(err).
				Str("name", info.Owner.Name).
				Str("surname", info.Owner.Surname).
				Str("patronymic", info.Owner.Patronymic).
				Msg("error while finding user in database")
			return err
		}
		cars[i] = &models.Car{
			RegNum:  regNum,
			Model:   info.Model,
			Mark:    info.Mark,
			Year:    int32(info.Year),
			OwnerID: &user.ID,
			Owner:   user,
		}
	}
	return s.CarRepo.CreateCars(cars)
}

func (s *CarService) GetCars(pag *helpers.PaginationParams, filters *models.CarFilter) ([]models.Car, int64, error) {
	cars, err := s.CarRepo.GetCars(int(pag.Offset), int(pag.Limit), filters)
	if err != nil {
		log.Error().Err(err).Msg("error while getting cars from database")
		return nil, 0, err
	}
	lastOffset, err := s.CarRepo.GetLastOffset(filters)
	if err != nil {
		log.Error().Err(err).Msg("error while getting count of cars")
		return nil, 0, err
	}
	log.Debug().Int64("lastOffset", lastOffset).Msg("got lastOffset from database")
	return cars, lastOffset, nil
}

func (s *CarService) GetCarByID(id uuid.UUID) (*models.Car, error) {
	carInfo, err := s.CarRepo.GetByID(id)
	if err != nil {
		log.Error().Err(err).Msg("error while getting car from database")
		return nil, err
	}
	return carInfo, nil
}

func (s *CarService) UpdateCarByID(id uuid.UUID, upd *helpers.CarUpdates) error {
	if upd.OwnerID != "" {
		userID, err := uuid.Parse(upd.OwnerID)
		if err != nil {
			log.Error().
				Err(err).
				Msg("failed to parse user ID")
			return err
		}
		_, err = s.UserProvider.GetByID(userID)
		if err != nil {
			log.Error().
				Err(err).
				Str("ownerID", userID.String()).
				Str("carID", id.String()).
				Msg("error while finding owner")
		}
	}
	return s.CarRepo.UpdateCar(id, upd)
}

func (s *CarService) DeleteCarByID(id uuid.UUID) error {
	return s.CarRepo.DeleteByID(id)
}
