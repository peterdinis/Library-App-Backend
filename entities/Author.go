package entities

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Author - entita pre autora
type Author struct {
	ID               uint      `gorm:"primaryKey"`
	Name             string    `gorm:"size:255;not null"`
	Description      string    `gorm:"size:500"`
	LiteraryPeriod   string    `gorm:"size:255"`
	DateOfBirth      time.Time `gorm:"not null"`
	DateOfDeath      *time.Time // Ukazuje, že dátum smrti je voliteľný
	NumberOfBooks    int       `gorm:"default:0"` // Počet kníh autora
	Books            []Book    `gorm:"foreignKey:AuthorID"` // Relácia s knihami
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
