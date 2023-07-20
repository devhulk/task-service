package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func checkPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name":    "devhulk tpl",
			"version": "0",
		})
	})

	app.Listen(checkPort())
}
