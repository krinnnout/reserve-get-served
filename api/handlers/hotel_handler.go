package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krinnnout/reserve-get-served/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	store *db.Store
}

func NewHotelHandler(s *db.Store) *HotelHandler {
	return &HotelHandler{
		store: s,
	}
}

func (handler *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	hotels, err := handler.store.Hotel.GetHotels(c.Context(), nil)
	if err != nil {
		return err
	}
	return c.JSON(hotels)
}

func (handler *HotelHandler) HandleGetHotel(c *fiber.Ctx) error {
	id := c.Params("id")
	hotel, err := handler.store.Hotel.GetHotelById(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(hotel)
}

func (handler *HotelHandler) HandleGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	rooms, err := handler.store.Room.GetRooms(c.Context(), bson.M{"hotelId": oid})
	if err != nil {
		return err
	}
	return c.JSON(rooms)
}
