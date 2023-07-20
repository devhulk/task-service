package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Task struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

func root(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"root": true,
	})
}

func create(c *fiber.Ctx, db *sql.DB) error {
	return c.JSON(fiber.Map{
		"name": "create stuff here",
	})
}

func list(c *fiber.Ctx, db *sql.DB) error {
	var task Task
	var tasks []Task

	rows, err := db.Query("SELECT title, description, status FROM tasks")
	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
		c.JSON("Error occured reading database.")
	}

	for rows.Next() {
		rows.Scan(&task.Title, &task.Description, &task.Status)
		tasks = append(tasks, task)
	}

	return c.JSON(tasks)

}

func update(c *fiber.Ctx, db *sql.DB) error {
	return c.JSON(fiber.Map{
		"name": "update stuff here",
	})
}

func remove(c *fiber.Ctx, db *sql.DB) error {
	return c.JSON(fiber.Map{
		"name": "delete stuff here",
	})
}
