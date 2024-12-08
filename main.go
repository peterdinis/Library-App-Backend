package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"os"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

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
