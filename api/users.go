package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"application/models/v1"
)

func SetUsersRoutes(path string, router fiber.Router) {
	users := router.Group(path)
	users.Get("/", GetAllUsers)
	users.Get("/:id", GetUser)
	users.Post("/", CreateUser)
	users.Put("/", UpdateUser)
	users.Delete("/", DeleteUser)

	users.Get("/:id/tasks", GetUserTasks)
}

func GetAllUsers(c *fiber.Ctx) error {
	users, err := models.UsersAll()
	if err != nil {
		return fmt.Errorf("getAllUsers: %w", err)
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

func GetUserTasks(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println(id)

	tasks, err := models.TasksByUserID(id)
	if err != nil {
		return fmt.Errorf("getUserTasks: %w", err)
	}

	return c.JSON(tasks)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return fmt.Errorf("createUser: %w", err)
	}

	user.ID = uuid.NewString()
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
