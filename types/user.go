package types

type User struct {
	Id        string `bson:"_id" json:"id,omitempty"`
	FirstName string `bson:"_firstName" json:"firstName"`
	LastName  string `bson:"_lastName" json:"lastName"`
}
