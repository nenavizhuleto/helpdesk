package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"application/auth"
	"application/megaplan"
	"application/models"
)

func HandleGetTaskNew(c *fiber.Ctx) error {
	return c.Render("pages/tasks", fiber.Map{})
}

func HandlePostTaskNew(c *fiber.Ctx) error {
	i := auth.GetIdentity(c)
	task := &models.Task{
		Name:    c.FormValue("title"),
		Subject: c.FormValue("description"),
	}

	c.Set("HX-Trigger", "issue-created")

	if err := megaplan.MP.HandleCreateTask(i, task); err != nil {
		return err
	}

	log.Printf("formData: %v", task)
	return c.SendStatus(200)
}
