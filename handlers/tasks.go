package handlers

import (
	"github.com/gofiber/fiber/v2"

	"application/auth"
	"application/megaplan"
	"application/models"
)

func HandleGetTaskNew(c *fiber.Ctx) error {
	return c.Render("partials/task-modal-new", fiber.Map{})
}

func HandlePostTaskNew(c *fiber.Ctx) error {
	i := auth.GetIdentity(c)
	name := c.FormValue("title")
	subject := c.FormValue("subject")
	task := &models.Task{
		Name:    name,
		Subject: subject,
	}

	newTask, err := megaplan.MP.HandleCreateTask(i, task)
	if err != nil {
		return err
	}

	println("task", newTask)

	if err := models.SaveTaskForUser(i.User.ID, newTask); err != nil {
		return err
	}
	c.Set("HX-Trigger", "task-created")

	return c.SendStatus(200)
}
