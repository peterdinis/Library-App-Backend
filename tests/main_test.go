package tests

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApp(t *testing.T) {
	app := fiber.New()

	// Define a simple route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	// Create a request
	req := httptest.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req)

	// Check if status code is 200
	assert.Equal(t, resp.StatusCode, 200)

	// Check if the response body is "Hello, Fiber!"
	body := string(resp.Body())
	assert.Equal(t, body, "Hello, Fiber!")
}