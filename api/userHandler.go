package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krinnnout/reserve-get-served/types"
)

func HandleGetUsers(contex *fiber.Ctx) error {
	user := types.User{
		"",
		"Boris",
		"Belka",
	}
	return contex.JSON(user)
}
