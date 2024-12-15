package entities

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Author struct {
	ID               uint      `gorm:"primaryKey"`
	Name             string    `gorm:"size:255;not null"`
	Description      string    `gorm:"size:500"`
	LiteraryPeriod   string    `gorm:"size:255"`
	DateOfBirth      time.Time `gorm:"not null"`
	DateOfDeath      *time.Time
	Image            string    `gorm:"size:500"`
	NumberOfBooks    int       `gorm:"default:0"`
	Books            []Book    `gorm:"foreignKey:AuthorID"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
