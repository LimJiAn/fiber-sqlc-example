package main

import (
	"log"

	"github.com/LimJiAn/fiber-sqlc-example/api/route"
	"github.com/LimJiAn/fiber-sqlc-example/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	database.ConnectDB()
}

func main() {
	app := fiber.New()

	// logger
	app.Use(logger.New())
	route.SetupRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
