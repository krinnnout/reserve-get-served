package api

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/krinnnout/reserve-get-served/db"
	"github.com/krinnnout/reserve-get-served/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (handler *UserHandler) HandleGetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := handler.userStore.GetUserById(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ctx.JSON(map[string]string{"error": "user not found"})
		}
		log.Fatal(err)
	}
	return ctx.JSON(user)
}

func (handler *UserHandler) HandleGetUsers(ctx *fiber.Ctx) error {
	users, err := handler.userStore.GetUsers(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(users)
}

func (handler *UserHandler) HandlePostUser(ctx *fiber.Ctx) error {
	var params types.UserParams
	if err := ctx.BodyParser(&params); err != nil {
		return err
	}
	if errs := params.Validate(); len(errs) > 0 {
		return ctx.JSON(errs)
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	insertedUser, err := handler.userStore.InsertUser(ctx.Context(), user)
	if err != nil {
		return err
	}
	return ctx.JSON(insertedUser)
}

func (handler *UserHandler) HandlerDeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := handler.userStore.DeleteUser(ctx.Context(), id); err != nil {
		return err
	}
	return ctx.JSON(map[string]string{"deleted": id})
}

func (handler *UserHandler) HandlerPutUser(ctx *fiber.Ctx) error {
	var (
		values types.ModifiableUserParams
		id     = ctx.Params("id")
	)
	if err := ctx.BodyParser(&values); err != nil {
		return err
	}
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": oid}
	if err := handler.userStore.UpdateUser(ctx.Context(), filter, values); err != nil {
		return err
	}
	return ctx.JSON(map[string]string{"updated document": id})
}
