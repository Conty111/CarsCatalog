package repositories

import (
	"github.com/Conty111/CarsCatalog/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CarRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) *CarRepository {
	return &CarRepository{
		db: db,
	}
}

func (r *CarRepository) GetByID(id uuid.UUID) (*models.Car, error) {
	var car models.Car
	res := r.db.Model(&car).
		Preload(clause.Associations).
		Where("id = ?", id).
		First(&car)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, CarNotFound
	}
	return &car, nil
}

func (r *CarRepository) GetLastOffset(filters *models.CarFilter) (int64, error) {
	tx := r.db.Model(&models.Car{})

	setFilters(tx, filters)

	var totalCount int64
	err := tx.Count(&totalCount).Error
	if err != nil {
		return 0, err
	}
	return totalCount, nil
}

func (r *CarRepository) GetCars(offset int, limit int, filters *models.CarFilter) ([]models.Car, error) {
	var cars []models.Car
	tx := r.db.Model(&models.Car{})
	setFilters(tx, filters)
	err := tx.
		Preload(clause.Associations).
		Offset(offset).
		Limit(limit).
		Find(&cars).
		Error
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func setFilters(tx *gorm.DB, filters *models.CarFilter) {
	if filters.Model != "" {
		tx.Where("model LIKE ?", "%"+filters.Model+"%")
	}
	if filters.Mark != "" {
		tx.Where("mark LIKE ?", "%"+filters.Mark+"%")
	}
	if filters.RegNum != "" {
		tx.Where("reg_num LIKE ?", "%"+filters.RegNum+"%")
	}
	if filters.MinYear >= 0 {
		tx.Where("year >= ?", filters.MinYear)
	}
	if filters.MaxYear > filters.MinYear {
		tx.Where("year <= ?", filters.MaxYear)
	}
}

func (r *CarRepository) DeleteByID(id uuid.UUID) error {
	car := models.Car{BaseModel: models.BaseModel{ID: id}}
	return r.db.Model(&models.Car{}).Delete(&car).Error
}

func (r *CarRepository) UpdateCar(id uuid.UUID, updates interface{}) error {
	return r.db.
		Model(&models.Car{}).
		Where("id = ?", id).
		Updates(updates).
		Error
}

func (r *CarRepository) CreateCars(cars []*models.Car) error {
	return r.db.Model(&models.Car{}).Create(cars).Error
}
