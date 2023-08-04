package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/krinnnout/reserve-get-served/api"
)

func main() {
	listenAddress := flag.String("Listen Address", ":5000", "The listen address for the api")
	app := fiber.New()
	apiv1 := app.Group("/api/v1")
	apiv1.Get("/users", api.HandleGetUsers)
	app.Listen(*listenAddress)

}
