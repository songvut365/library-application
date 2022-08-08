package main

import (
	"library-app/database"
	"library-app/router"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	// Config
	InitDotENV()
	database.InitDatabase()
	database.InitRedis()
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	// Front-end
	app.Static("/", "./frontend/")

	// Router
	router.SetupRoutes(app)

	// Run
	PORT := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + PORT))
}

func InitDotENV() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
