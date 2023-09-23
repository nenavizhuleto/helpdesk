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


  app.Get("/", handlers.HandleIndex)
  log.Fatal(app.Listen(":3000"))
}
