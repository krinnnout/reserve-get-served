package db

import (
	"context"
	"github.com/krinnnout/reserve-get-served/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomStore interface {
	InsertRoom(ctx context.Context, hotel *types.Room) (*types.Room, error)
}

type MongoRoomStore struct {
	client *mongo.Client
	coll   *mongo.Collection

	HotelStore
}

func NewMongoRoomStore(client *mongo.Client, hotelStore HotelStore) *MongoRoomStore {
	return &MongoRoomStore{
		client:     client,
		coll:       client.Database(DBNAME).Collection("rooms"),
		HotelStore: hotelStore,
	}
}

func (store *MongoRoomStore) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	resp, err := store.coll.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	room.Id = resp.InsertedID.(primitive.ObjectID)

	filter := bson.M{"_id": room.HotelId}
	update := bson.M{"$push": bson.M{"rooms": room.Id}}
	if err := store.HotelStore.Update(ctx, filter, update); err != nil {
		return nil, err
	}

	return room, nil
}
