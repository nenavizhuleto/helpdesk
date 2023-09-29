package handlers

import (
	"github.com/gofiber/fiber/v2"

	"application/models"
)

func HandleAuth(c *fiber.Ctx) error {
	return c.Render("auth", nil)
}

func HandleSignup(c *fiber.Ctx) error {
	name := c.FormValue("name")
	phone := c.FormValue("phone")
	if err := models.NewUser(c.IP(), name, phone); err != nil {
		return err
	}
	c.Set("HX-Redirect", "/")
	return c.SendStatus(200)
}
