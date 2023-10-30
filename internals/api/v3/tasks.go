package api

import (
	"fmt"

	"helpdesk/internals/megaplan"
	"helpdesk/internals/models"
	"helpdesk/internals/models/v3/comment"
	"helpdesk/internals/models/v3/task"

	"github.com/gofiber/fiber/v2"
)

func SetTasksRoutes(path string, router fiber.Router) {
	tasks := router.Group(path)
	tasks.Get("/", GetTasks)
	tasks.Get("/:id", GetTask)
	tasks.Post("/", CreateTask)
	tasks.Put("/:id/comment", CommentTask)
	tasks.Delete("/:id", DeleteTask)
}

func CommentTask(c *fiber.Ctx) error {
	var body struct {
		Content string
	}

	if err := c.BodyParser(&body); err != nil {
		return models.NewParseError("comment_task", err)
	}

	id := c.Params("id")
	t, err := task.Get(id)
	if err != nil {
		return err
	}

	content := fmt.Sprintf("#[FROMUSER]: %s", body.Content)
	com, err := megaplan.MP.CommentTask(t.ID, content)
	if err != nil {
		return err
	}

	var comm comment.Comment
	comm.ID = com.ID
	comm.Content = com.Content
	comm.Direction = comment.DirectionFrom

	t.Comments = append(t.Comments, comm)

	return c.JSON(comm)
}

func GetTasks(c *fiber.Ctx) error {
	tasks, err := task.All()
	fmt.Println(tasks)
	if err != nil {
		return fmt.Errorf("getTasks: %w", err)
	}

	return c.JSON(tasks)
}

func GetTask(c *fiber.Ctx) error {
	id := c.Params("id")
	task, err := task.Get(id)
	if err != nil {
		return err
	}

	return c.JSON(task)
}

func CreateTask(c *fiber.Ctx) error {
	var body struct {
		Name    string
		Subject string
	}
	if err := c.BodyParser(&body); err != nil {
		return models.NewParseError("task", err)
	}

	task, err := task.New()
	if err != nil {
		return err
	}

	task.Name = body.Name
	task.Subject = body.Subject

	if err := task.Save(); err != nil {
		return err
	}

	return c.JSON(task)
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")

	task, err := task.Get(id)
	if err != nil {
		return err
	}

	if err := task.Delete(); err != nil {
		return err
	}

	return c.JSON(task)
}
