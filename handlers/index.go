package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func HandleIndex(c *fiber.Ctx) error {
	return c.Render("layouts/main", fiber.Map{})
}

func HandleIdentityInfo(c *fiber.Ctx) error {
	return c.Render("pages/identity", fiber.Map{})
}
