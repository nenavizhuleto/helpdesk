package handlers

import (
	"log"

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

func HandleMegaPlanGetEntity(c *fiber.Ctx) error {
	entity := c.Params("entity")
	entities, err := megaplan.MP.Get("/" + entity)
	if err != nil {
		return err
	}

	log.Printf("Entities: %v", entities)

	return c.JSON(entities)
}
