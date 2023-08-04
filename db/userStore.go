package db

import (
	"github.com/krinnnout/reserve-get-served/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore interface {
	GetUserById(string) (*types.User, error)
}

type MongoUserStore struct {
	UserStore
	client *mongo.Client
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: client,
	}
}
