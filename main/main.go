package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/krinnnout/reserve-get-served/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gofiber/fiber/v2"
	"github.com/krinnnout/reserve-get-served/api"
)

const dbUri = "mongodb://localhost:27017"
const dbName = "reserve-get-served"
const userCollection = "users"

func main() {
	listenAddress := flag.String("Listen Address", ":5000", "The listen address for the api")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(client)
	//Handlers initialization
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/users", userHandler.HandleGetUsers)
	apiv1.Get("/users/:id", userHandler.HandleGetUser)

	if err = app.Listen(*listenAddress); err != nil {
		log.Fatal(err)
	}

}
