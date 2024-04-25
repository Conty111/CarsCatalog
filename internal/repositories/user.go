package repositories

import (
	"github.com/Conty111/CarsCatalog/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByID(id uuid.UUID) (*models.User, error) {
	user := &models.User{}
	tx := r.db.Where("id = ?", id).First(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, UserNotFound
	}
	return user, nil
}

func (r *UserRepository) GetByFullName(name, surname, patronymic string) (*models.User, error) {
	user := models.User{}
	tx := r.db.
		Where("name = ?", name).
		Where("surname = ?", surname).
		Where("patronymic = ?", patronymic).
		First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, UserNotFound
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) UpdateByID(id uuid.UUID, updates interface{}) error {
	return r.db.
		Model(models.User{}).
		Where("id = ?", id).
		Updates(updates).
		Error
}

func (r *UserRepository) DeleteByID(id uuid.UUID) error {
	user := models.User{}
	user.ID = id

	return r.db.
		Model(models.User{}).
		Delete(&user).
		Error
}
