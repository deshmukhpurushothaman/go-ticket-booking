package main

import (
	"ticket-booking/handlers"
	"ticket-booking/repositories"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "TickerBooking",
		ServerHeader: "Fiber",
	})

	// repositories
	eventRepository := repositories.NewEventRepository(nil)

	// Routing
	server := app.Group("/api")

	// handlers
	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(":8000")
}
