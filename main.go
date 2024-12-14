package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/peterdinis/library-app-backend/database"
	"github.com/peterdinis/library-app-backend/routes"
	"os"
	"log"
	"github.com/gofiber/fiber/v2/middleware/logger"
 _ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	database.Connect()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())
	routes.ExampleRoute(app)

	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = ":5000"
	} else if appPort[0] != ':' {
		appPort = ":" + appPort
	}

	err = app.Listen(appPort)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
