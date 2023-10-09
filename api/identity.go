package api

import (
	"github.com/gofiber/fiber/v2"

	"application/auth"
	"application/models"
)

func GetIdentity(c *fiber.Ctx) error {
	i := auth.GetIdentity(c)
	if i == nil {
		return models.ErrNotIdentified
	}

	return c.JSON(i)
}
