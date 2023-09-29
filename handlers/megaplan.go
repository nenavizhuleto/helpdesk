package handlers

import (
	"bufio"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"

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

var (
	MegaPlanEventChannel     = make(chan bool, 255)
	MegaPlanEventDataChannel = make(chan MegaPlanEvent, 255)
)

func HandleMegaPlan(c *fiber.Ctx) error {
	data := MegaPlanEvent{}
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	MegaPlanEventChannel <- true
	MegaPlanEventDataChannel <- data

	log.Printf("Data: %v", data)

	return nil
}

func HandleMegaPlanEventStream(c *fiber.Ctx) error {
	event := <-MegaPlanEventDataChannel

	return c.Render("event", event)
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

func HandleMegaPlanSSE(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Tranfer-Encoding", "chunked")

	clientID := GetUserID(c)

	eventStream := make(chan MegaPlanEvent)
	// go ServeClient(eventStream, done, isDied, clientID)

	c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		log.Printf("New SSE connection with ClientID: %s\n", clientID)
		for {
			event := <-MegaPlanEventChannel
			fmt.Fprintf(w, "event: %s\n", "megaplan")
			fmt.Fprintf(w, "data: %s\n\n", "megaplan")

			log.Printf("event: %v", event)
			err := w.Flush()
			if err != nil {
				log.Printf("Error while flushing: %v. Closing http connection.\n", err)
				break
			}
		}
		close(eventStream)
	}))

	return nil
}
