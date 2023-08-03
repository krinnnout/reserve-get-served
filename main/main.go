package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddress := flag.String("Listen Address", ":5000", "The listen address for the api")
	app := fiber.New()
	apiv1 := app.Group("/api/v1")
	apiv1.Get("/users", handleUsers)
	app.Listen(*listenAddress)

}

func handleUsers(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user1": "Bogdan"})
}
