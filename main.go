package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	"application/api"
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
	app := fiber.New(fiber.Config{
		ErrorHandler: api.HandleError,
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	app.Post("/register", handlers.HandleRegister)

	apiRouter := app.Group("/api")
	apiRouter.Use(handlers.IdentityMiddlewareDevice)

	// Identity
	apiRouter.Get("/identity", api.GetIdentity)

	// Tasks
	apiRouter.Get("/tasks", api.GetTasks)
	apiRouter.Post("/tasks", api.CreateTask)

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
