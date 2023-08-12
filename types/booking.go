package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Booking struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserId      primitive.ObjectID `bson:"user_id,omitempty" json:"userId,omitempty"`
	RoomId      primitive.ObjectID `bson:"room_id,omitempty" json:"roomId,omitempty"`
	NumOfPeople int                `bson:"num_of_people,omitempty" json:"numOfPeople,omitempty"`
	FromDate    time.Time          `bson:"from_Date,omitempty" json:"fromDate,omitempty"`
	TillDate    time.Time          `bson:"till_Date,omitempty" json:"tillDate,omitempty"`
}
