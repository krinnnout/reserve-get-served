package main

import (
	"context"
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/krinnnout/reserve-get-served/api"
	"github.com/krinnnout/reserve-get-served/api/middleware"
	"github.com/krinnnout/reserve-get-served/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddress := flag.String("Listen Address", ":5000", "The listen address for the api")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	//Handlers initialization
	var (
		userStore    = db.NewMongoUserStore(client)
		hotelStore   = db.NewMongoHotelStore(client)
		roomStore    = db.NewMongoRoomStore(client, hotelStore)
		bookingStore = db.NewMongoBookingStore(client)
		store        = &db.Store{
			User:    userStore,
			Hotel:   hotelStore,
			Room:    roomStore,
			Booking: bookingStore,
		}
		hotelHandler   = api.NewHotelHandler(store)
		userHandler    = api.NewUserHandler(userStore)
		authHandler    = api.NewAuthHandler(userStore)
		roomHandler    = api.NewRoomHandler(store)
		bookingHandler = api.NewBookingHandler(store)
		app            = fiber.New(config)
		apiv1          = app.Group("/api/v1", middleware.JWTAuthentication(userStore))
		auth           = app.Group("/api")
		admin          = apiv1.Group("/admin", middleware.AdminAuth)
	)

	//auth handlers
	auth.Post("/auth", authHandler.HandleAuthenticate)
	//user handlers
	apiv1.Get("/users", userHandler.HandleGetUsers)
	apiv1.Get("/users/:id", userHandler.HandleGetUser)
	apiv1.Post("/users", userHandler.HandlePostUser)
	apiv1.Delete("/users/:id", userHandler.HandlerDeleteUser)
	apiv1.Put("/users/:id", userHandler.HandlerPutUser)

	//hotel handlers
	apiv1.Get("/hotels", hotelHandler.HandleGetHotels)
	apiv1.Get("/hotels/:id/rooms", hotelHandler.HandleGetRooms)
	apiv1.Get("/hotels/:id", hotelHandler.HandleGetHotel)
	//rooms handlers
	apiv1.Post("/rooms/:id/book", roomHandler.HandleBookRoom)
	apiv1.Get("/rooms", roomHandler.HandleGetRooms)
	//booking handlers
	admin.Get("/bookings", bookingHandler.HandleGetBookings)
	apiv1.Get("/bookings/:id", bookingHandler.HandleGetBooking)
	if err = app.Listen(*listenAddress); err != nil {
		log.Fatal(err)
	}

}
