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

func GetTask(c *fiber.Ctx) error {
	i := auth.GetIdentity(c)
	if i == nil {
		return models.ErrNotIdentified
	}

	task_id := c.Params("id")

	task, err := models.GetTaskForUser(i.User.ID, task_id)
	if err != nil {
		return models.ErrEntityNotFound
	}
	if taskMP, err := megaplan.MP.HandleFetchTaskUpdates(i, task); err != nil {
		return models.ErrMegaplan
	} else {
		if taskMP.Status != task.Status {
			task.Status = taskMP.Status
			task.Name = taskMP.Name
			models.UpdateTaskForUser(i.User.ID, task)
		}
	}

	return c.JSON(task)
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

	// FIXME: Replacing returned task from megaplan with original subject
	newTask.Subject = task.Subject

	if err := models.SaveTaskForUser(i.User.ID, newTask); err != nil {
		return models.ErrDatabase
	}
	if err = newTask.Prettify(); err != nil {
		return fiber.NewError(500, err.Error())
	}

	return c.JSON(newTask)
}
