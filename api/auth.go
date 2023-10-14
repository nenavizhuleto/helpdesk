package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"application/auth/v2"
	"application/models/v2"
)

func GetIdentity(c *fiber.Ctx) error {
	dev, err := auth.GetIdentity(c.IP())
	if err != nil {
		return fmt.Errorf("identity: %w", err)
	}

	return c.JSON(dev)
}

func Register(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return fmt.Errorf("register: %w", err)
	}

	user.ID = uuid.NewString()

	device, err := auth.Register(c.IP(), &user)
	if err != nil {
		return fmt.Errorf("register: %w", err)
	}
	return c.JSON(device)
}
