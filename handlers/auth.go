package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"application/models"
)

func HandleRegister(c *fiber.Ctx) error {
	var user models.User
	log.Println(string(c.Body()))
	if err := c.BodyParser(&user); err != nil {
		return models.ErrAPIInvalidRequestBody
	}
	newUser, err := models.NewUser(c.IP(), user.Name, user.Phone)
	if err != nil {
		return models.ErrDatabase
	}
	return c.JSON(newUser)
}
