package handlers

import "github.com/gofiber/fiber/v2"

func HandleIndex(c *fiber.Ctx) error {
  message := "I'm the message"
  return c.Render("landing/index", fiber.Map{
    "Message": message,
  })
}
