package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/peterdinis/library-app/backend/entities"
	"github.com/peterdinis/library-app-backend/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/postgres"
)

// Database instance
type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// Connect function
func Connect() {
	// Fetch the database port from config
	p := config.Config("DB_PORT")

	// Parse the port from string to int
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println("Error parsing str to int")
	}

	// Form the Data Source Name (DSN) for the database connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", 
		config.Config("DB_HOST"), 
		config.Config("DB_USER"), 
		config.Config("DB_PASSWORD"),  
		config.Config("DB_NAME"), 
		port)

	// Open a connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	// Log the successful connection
	log.Println("Connected")
	// Set the logger for the database
	db.Logger = logger.Default.LogMode(logger.Info)

	// Log running migrations (if applicable)
	log.Println("running migrations")
	err = db.AutoMigrate(&Category{}, &Book{}, &Author{}, &Publisher{})
	// Assign the database instance to the global DB variable
	DB = Dbinstance{
		Db: db,
	}
}
