package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"application/auth"
	"application/models"
)

func HandleMain(c *fiber.Ctx) error {
	identity := auth.GetIdentity(c)
	tasks, err := models.GetTasksForUser(identity.User.ID)
	log.Println(tasks)
	if err != nil {
		return err
	}

	return c.Render("screens/system", fiber.Map{
		"Tasks": tasks,
	}, "layout/system")
}

func HandleIndex(c *fiber.Ctx) error {
	// return c.Render("layouts/main", fiber.Map{})
	// return c.Render("index", fiber.Map{}, "layout/main")
	return c.Redirect("/system")
}

func HandleIdentityInfo(c *fiber.Ctx) error {
	return c.Render("pages/identity", fiber.Map{})
}
