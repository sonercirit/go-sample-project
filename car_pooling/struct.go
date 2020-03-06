package car_pooling

type Car struct {
	Id        int
	Seats     int
	FreeSeats *int
	Groups    []*Group `json:"-"`
}

type Group struct {
	Id     int
	People int
	Car    *Car
}

var cars []*Car
var groups []*Group
