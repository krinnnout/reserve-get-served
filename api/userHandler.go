package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krinnnout/reserve-get-served/db"
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
	user, err := userHandler.userStore.GetUserById(ctx.Context(), id)
	if err != nil {
		log.Fatal(err)
	}
	return ctx.JSON(user)
}

func (userHandler *UserHandler) HandleGetUsers(ctx *fiber.Ctx) error {
	users, err := userHandler.userStore.GetUsers(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(users)
}

func (UserHandler *UserHandler) HandlePostUser(ctx *fiber.Ctx) error {
	return nil
}
