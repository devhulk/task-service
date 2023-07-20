package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Task struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

func create(c *fiber.Ctx, db *sql.DB) error {

	var task Task

	if err := c.BodyParser(&task); err != nil {
		log.Fatalln("Error parsing insert response body.", err)
	}

	id, err1 := uuid.NewRandom()
	if err1 != nil {
		log.Fatalln("Could not create uuid.", err1)

	}

	insert := `
	INSERT INTO tasks (id, title, description, status)
	VALUES ($1, $2, $3, $4)
	`

	_, err2 := db.Exec(insert, id.String(), task.Title, task.Description, task.Status)
	if err2 != nil {
		log.Fatalln("Error executing insert", err2)
		return c.JSON(fiber.Map{
			"error": "Could not insert. Try a unique id.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "new task created successfully",
		"id":      id.String(),
	})
}

func list(c *fiber.Ctx, db *sql.DB) error {
	var task Task
	var tasks []Task

	rows, err := db.Query("SELECT id, title, description, status FROM tasks")
	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
		c.JSON("Error occured reading database.")
	}

	for rows.Next() {
		rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)
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
