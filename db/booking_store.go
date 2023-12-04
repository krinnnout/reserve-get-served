package db

import (
	"context"
	"github.com/krinnnout/reserve-get-served/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingStore interface {
	InsertBooking(ctx context.Context, booking *models.Booking) (*models.Booking, error)
	GetBookings(ctx context.Context, filter bson.M) ([]*models.Booking, error)
	GetBookingById(ctx context.Context, id string) (*models.Booking, error)
}

type MongoBookingStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoBookingStore(client *mongo.Client) *MongoBookingStore {
	return &MongoBookingStore{
		client: client,
		coll:   client.Database(DBNAME).Collection("bookings"),
	}
}

func (store *MongoBookingStore) InsertBooking(ctx context.Context, booking *models.Booking) (*models.Booking, error) {
	resp, err := store.coll.InsertOne(ctx, booking)
	if err != nil {
		return nil, err
	}
	booking.Id = resp.InsertedID.(primitive.ObjectID)
	return booking, nil
}

func (store *MongoBookingStore) GetBookings(ctx context.Context, filter bson.M) ([]*models.Booking, error) {
	resp, err := store.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var bookings []*models.Booking
	if err = resp.All(ctx, &bookings); err != nil {
		return nil, err
	}
	return bookings, nil
}

func (store *MongoBookingStore) GetBookingById(ctx context.Context, id string) (*models.Booking, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var booking models.Booking
	if err = store.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&booking); err != nil {
		return nil, err
	}

	return &booking, nil
}
