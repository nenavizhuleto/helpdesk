package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"application/auth/v1"
	"application/megaplan"
	"application/models/v1"
)

func SetTasksRoutes(path string, router fiber.Router) {
	tasks := router.Group(path)
	tasks.Get("/", GetTasks)
	tasks.Post("/", CreateTask)
	tasks.Put("/", UpdateTask)
	tasks.Delete("/", DeleteTask)
}

func GetTasks(c *fiber.Ctx) error {
	tasks, err := models.TasksAll()
	if err != nil {
		return fmt.Errorf("getTasks: %w", err)
	}

	return c.JSON(tasks)
}

func CreateTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return fmt.Errorf("createTask: %w", err)
	}

	task.BeforeCreate = func(t *models.Task) error {
		fmt.Println("TODO: Create Task in Megaplan")
		t.ID = uuid.NewString()
		return nil
	}

	if err := task.Create(); err != nil {
		return fmt.Errorf("createTask: %w", err)
	}

	return c.JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return fmt.Errorf("updateTask: %w", err)
	}

	if err := task.Update(); err != nil {
		return fmt.Errorf("updateTask: %w", err)
	}

	return c.JSON(task)
}

func DeleteTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return fmt.Errorf("deleteTask: %w", err)
	}

	if err := task.Delete(); err != nil {
		return fmt.Errorf("deleteTask: %w", err)
	}

	return c.JSON(task)
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

	if relevantTask, err := megaplan.MP.HandleFetchTaskUpdates(task); err != nil {
		return models.ErrMegaplan
	} else {
		if relevantTask.Status != task.Status {
			task.Status = relevantTask.Status
			models.UpdateTaskForUser(i.User.ID, task)
		} else {
			task.Comments = relevantTask.Comments
		}
	}

	return c.JSON(task)
}

func CreateTask1(c *fiber.Ctx) error {
	i := auth.GetIdentity(c)
	if i == nil {
		return models.ErrNotIdentified
	}
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return models.ErrAPIInvalidRequestBody
	}

	createdTask, err := megaplan.MP.HandleCreateTask(i, &task)
	if err != nil {
		return models.ErrMegaplan
	}

	// FIXME: Replacing returned task from megaplan with original subject
	createdTask.Subject = task.Subject

	if err := models.SaveTaskForUser(i.User.ID, createdTask); err != nil {
		return models.ErrDatabase
	}
	if err = createdTask.Prettify(); err != nil {
		return fiber.NewError(500, err.Error())
	}

	return c.JSON(createdTask)
}
