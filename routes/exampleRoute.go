package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/peterdinis/library-app-backend/controllers"
)

func ExampleRoute(app *fiber.App) {
	app.Get("/", controllers.ExampleController)
}