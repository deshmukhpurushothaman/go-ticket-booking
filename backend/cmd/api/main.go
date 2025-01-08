package main

import (
	"fmt"
	"ticket-booking/config"
	"ticket-booking/db"
	"ticket-booking/handlers"
	"ticket-booking/repositories"

	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)
	app := fiber.New(fiber.Config{
		AppName:      "TickerBooking",
		ServerHeader: "Fiber",
	})

	// repositories
	eventRepository := repositories.NewEventRepository(db)

	// Routing
	server := app.Group("/api")

	// handlers
	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
