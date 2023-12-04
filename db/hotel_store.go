package db

import (
	"context"
	"github.com/krinnnout/reserve-get-served/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStore interface {
	InsertHotel(ctx context.Context, hotel *models.Hotel) (*models.Hotel, error)
	Update(ctx context.Context, filter bson.M, update bson.M) error
	GetHotels(ctx context.Context, filter bson.M) ([]*models.Hotel, error)
	GetHotelById(ctx context.Context, id string) (*models.Hotel, error)
}

type MongoHotelStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client) *MongoHotelStore {
	return &MongoHotelStore{
		client: client,
		coll:   client.Database(DBNAME).Collection("hotels"),
	}
}

func (store *MongoHotelStore) InsertHotel(ctx context.Context, hotel *models.Hotel) (*models.Hotel, error) {
	resp, err := store.coll.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.Id = resp.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (store *MongoHotelStore) Update(ctx context.Context, filter bson.M, update bson.M) error {
	_, err := store.coll.UpdateOne(ctx, filter, update)
	return err
}

func (store *MongoHotelStore) GetHotels(ctx context.Context, filter bson.M) ([]*models.Hotel, error) {
	resp, err := store.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var hotels []*models.Hotel
	if err := resp.All(ctx, &hotels); err != nil {
		return nil, err
	}
	return hotels, err

}

func (store *MongoHotelStore) GetHotelById(ctx context.Context, id string) (*models.Hotel, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var hotel *models.Hotel
	if err := store.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&hotel); err != nil {
		return nil, err
	}

	return hotel, nil
}
