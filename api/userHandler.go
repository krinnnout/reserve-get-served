package api

import "github.com/gofiber/fiber/v2"

func HandleGetUsers(contex *fiber.Ctx) error {

	return contex.JSON("James")
}
