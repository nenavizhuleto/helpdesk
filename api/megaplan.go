package api

import (
	"application/megaplan"
	"application/models/v2"
	"log"

	"github.com/gofiber/fiber/v2"
)


func HandleMegaplanEvent(c *fiber.Ctx) error {
	var event megaplan.TaskEvent
	if err := c.BodyParser(&event); err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}

	if event.Model != "Task" {
		return c.SendStatus(200)
	}

	log.Printf("Event: %#v", event)

	dto := event.Data
	task, err := models.TaskByID(dto.ID)
	if err != nil {
		log.Printf("event: %s", err.Error())
		return c.SendStatus(200)
	}

	

	task.Status = dto.GetStatus()

	if err := task.Update(); err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}

	return c.SendStatus(200)
}
