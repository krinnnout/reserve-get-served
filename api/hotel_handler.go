package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krinnnout/reserve-get-served/db"
)

type HotelHandler struct {
	roomStore  db.RoomStore
	hotelStore db.HotelStore
}

func NewHotelHandler(rs db.RoomStore, hs db.HotelStore) *HotelHandler {
	return &HotelHandler{
		roomStore:  rs,
		hotelStore: hs,
	}
}

type HotelQueryParams struct {
	Rooms  bool
	Rating int
}

func (handler *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var qparams HotelQueryParams
	if err := c.QueryParser(&qparams); err != nil {
		return err
	}
	hotels, err := handler.hotelStore.GetHotels(c.Context(), nil)
	if err != nil {
		return err
	}
	return c.JSON(hotels)
}
