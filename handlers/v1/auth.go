package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"application/models/v1"
)

func HandleRegister(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return fmt.Errorf("register: %w", err)
	}

	user.ID = uuid.NewString()
	if err := user.Create(); err != nil {
		return fmt.Errorf("register: %w", err)
	}

	return c.JSON(user)
}
