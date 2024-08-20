package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	// Define a simple GET route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Define another route for a student endpoint
	app.Get("/student", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name":   "faizan",
			"rollNo": "14",
			"grade":  "A",
			"cgpa":   9.7,
		})
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
