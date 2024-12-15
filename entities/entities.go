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
	Author      Author    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PublisherID uint      `gorm:"not null"`              
	Publisher   Publisher `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` 
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Category struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:255;not null;unique"`
	Description string    `gorm:"size:500"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Books       []Book    `gorm:"foreignKey:CategoryID"`
}

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
