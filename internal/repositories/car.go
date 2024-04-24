package repositories

import (
	"github.com/Conty111/CarsCatalog/internal/models"
	"gorm.io/gorm"
)

type CarRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) *CarRepository {
	return &CarRepository{
		db: db,
	}
}

func (r *CarRepository) GetCars(offset int, limit int, filters *models.CarFilter) ([]models.Car, error) {
	var cars []models.Car
	err := r.db.
		Model(models.Car{}).
		Offset(offset).
		Limit(limit).
		Find(&cars).
		Error

	if err != nil {
		return nil, err
	}
	return cars, nil
}

func setFilters(tx *gorm.DB, filters *models.CarFilter) *gorm.DB {

}
