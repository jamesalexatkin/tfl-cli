package model

type Departure struct {
	Destination         string
	MinutesUntilArrival int
}

type Platform struct {
	Name       string
	LineName   string
	Color      RoundelColour
	Departures []Departure
}

type Board struct {
	StationName string
	Platforms   []Platform
}
