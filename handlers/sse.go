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


func ServeClient(c chan<- Event, done <-chan bool, isDied <-chan bool, clientID string) {
    c <- NewEvent(clientID, "Started serving " + clientID)
    notify := data.DB.Subscribers.Subscribe(clientID)
    go func() {
        for {
            select {
            case <-isDied:
                return
            default:
                c <- NewEvent(clientID, "isAlive?")
                time.Sleep(time.Second * 10)
            }
        }
    }()
    for {
        select {
        case <-done:
            return
        case <-notify:
            c <- NewEvent(clientID, "update")
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
            fmt.Fprintf(w, "event: %s\n", event.Name)
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
