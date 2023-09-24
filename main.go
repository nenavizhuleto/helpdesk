package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"application/data"
	"application/handlers"
)

func main() {
	engine := html.New("./www", ".html")

	// For debug
	engine.Reload(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	db, err := data.NewDB()
	if err != nil {
		log.Fatalf("Error initializing SQlite database: %v", err)
	}
	data.DB = db

	app.Static("/", "./public")

	app.Use(handlers.IdentityMiddleware)

	app.Get("/", handlers.HandleIndex)
	app.Get("/sse", handlers.HandleSSE)
	app.Post("/issues/:client_id", handlers.HandleIssues)
	app.Get("/issues/:client_id", handlers.HandleIssues)
	app.Post("/issue/send", handlers.HandleIssueSend)
	app.Get("/test", handlers.HandleTest)
	app.Post("/issue/:client_id/:issue_id", handlers.HandleTestChangeStatus)
	log.Fatal(app.Listen("192.168.31.18:3000"))
}
