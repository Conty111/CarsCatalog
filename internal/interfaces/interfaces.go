package interfaces

type CarManager interface {
	GetByID()
	DeleteByID()
	GetCars()
	UpdateCars()
	CreateCars()
}
