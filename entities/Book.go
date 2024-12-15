package entities

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:255;not null"`
	Description string    `gorm:"size:500"`
	Year        int       `gorm:"not null"`
	Pages       int       `gorm:"not null"`
	Images      string    `gorm:"size:500"`
	IsAvailable bool      `gorm:"default:true"`
	CategoryID  uint      `gorm:"not null"`
	Category    Category  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AuthorID    uint      `gorm:"not null"`
	Author      Author    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Relácia s autorom
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
