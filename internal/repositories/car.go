package repositories

import (
	"errors"
	"github.com/Conty111/CarsCatalog/internal/errs"
	"github.com/Conty111/CarsCatalog/internal/models"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
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
	err := r.db.Model(&car).
		Preload(clause.Associations).
		Where("id = ?", id).
		First(&car).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewCarNotFoundError(id)
		}
		return nil, err
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
	log.Debug().Msg("got cars from database successfully")
	return cars, nil
}

func setFilters(tx *gorm.DB, filters *models.CarFilter) {
	if filters.Model != "" {
		tx.Where("model ILIKE ?", "%"+filters.Model+"%")
	}
	if filters.Mark != "" {
		tx.Where("mark ILIKE ?", "%"+filters.Mark+"%")
	}
	if filters.RegNum != "" {
		tx.Where("reg_num ILIKE ?", "%"+filters.RegNum+"%")
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

	tx := r.db.Model(&models.Car{}).Delete(&car)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return errs.NewCarNotFoundError(id)
		}
	}
	if tx.RowsAffected == 0 {
		return errs.NewCarNotFoundError(id)
	}
	return nil
}

func (r *CarRepository) UpdateCar(id uuid.UUID, updates interface{}) error {
	tx := r.db.
		Debug().
		Model(&models.Car{}).
		Where("id = ?", id).
		Updates(updates)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return errs.NewCarNotFoundError(id)
		}
	}
	if tx.RowsAffected == 0 {
		return errs.NewCarNotFoundError(id)
	}
	return nil
}

func (r *CarRepository) CreateCars(cars []*models.Car) error {
	return r.db.Debug().Model(&models.Car{}).Create(cars).Error
}
