package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User database model
type User struct {
	BaseModel
	Name       string  `gorm:"column:name;co"`
	Surname    string  `gorm:"column:surname"`
	Patronymic *string `gorm:"column:patronymic"`
	Cars       []*Car  `gorm:"foreignKey:OwnerID;references:ID"`
}

func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
}
