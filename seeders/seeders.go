package seeders

import (
	"log"

	"github.com/peterdinis/library-app-backend/database"
	"github.com/peterdinis/library-app-backend/models"
)

/* Later this will be global seeder for all tables */
// SeedCategories inserts initial data into the categories table
func SeedCategories() {
	categories := []models.Category{
		{Name: "Fiction", Description: "Books that contain fictional stories."},
		{Name: "Non-Fiction", Description: "Books based on real events and facts."},
		{Name: "Science", Description: "Books that explore scientific concepts and discoveries."},
		{Name: "History", Description: "Books about historical events and figures."},
	}

	// Use GORM's Create method to insert data
	for _, category := range categories {
		if err := database.DB.Db.Create(&category).Error; err != nil {
			log.Printf("Failed to seed category: %s, error: %v", category.Name, err)
		} else {
			log.Printf("Seeded category: %s", category.Name)
		}
	}
}