package api

import (
	"fmt"

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

	// users.Get("/:id/tasks", GetUserTasks)
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
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return fmt.Errorf("createUser: %w", err)
	}

	user.ID = uuid.NewString()
	if err := user.Create(c.IP()); err != nil {
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
