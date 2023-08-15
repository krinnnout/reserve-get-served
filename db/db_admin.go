package db

const (
	DBNAME     = "reserve-get-served"
	TestDBNAME = "reserve-get-served-test"
	DBURI      = "mongodb://localhost:27017"
)

type Store struct {
	User    UserStore
	Hotel   HotelStore
	Room    RoomStore
	Booking BookingStore
}
