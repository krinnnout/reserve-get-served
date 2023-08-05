package api

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/krinnnout/reserve-get-served/db"
	"github.com/krinnnout/reserve-get-served/types"
	"log"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (userHandler *UserHandler) HandleGetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	c := context.Background()
	user, err := userHandler.userStore.GetUserById(c, id)
	if err != nil {
		log.Fatal(err)
	}
	return ctx.JSON(user)
}

func (userHandler *UserHandler) HandleGetUsers(ctx *fiber.Ctx) error {
	user := types.User{
		FirstName: "Boris",
		LastName:  "Belka",
	}
	return ctx.JSON(user)
}
