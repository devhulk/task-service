package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New())

	db, err := initDB()

	if err != nil {
		log.Fatalln("DB Connection Failed")
	}

	defer db.Close()

	app.Get("/", func(c *fiber.Ctx) error {
		return list(c, db)
	})
	app.Get("/:id", func(c *fiber.Ctx) error {
		return getTask(c, db)
	})
	app.Post("/create", func(c *fiber.Ctx) error {
		return create(c, db)
	})
	app.Put("/update/:id", func(c *fiber.Ctx) error {
		return update(c, db)
	})
	app.Delete("/remove/:id", func(c *fiber.Ctx) error {
		return remove(c, db)
	})

	log.Fatalln(app.Listen(checkPort()))
}
