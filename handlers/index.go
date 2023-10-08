package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func HandleMain(c *fiber.Ctx) error {
	return c.Render("screens/system", fiber.Map{}, "layout/system")
}

func HandleDevNull(c *fiber.Ctx) error {
	return c.SendString("")
}

func HandleIndex(c *fiber.Ctx) error {
	// return c.Render("layouts/main", fiber.Map{})
	// return c.Render("index", fiber.Map{}, "layout/main")
	return c.Redirect("/system")
}

func HandleIdentityInfo(c *fiber.Ctx) error {
	return c.Render("pages/identity", fiber.Map{})
}

func HandleSvelte(c *fiber.Ctx) error {
	return c.Render("index.html", nil)
}
