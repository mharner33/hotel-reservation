package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mharner33/hotel-reservation/db"
	"github.com/mharner33/hotel-reservation/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://192.168.0.50:27017"

// const dburi = "mongodb://192.168.86.30:27017"
//const dbname = "hotel-reservation"
//const userCollection = "users"

// Create a new fiber instance with custom config
var config = fiber.Config{
	// Override default error handler
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	// Create the Datadog tracer
	// tracer.Start(
	// 	tracer.WithServiceName("hotel-reservation"),
	// 	tracer.WithEnv("dev"),
	// 	tracer.WithServiceVersion("0.0.1"),
	// )
	// create a mongo client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	listenPort := flag.String("listenPort", ":3000", "API server listen address")
	flag.Parse()
	// Fiber instance
	app := fiber.New(config)
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Routes
	userHandler := handlers.NewUserHandler(db.NewMongoUserStore(client))
	v1.Get("/users", userHandler.HandleGetUsers)
	v1.Get("/user/:id", userHandler.HandleGetUser)

	// Start server
	log.Fatal(app.Listen(*listenPort))
	//defer tracer.Stop()
}
