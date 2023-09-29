package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"

	"application/data"
	"application/handlers"
	"application/megaplan"
	"application/util"
)

func main() {
	initEnv()
	initDb()
	initMegaplan()

	// initializing fiber application
	engine := html.New("./www", ".html")
	engine.Reload(true)
	app := fiber.New(fiber.Config{
		Views:             engine,
		PassLocalsToViews: true,
	})
	app.Static("/", "./public")

	app.Use(logger.New())

	app.Get("/auth", handlers.HandleAuth)
	app.Post("/auth/signup", handlers.HandleSignup)
	app.Use(handlers.IdentityMiddlewareDevice)
	app.Get("/", handlers.HandleIndex)
	app.Get("/metrics", monitor.New())
	app.Get("/chat", handlers.HandleChat)
	app.Get("/task/new", handlers.HandleGetTaskNew)
	app.Post("/task/new", handlers.HandlePostTaskNew)

	// app.Use("/ws", handlers.HandleWebSocket)
	// app.Get("/ws/:id", websocket.New(handlers.HandleGetWebSocket))

	log.Fatal(app.Listen(":3000"))
}

func initEnv() {
	// initializing environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func initDb() {
	// initializing database connection
	if db, err := data.NewDB(); err != nil {
		log.Fatalf("Error initializing SQlite database: %v", err)
	} else {
		data.DB = db
	}
}

func initMegaplan() {
	// initializing megaplan connection
	opts := megaplan.NewAuthOpt(util.MustGetEnvVar("MEGAPLAN_USER"), util.MustGetEnvVar("MEGAPLAN_PASSWORD"))
	megaplan.MP = megaplan.New(util.MustGetEnvVar("MEGAPLAN_URL"), opts).MustAuthenticateWithPassword(opts)
}
