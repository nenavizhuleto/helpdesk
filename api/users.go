package api

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"application/models/v2"
)

func SetUsersRoutes(path string, router fiber.Router) {
	users := router.Group(path)
	users.Get("/", GetUsers)
	users.Get("/:id", GetUser)
	users.Post("/", CreateUser)
	users.Put("/", UpdateUser)
	users.Delete("/", DeleteUser)

	users.Get("/:id/tasks", GetUserTasks)
	users.Post("/:id/tasks", CreateUserTask)
}

func GetUsers(c *fiber.Ctx) error {
	users, err := models.UsersAll()
	if err != nil {
		return fmt.Errorf("getUsers: %w", err)
	}

	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := models.User{
		ID: id,
	}

	if err := user.Fetch(); err != nil {
		return fmt.Errorf("getUser: %w", err)
	}

	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	var user = models.NewUser()
	if err := c.BodyParser(user); err != nil {
		return fmt.Errorf("createUser: %w", err)
	}

	if err := user.Create(); err != nil {
		return fmt.Errorf("createUser: %w", err)
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return fmt.Errorf("updateUser: %w", err)
	}

	if err := user.Update(); err != nil {
		return fmt.Errorf("updateUser: %w", err)
	}

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return fmt.Errorf("deleteUser: %w", err)
	}

	if err := user.Delete(); err != nil {
		return fmt.Errorf("deleteUser: %w", err)
	}

	return c.JSON(user)
}

func GetUserTasks(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := models.UserByID(id)
	if err != nil {
		return err
	}

	tasks, err := models.TasksByUser(user)
	if err != nil {
		return err
	}

	return c.JSON(tasks)

}

func CreateUserTask(c *fiber.Ctx) error {
	id := c.Params("id")

	dev, err := models.DeviceGetByIP(c.IP())
	if err != nil {
		return fmt.Errorf("user: %w", err)
	}

	var task = models.NewTask()

	if err := c.BodyParser(&task); err != nil {
		return fmt.Errorf("user: %w", err)
	}

	user, err := models.UserByID(id)
	if err != nil {
		return err
	}

	task.User = *user
	task.Company = dev.Company
	task.Branch = dev.Branch

	if err := task.Create(); err != nil {
		return err
	}

	return c.JSON(task)

}
