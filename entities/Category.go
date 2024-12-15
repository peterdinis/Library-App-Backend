package entities

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Category struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:255;not null;unique"`
	Description string    `gorm:"size:500"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Books       []Book    `gorm:"foreignKey:CategoryID"`
}