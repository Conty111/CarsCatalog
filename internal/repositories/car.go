package repositories

import (
	"github.com/Conty111/CarsCatalog/internal/models"
	"github.com/google/uuid"
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
	tx := setFilters(r.db.Model(models.Car{}), filters)
	err := tx.Offset(offset).Limit(limit).Find(&cars).Error
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func setFilters(tx *gorm.DB, filters *models.CarFilter) *gorm.DB {
	tx.Where("model LIKE '%?%'", filters.Model)
	tx.Where("mark LIKE '%?%'", filters.Mark)
	tx.Where("reg_num LIKE '%?%'", filters.RegNum)
	tx.Where("year >= ?", filters.MinYear)
	if filters.MaxYear > filters.MinYear {
		tx.Where("year =< ?", filters.MaxYear)
	}
	return tx
}

func (r *CarRepository) DeleteCar(id uuid.UUID) error {
	car := models.Car{BaseModel: models.BaseModel{ID: id}}
	return r.db.Model(models.Car{}).Delete(&car).Error
}

func (r *CarRepository) UpdateCar(id uuid.UUID, updates map[string]interface{}) error {
	tx := r.db.
		Model(models.Car{}).
		Where("id = ?", id)

}
