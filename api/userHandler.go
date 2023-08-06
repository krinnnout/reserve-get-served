package api

import (
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

func (userHandler *UserHandler) HandlePostUser(ctx *fiber.Ctx) error {
	var params types.UserParams
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	if errors := params.Validate(); len(errors) > 0 {
		return ctx.JSON(errors)
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	insertedUser, err := userHandler.userStore.InsertUser(ctx.Context(), user)
	if err != nil {
		return err
	}
	return ctx.JSON(insertedUser)
}
