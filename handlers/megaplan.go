package handlers

import (
	"github.com/gofiber/fiber/v2"

	"application/megaplan"
)

type MegaPlanEvent struct {
	Uuid  string            `json:"uuid"`
	Event string            `json:"event"`
	Model string            `json:"model"`
	Data  MegaPlanEventData `json:"data"`
}

type MegaPlanEventData struct {
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Status  string `json:"status"`
}

func HandleMegaPlanInfo(c *fiber.Ctx) error {
	return c.JSON(megaplan.MP)
}
