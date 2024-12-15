package entities

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Publisher struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:255;not null;unique"`
	Description string    `gorm:"size:500"`
	FoundedDate time.Time `gorm:"not null"`
	IsActive    bool      `gorm:"default:true"`
	LogoImage   string    `gorm:"size:500"`
	Books       []Book    `gorm:"foreignKey:PublisherID"` // Relácia s knihami
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
