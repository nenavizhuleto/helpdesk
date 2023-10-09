package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"application/models"
)

type Event struct {
	Data  map[string]interface{} `json:"data"`
	Event string                 `json:"event"`
	Model string                 `json:"model"`
}
type Task struct {
	ID     string `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Status string `json:"status" db:"status"`
}

func HandleMegaplanEvent(c *fiber.Ctx) error {
	var event Event
	if err := c.BodyParser(&event); err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}

	log.Printf("Event: %v", event)

	if event.Model == "Task" {
		task := event.Data
		if err := models.UpdateTaskStatus(task["id"].(string), task["status"].(string)); err != nil {
			log.Println(err)
			return c.SendStatus(500)
		}
	}

	return c.SendStatus(200)
}
