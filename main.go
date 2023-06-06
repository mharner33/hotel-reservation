package main

import (
	"flag"
	"log"

	"github.com/mharner33/hotel-reservation/handlers"

	"github.com/gofiber/fiber/v2"
)

 func main() {
	listenPort := flag.String("listenPort", ":3000", "API server listen address")
	flag.Parse()
	// Fiber instance
	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")


	// Routes
	v1.Get("/users", handlers.HandleGetUsers)
	v1.Get("/user/:id", handlers.HandleGetUser)

	// Start server
	log.Fatal(app.Listen(*listenPort))
}


