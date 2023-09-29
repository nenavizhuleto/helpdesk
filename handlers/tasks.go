package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"application/megaplan"
	"application/models"
)

func HandleGetTaskNew(c *fiber.Ctx) error {
	return c.Render("pages/tasks", fiber.Map{})
}

func HandlePostTaskNew(c *fiber.Ctx) error {
	user := c.Locals("User").(models.User)
	task := &models.Task{
		ID:      uuid.NewString(),
		User:    user,
		Name:    c.FormValue("title"),
		Subject: c.FormValue("description"),
		Status:  "Новое обращение",
	}

	c.Set("HX-Trigger", "issue-created")

	if err := megaplan.MP.HandleCreateTask(task); err != nil {
		return err
	}

	log.Printf("formData: %v", task)
	return c.SendStatus(200)
}
