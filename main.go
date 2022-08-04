package main

import (
	"library-app/database"
	"library-app/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	// Config
	database.InitDatabase()
	database.InitRedis()

	// Router
	router.SetupRoutes(app)

	// Run
	log.Fatal(app.Listen(":3000"))
}
