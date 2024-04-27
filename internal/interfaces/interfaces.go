package interfaces

import (
	"github.com/Conty111/CarsCatalog/internal/models"
	"github.com/google/uuid"
)

//go:generate go run github.com/vektra/mockery/v3 --name CarManager --output ../../test/mocks
type CarManager interface {
	GetByID(id uuid.UUID) (*models.Car, error)
	DeleteByID(id uuid.UUID) error
	GetCars(offset int, limit int, filters *models.CarFilter) ([]models.Car, error)
	GetLastOffset(filters *models.CarFilter) (int64, error)
	UpdateCar(id uuid.UUID, updates interface{}) error
	CreateCars(cars []*models.Car) error
}

//go:generate go run github.com/vektra/mockery/v3 --name UserProvider --output ../../test/mocks
type UserProvider interface {
	GetByID(id uuid.UUID) (*models.User, error)
	GetByFullName(name, surname, patronymic string) (*models.User, error)
	CreateUser(user *models.User) error
}

//go:generate go run github.com/vektra/mockery/v3 --name UserManager --output ../../test/mocks
type UserManager interface {
	GetByID(id uuid.UUID) (*models.User, error)
	GetByFullName(name, surname, patronymic string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateByID(id uuid.UUID, updates interface{}) error
	DeleteByID(id uuid.UUID) error
}
