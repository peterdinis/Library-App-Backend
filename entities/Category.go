import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Category struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255;not null;unique"`
	Description string `gorm:"size:500"`
	CreatedAt   gorm.DeletedAt
	UpdatedAt   gorm.DeletedAt
}