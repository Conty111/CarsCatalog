package interfaces

import (
	"github.com/Conty111/CarsCatalog/internal/external_api"
	"github.com/Conty111/CarsCatalog/internal/models"
	"github.com/google/uuid"
)

type CarManager interface {
	GetByID(id uuid.UUID) (*models.Car, error)
	DeleteByID(id uuid.UUID) error
	GetCars(offset int, limit int, filters *models.CarFilter) ([]models.Car, error)
	GetLastOffset(filters *models.CarFilter) (int64, error)
	UpdateCar(id uuid.UUID, updates interface{}) error
	CreateCars(cars []*models.Car) error
}

type UserProvider interface {
	GetByID(id uuid.UUID) (*models.User, error)
	GetByFullName(name, surname, patronymic string) (*models.User, error)
	CreateUser(user *models.User) error
}

type UserManager interface {
	GetByID(id uuid.UUID) (*models.User, error)
	DeleteByID(id uuid.UUID) error
	UpdateCar(id uuid.UUID, updates interface{}) error
	CreateUser(user *models.User) error
}

type CarAPIClient interface {
	GetCarInfo(regNum string) (*external_api.CarData, error)
}
