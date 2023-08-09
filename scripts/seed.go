package main

import (
	"context"
	"github.com/krinnnout/reserve-get-served/db"
	"github.com/krinnnout/reserve-get-served/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	client     *mongo.Client
	roomStore  db.RoomStore
	hotelStore db.HotelStore
	c          = context.Background()
)

func seedHotel(name, location string, rating int) {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
		Rating:   rating,
	}
	rooms := []types.Room{
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

func main() {
	seedHotel("Crystal Hotel", "France", 5)
	seedHotel("SpaceY", "Germany", 4)
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
}
