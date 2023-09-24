package handlers

import (
	"bufio"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"

	"application/data"
)

type Event struct {
	Name string
	Kind string
	Data string
}

func NewEvent(name, kind, data string) Event {
	return Event{
		Name: name,
		Kind: kind,
		Data: data,
	}
}

func (e *Event) GetName() string {
	return e.Kind + ":" + e.Name
}

func ServeClient(c chan<- Event, done <-chan bool, isDied <-chan bool, clientID string) {
	c <- NewEvent(clientID, "update", "Started serving "+clientID)
	notify := data.DB.Subscribers.Subscribe(clientID)
	go func() {
		for {
			select {
			case <-isDied:
				return
			default:
				c <- NewEvent(clientID, "connection", "isAlive?")
				time.Sleep(time.Second * 10)
			}
		}
	}()
	for {
		select {
		case <-done:
			data.DB.Subscribers.Unsubscribe(clientID)
			return
		case <-notify.Channel:
			c <- NewEvent(clientID, "update", "update")
		}
	}
}

func HandleSSE(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Tranfer-Encoding", "chunked")

	clientID := GetClientID(c)

	eventStream := make(chan Event)
	done := make(chan bool, 1)
	isDied := make(chan bool, 1)
	go ServeClient(eventStream, done, isDied, clientID)

	c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		log.Printf("New SSE connection with ClientID: %s\n", clientID)
		for {
			event := <-eventStream
			fmt.Fprintf(w, "event: %s\n", event.GetName())
			fmt.Fprintf(w, "data: %s\n\n", event.Data)

			log.Printf("event: %v", event)
			err := w.Flush()
			if err != nil {
				isDied <- true
				log.Printf("Error while flushing: %v. Closing http connection.\n", err)
				break
			}
		}
		done <- true
		close(eventStream)
		close(isDied)
		close(done)
	}))

	return nil
}
