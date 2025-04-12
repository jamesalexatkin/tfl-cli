package presenter

import "errors"

const (
	// ErrNoStationFoundStr is the string for an error returned when no such station is found.
	ErrNoStationFoundStr = "no station found"
)

// ErrNoStationFound is returned when no such station is found.
var ErrNoStationFound = errors.New(ErrNoStationFoundStr)
