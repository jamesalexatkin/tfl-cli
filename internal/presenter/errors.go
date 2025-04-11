package presenter

import "errors"

const (
	ErrNoStationFoundStr = "no station found"
)

var (
	ErrNoStationFound = errors.New(ErrNoStationFoundStr)
)
