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
	"application/handlers/v1"
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

	app.Static("/", "./dist")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("dist/index.html")
	})
	app.Get("/system", func(c *fiber.Ctx) error {
		return c.SendFile("dist/system.html")
	})

	apiRouter := app.Group("/api")
	apiRouter.Post("/megaplan/event", api.HandleMegaplanEvent)

	apiV1 := apiRouter.Group("/v1")
	apiV1.Post("/register", handlers.HandleRegister)
	apiV1.Use(handlers.IdentityMiddlewareDevice)
	// Identity
	apiV1.Get("/identity", api.GetIdentity)
	// Users API endpoints
	api.SetUsersRoutes("/users", apiV1)
	api.SetTasksRoutes("/tasks", apiV1)
	api.SetDevicesRoutes("/devices", apiV1)

	apiV2 := apiRouter.Group("/v2")

	apiV2.Get("/identity", api.GetIdentity)
	apiV2.Post("/register", api.Register)
	// Users API endpoints
	api.SetUsersRoutes("/users", apiV2)
	api.SetTasksRoutes("/tasks", apiV2)
	api.SetDevicesRoutes("/devices", apiV2)

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

	data.MustConnectMongo("helpdesk")
}

func initMegaplan() {
	// initializing megaplan connection
	opts := megaplan.NewAuthOpt(util.MustGetEnvVar("MEGAPLAN_USER"), util.MustGetEnvVar("MEGAPLAN_PASSWORD"), util.MustGetEnvVar("MEGAPLAN_RESPONSIBLE"))
	megaplan.MP = megaplan.New(util.MustGetEnvVar("MEGAPLAN_URL"), opts).MustAuthenticateWithPassword(opts)
}
