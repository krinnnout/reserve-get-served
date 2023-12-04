package main

import (
	"context"
	"github.com/krinnnout/reserve-get-served/db"
	"github.com/krinnnout/reserve-get-served/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	client     *mongo.Client
	roomStore  db.RoomStore
	hotelStore db.HotelStore
	userStore  db.UserStore
	c          = context.Background()
)

func seedHotel(name, location string, rating int) {
	hotel := models.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
		Rating:   rating,
	}
	rooms := []models.Room{
		{
			Size:  "small",
			Price: 99.9,
		},
		{
			Size:  "normal",
			Price: 119.9,
		},
		{
			Size:  "kingsize",
			Price: 199.9,
		},
	}
	insertedHotel, err := hotelStore.InsertHotel(c, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	for _, room := range rooms {
		room.HotelId = insertedHotel.Id
		_, err := roomStore.InsertRoom(c, &room)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func seedUser(fname, lname, password, email string, isAdmin bool) {
	user, err := models.NewUserFromParams(models.UserParams{FirstName: fname, LastName: lname, Email: email, Password: password})
	if err != nil {
		log.Fatal(err)
	}
	user.IsAdmin = isAdmin
	_, err = userStore.InsertUser(c, user)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	seedHotel("Crystal Hotel", "France", 5)
	seedHotel("SpaceY", "Germany", 4)
	seedUser("James", "Jameski", "supersecurepassword", "james@mail.com", false)
	seedUser("admin", "admin", "admin123", "admin@mail.com", true)
}

func init() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Database(db.DBNAME).Drop(c); err != nil {
		log.Fatal(err)
	}
	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
	userStore = db.NewMongoUserStore(client)
}
