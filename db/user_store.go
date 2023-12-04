package db

import (
	"context"
	"fmt"
	"github.com/krinnnout/reserve-get-served/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const UserCollection = "users"

type Dropper interface {
	Drop(context.Context) error
}

type UserStore interface {
	Dropper
	GetUserById(context.Context, string) (*models.User, error)
	GetUserByEmail(context.Context, string) (*models.User, error)
	GetUsers(context.Context) ([]*models.User, error)
	InsertUser(context.Context, *models.User) (*models.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(context.Context, bson.M, models.ModifiableUserParams) error
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(UserCollection),
	}
}

func (store *MongoUserStore) GetUsers(c context.Context) ([]*models.User, error) {
	var users []*models.User
	cursor, err := store.coll.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(c, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (store *MongoUserStore) GetUserById(ctx context.Context, id string) (*models.User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user models.User
	if err := store.coll.FindOne(ctx, bson.M{"_id": objId}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (store *MongoUserStore) InsertUser(ctx context.Context, user *models.User) (*models.User, error) {
	result, err := store.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (store *MongoUserStore) DeleteUser(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = store.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}

func (store *MongoUserStore) UpdateUser(ctx context.Context, filter bson.M, values models.ModifiableUserParams) error {
	update := bson.M{
		"$set": values.ToBSON(),
	}
	_, err := store.coll.UpdateOne(ctx, filter, update)
	if err != nil {
	}
	return nil

}

func (store *MongoUserStore) Drop(ctx context.Context) error {
	fmt.Println("---dropping user collection")
	return store.coll.Drop(ctx)
}

func (store *MongoUserStore) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := store.coll.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil

}
