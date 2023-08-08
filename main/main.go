package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/krinnnout/reserve-get-served/api"
	"github.com/krinnnout/reserve-get-served/db"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbUri          = "mongodb://localhost:27017"
	dbName         = "reserve-get-served"
	userCollection = "users"
)

var config = fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddress := flag.String("Listen Address", ":5000", "The listen address for the api")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(client)
	//Handlers initialization
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client, dbName))

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/users", userHandler.HandleGetUsers)
	apiv1.Get("/users/:id", userHandler.HandleGetUser)
	apiv1.Post("/users", userHandler.HandlePostUser)
	apiv1.Delete("/users/:id", userHandler.HandlerDeleteUser)
	apiv1.Put("/users/:id", userHandler.HandlerPutUser)

	if err = app.Listen(*listenAddress); err != nil {
		log.Fatal(err)
	}

}
