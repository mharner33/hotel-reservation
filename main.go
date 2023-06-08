package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/mharner33/hotel-reservation/handlers"
	"github.com/mharner33/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gofiber/fiber/v2"
)
const dburi = "mongodb://192.168.86.30:27017"
const dbname = "hotel-reservation"
const userCollection = "users"

 func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	coll := client.Database(dbname).Collection(userCollection)
	// user := types.User{
	// 	FirstName: "Mike",
	// 	LastName: "Jones",
	// }
	// res, err := coll.InsertOne(ctx, user)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(res.InsertedID)

	var mike types.User
	err = coll.FindOne(ctx, bson.M{}).Decode(&mike)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mike)
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


