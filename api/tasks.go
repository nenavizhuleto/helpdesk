package api

import (
	"github.com/gofiber/fiber/v2"

	"application/auth"
	"application/megaplan"
	"application/models"
)

func GetTasks(c *fiber.Ctx) error {
	i := auth.GetIdentity(c)
	if i == nil {
		return models.ErrNotIdentified
	}
	tasks, err := models.GetTasksForUser(i.User.ID)
	if err != nil {
		return models.ErrEntityNotFound
	}
	return c.JSON(tasks)
}

func CreateTask(c *fiber.Ctx) error {
	i := auth.GetIdentity(c)
	if i == nil {
		return models.ErrNotIdentified
	}
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return models.ErrAPIInvalidRequestBody
	}

	newTask, err := megaplan.MP.HandleCreateTask(i, &task)
	if err != nil {
		return models.ErrMegaplan
	}

	if err := models.SaveTaskForUser(i.User.ID, newTask); err != nil {
		return models.ErrDatabase
	}

	return c.JSON(newTask)
}
