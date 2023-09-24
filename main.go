package main

import (
	"application/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)


func main() {
  engine := html.New("./www", ".html")

  // For debug
  engine.Reload(true)

  app := fiber.New(fiber.Config{
    Views: engine,
  })


  app.Static("/", "./public")

  app.Use(handlers.IdentityMiddleware)

  app.Get("/", handlers.HandleIndex)
  app.Get("/sse", handlers.HandleSSE)
  app.Post("/issues/:uuid", handlers.HandleIssues)
  app.Post("/issue/send", handlers.HandleIssueSend)
  app.Get("/test", handlers.HandleTest)
  app.Post("/issue/:uuid/:index", handlers.HandleTestChangeStatus)
  log.Fatal(app.Listen(":3000"))
}
