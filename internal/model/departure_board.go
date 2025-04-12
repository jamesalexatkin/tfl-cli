package model

// Departure represents a given departing train.
type Departure struct {
	Destination         string
	MinutesUntilArrival int
}

// Platform represents the platform of a station, including departing trains.
type Platform struct {
	Name       string
	LineName   string
	Color      RoundelColour
	Departures []Departure
}

// Board represents a departure board for a particular station.
type Board struct {
	StationName string
	Platforms   []Platform
}
