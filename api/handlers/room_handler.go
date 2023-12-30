package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/krinnnout/reserve-get-served/db"
	"github.com/krinnnout/reserve-get-served/types"
	"go.mongodb.org/mongo-driver/bson"
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
	insertedBooking, err := handler.store.Booking.InsertBooking(c.Context(), &booking)
	if err != nil {
		return err
	}
	ok, err = handler.isRoomAvailableForBooking(c, roomId, params)
	if err != nil {
		return err
	}
	if !ok {
		return c.Status(http.StatusBadRequest).JSON(GenericResponse{Type: "error", Msg: fmt.Sprintf("room %s already booked", roomId.String())})

	}
	return c.JSON(insertedBooking)
}

func (params BookRoomParams) validate() error {
	now := time.Now()
	if now.After(params.FromDate) || now.After(params.TillDate) {
		return fmt.Errorf("cannot book room in the past")
	}
	return nil
}

func (handler *RoomHandler) isRoomAvailableForBooking(c *fiber.Ctx, roomId primitive.ObjectID, params BookRoomParams) (bool, error) {
	filter := bson.M{
		"roomId": roomId,
		"fromDate": bson.M{
			"$gte": params.FromDate,
		},
		"tillDate": bson.M{
			"$lte": params.TillDate,
		},
	}
	bookings, err := handler.store.Booking.GetBookings(c.Context(), filter)
	if err != nil {
		return false, err
	}
	fmt.Println(bookings)

	return len(bookings) == 0, nil
}

func (handler *RoomHandler) HandleGetRooms(c *fiber.Ctx) error {
	rooms, err := handler.store.Room.GetRooms(c.Context(), bson.M{})
	if err != nil {
		return err
	}
	return c.JSON(rooms)
}
