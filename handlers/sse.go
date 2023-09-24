package handlers

import (
	"application/data"
	"bufio"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type Event struct {
    Name string
    Data string
}

func NewEvent(name, data string) Event {
    return Event{
        Name: name,
        Data: data,
    }
}

func ServeClient(c chan<- Event, done <-chan bool, id string) {
    c <- NewEvent(id, "Started serving " + id)
    notifyChan := data.Subs.Subscribe(id)
    
    for {
        select {
        case <-done:
            return
        case <-notifyChan:
            time.Sleep(time.Second)
            c <- NewEvent(id, time.Now().String())
        }
    }
}

func HandleSSE(c *fiber.Ctx) error {
    c.Set("Content-Type", "text/event-stream")
    c.Set("Cache-Control", "no-cache")
    c.Set("Connection", "keep-alive")
    c.Set("Tranfer-Encoding", "chunked")
    
    id := c.Cookies("uuid")

    eventStream := make(chan Event)
    done := make(chan bool)
    go ServeClient(eventStream, done, id)

    c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
        log.Printf("New SSE connection with ID: %s\n", id)
        for {
            event := <-eventStream
            fmt.Fprintf(w, "event: %s\n", event.Name)
            fmt.Fprintf(w, "data: %s\n\n", event.Data)
            
            log.Printf("event: %v", event)
            err := w.Flush()
            if err != nil {
                log.Printf("Error while flushing: %v. Closing http connection.\n", err)
                break
            }
        }
        done <- true
        close(eventStream)
        close(done)
    }))


    return nil
}
