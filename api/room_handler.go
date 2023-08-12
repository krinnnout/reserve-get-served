package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/krinnnout/reserve-get-served/db"
	"github.com/krinnnout/reserve-get-served/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type BookRoomParams struct {
	FromDate    time.Time `json:"fromDate"`
	TillDate    time.Time `json:"tillDate"`
	NumOfPeople int       `json:"numOfPeople"`
}
type RoomHandler struct {
	store *db.Store
}

func NewRoomHandler(store *db.Store) *RoomHandler {
	return &RoomHandler{
		store: store,
	}
}

func (handler *RoomHandler) HandleBookRoom(c *fiber.Ctx) error {
	var params BookRoomParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	roomId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return err
	}
	user, ok := c.Context().Value("user").(*types.User)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(GenericResponse{Type: "error", Msg: "internal server error"})

	}

	booking := types.Booking{
		UserId:      user.Id,
		RoomId:      roomId,
		FromDate:    params.FromDate,
		TillDate:    params.TillDate,
		NumOfPeople: params.NumOfPeople,
	}
	fmt.Printf("%+v\n", booking)
	return nil
}
