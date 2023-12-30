package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/krinnnout/reserve-get-served/db"
	"github.com/krinnnout/reserve-get-served/types"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type BookingHandler struct {
	store *db.Store
}

func NewBookingHandler(store *db.Store) *BookingHandler {
	return &BookingHandler{
		store: store,
	}
}

func (handler *BookingHandler) HandleGetBookings(c *fiber.Ctx) error {
	bookings, err := handler.store.Booking.GetBookings(c.Context(), bson.M{})
	if err != nil {
		return err
	}
	return c.JSON(bookings)
}

func (handler *BookingHandler) HandleGetBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := handler.store.Booking.GetBookingById(c.Context(), id)
	if err != nil {
		return err
	}
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return fmt.Errorf("user not found")
	}
	if booking.UserId != user.Id {
		return c.Status(http.StatusUnauthorized).JSON(GenericResponse{Type: "error", Msg: "not authorized"})
	}

	return c.JSON(booking)
}
