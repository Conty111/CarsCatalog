package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Car database model
type Car struct {
	BaseModel
	RegNum  string    `gorm:"primaryKey"`
	Mark    string    `gorm:"column:mark"`
	Model   string    `gorm:"column:model"`
	Year    int32     `gorm:"column:year"`
	OwnerID uuid.UUID `gorm:"column:owner_id;index"`
	Owner   *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (c *Car) BeforeCreate(_ *gorm.DB) (err error) {
	c.ID = uuid.New()
	return nil
}
