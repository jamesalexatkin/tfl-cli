package service

import (
	"context"
	"log/slog"
	"strings"
	"time"

	"github.com/jamesalexatkin/tfl-cli/internal/model"
	tfl "github.com/jamesalexatkin/tfl-golang"
)

// Service is used to fetch and marshal data from TfL's API.
type Service struct {
	TFLClient *tfl.Client
}

func convertLine(s tfl.Status) model.Line {
	l := model.Line{
		Name:   s.Name,
		Status: "Unknown",
	}

	for _, ls := range s.LineStatuses {
		l.LineStatuses = append(l.LineStatuses, model.LineStatus{
			StatusSeverityDescription: ls.StatusSeverityDescription,
			Reason:                    ls.Reason,
		})
	}

	return l
}

// GetStatus fetches the status for all lines.
//
//nolint:cyclop,funlen
func (s *Service) GetStatus(ctx context.Context) (*model.TfLStatus, error) {
	statuses, err := s.TFLClient.GetLineStatusByMode(ctx, []string{"tube", "overground", "elizabeth-line", "dlr"})
	if err != nil {
		return nil, err
	}

	tflStatus := model.TfLStatus{
		Time: time.Now(),
	}

	for _, s := range statuses {
		switch s.ModeName {
		case "tube":
			switch s.Name {
			case "Bakerloo":
				tflStatus.Underground.Bakerloo = convertLine(s)
			case "Central":
				tflStatus.Underground.Central = convertLine(s)
			case "Circle":
				tflStatus.Underground.Circle = convertLine(s)
			case "District":
				tflStatus.Underground.District = convertLine(s)
			case "Hammersmith & City":
				tflStatus.Underground.HammersmithAndCity = convertLine(s)
			case "Jubilee":
				tflStatus.Underground.Jubilee = convertLine(s)
			case "Metropolitan":
				tflStatus.Underground.Metropolitan = convertLine(s)
			case "Northern":
				tflStatus.Underground.Northern = convertLine(s)
			case "Piccadilly":
				tflStatus.Underground.Piccadilly = convertLine(s)
			case "Victoria":
				tflStatus.Underground.Victoria = convertLine(s)
			case "Waterloo & City":
				tflStatus.Underground.WaterlooAndCity = convertLine(s)
			}
		case "overground":
			switch s.Name {
			case "Liberty":
				tflStatus.Overground.Liberty = convertLine(s)
			case "Lioness":
				tflStatus.Overground.Lioness = convertLine(s)
			case "Mildmay":
				tflStatus.Overground.Mildmay = convertLine(s)
			case "Suffragette":
				tflStatus.Overground.Suffragette = convertLine(s)
			case "Weaver":
				tflStatus.Overground.Weaver = convertLine(s)
			case "Windrush":
				tflStatus.Overground.Windrush = convertLine(s)
			}
		case "dlr":
			tflStatus.DLR = convertLine(s)
		case "elizabeth-line":
			tflStatus.ElizabethLine = convertLine(s)
		default:
			slog.Info("Unknown mode: " + s.ModeName)
		}
	}

	return &tflStatus, nil
}

/// STATION

// FetchStationArrivalsBoard fetches the arrivals for a station and formats it into a board.
func (s *Service) FetchStationArrivalsBoard(ctx context.Context, station string, numDepartures int) (*model.Board, error) {
	arrivals, err := s.fetchArrivals(ctx)
	if err != nil {
		return nil, err
	}

	board := s.convertArrivalsToBoard(ctx, station, arrivals, numDepartures)

	return &board, nil
}

func (s *Service) fetchArrivals(ctx context.Context) ([]tfl.Prediction, error) {
	// TODO: Get StopPoint using /StopPoint/Search/{query}

	totalArrivals := []tfl.Prediction{}

	undergroundArrivals, err := s.TFLClient.GetArrivalPredictionsForMode(ctx, "tube", 10)
	if err != nil {
		return nil, err
	}
	totalArrivals = append(totalArrivals, undergroundArrivals...)

	overgroundArrivals, err := s.TFLClient.GetArrivalPredictionsForMode(ctx, "overground", 10)
	if err != nil {
		return nil, err
	}
	totalArrivals = append(totalArrivals, overgroundArrivals...)

	elizabethLineArrivals, err := s.TFLClient.GetArrivalPredictionsForMode(ctx, "elizabeth-line", 10)
	if err != nil {
		return nil, err
	}
	totalArrivals = append(totalArrivals, elizabethLineArrivals...)

	return totalArrivals, nil
}

func (s *Service) convertArrivalsToBoard(ctx context.Context, station string, arrivals []tfl.Prediction, numDepartures int) model.Board {
	board := model.Board{
		StationName: station,
	}

	platforms := map[string]model.Platform{}

	for _, a := range arrivals {
		if stripStationName(a.StationName) != station {
			continue
		}

		currentPlatform, ok := platforms[a.PlatformName]
		if !ok {
			currentPlatform = model.Platform{
				Name:       a.PlatformName,
				LineName:   a.LineName,
				Color:      model.CreateRoundelColourFromLineName(a.LineName),
				Departures: []model.Departure{},
			}
		}

		// Cap at number of departures config setting
		if len(currentPlatform.Departures) >= numDepartures {
			continue
		}

		d := time.Duration(a.TimeToStation) * time.Second

		currentPlatform.Departures = append(currentPlatform.Departures, model.Departure{
			Destination:         stripStationName(a.DestinationName),
			MinutesUntilArrival: int(d.Minutes()),
		})
		platforms[a.PlatformName] = currentPlatform
	}

	platformsSlice := []model.Platform{}
	for _, p := range platforms {
		platformsSlice = append(platformsSlice, p)
	}

	board.Platforms = platformsSlice

	return board
}

func stripStationName(station string) string {
	stringsToStrip := []string{
		" Underground Station",
		" Rail Station",
		" (Berks)",
	}

	strippedStation := station
	for _, s := range stringsToStrip {
		strippedStation = strings.Split(strippedStation, s)[0]
	}

	return strippedStation
}
