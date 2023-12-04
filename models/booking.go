package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Booking struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserId      primitive.ObjectID `bson:"userId,omitempty" json:"userId,omitempty"`
	RoomId      primitive.ObjectID `bson:"roomId,omitempty" json:"roomId,omitempty"`
	NumOfPeople int                `bson:"numOfPeople,omitempty" json:"numOfPeople,omitempty"`
	FromDate    time.Time          `bson:"fromDate,omitempty" json:"fromDate,omitempty"`
	TillDate    time.Time          `bson:"tillDate,omitempty" json:"tillDate,omitempty"`
}
