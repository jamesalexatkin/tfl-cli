package internal

import (
	"context"
	"fmt"
	"jamesalexatkin/tfl-cli/internal/model"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/jamesalexatkin/tfl-golang"
)

type Service struct {
	TFLClient *tfl.Client
}

func convertLine(s tfl.Status) model.Line {
	l := model.Line{
		Name:   s.Name,
		Status: "Unknown",
	}

	// Handle empty list (shouldn't happen but you never know)
	// if s.LineStatuses == nil {
	// 	return l
	// }

	for _, ls := range s.LineStatuses {
		l.LineStatuses = append(l.LineStatuses, model.LineStatus{
			StatusSeverityDescription: ls.StatusSeverityDescription,
			Reason:                    ls.Reason,
		})
	}

	// // Sort out status
	// l.Status = s.LineStatuses[0].StatusSeverityDescription

	// // Sort out disruption
	// switch l.Status {
	// case "Minor Delays", "Severe Delays", "Part Closure":
	// 	l.Disruption = &s.LineStatuses[0].Reason
	// }

	return l
}

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
			tflStatus.Overground = convertLine(s)
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

// TODO: move this into some render layer
func (s *Service) RenderStatus(ctx context.Context, status *model.TfLStatus) error {
	bold := color.New(color.Bold)

	fmt.Println("┌───────────────────────────")
	bold.Println("LONDON UNDERGROUND")
	renderASCIIRoundel(color.New(color.FgRed), color.New(color.FgBlue))
	renderLine(status.Underground.Bakerloo)
	renderLine(status.Underground.Central)
	renderLine(status.Underground.Circle)
	renderLine(status.Underground.District)
	renderLine(status.Underground.HammersmithAndCity)
	renderLine(status.Underground.Jubilee)
	renderLine(status.Underground.Metropolitan)
	renderLine(status.Underground.Northern)
	renderLine(status.Underground.Piccadilly)
	renderLine(status.Underground.Victoria)
	renderLine(status.Underground.WaterlooAndCity)

	fmt.Println("┌───────────────────────────")
	bold.Println("LONDON OVERGROUND")
	renderASCIIRoundel(color.RGB(239, 123, 16), color.New(color.FgBlue))

	fmt.Println("┌───────────────────────────")
	bold.Println("ELIZABETH LINE")
	renderASCIIRoundel(color.New(color.FgMagenta), color.New(color.FgBlue))
	renderLine(status.ElizabethLine)

	fmt.Println("┌───────────────────────────")
	bold.Println("DLR")
	renderASCIIRoundel(color.New(color.FgCyan), color.New(color.FgBlue))
	renderLine(status.DLR)

	fmt.Printf("(Correct as of %s)\n", status.Time.Format(time.DateTime))

	return nil
}

var box = `┌──────────────────────────────────┐
│               %s                │
├──────────────────────────────────┤
│ %s │
└──────────────────────────────────┘
`

var smallRoundel = `      
       RRRRRRRRR          
    RRRRR     RRRRR      
   RRRR         RRRR   
 BBBBBBBBBBBBBBBBBBBBB 
 BBBBBBBBBBBBBBBBBBBBB 
   RRRR         RRRR   
    RRRRR     RRRRR     
       RRRRRRRRR        

`

var tinyRoundel = `      
      RRRRRR          
    RRR    RRR       
   BBBBBBBBBBBB        
    RRR    RRR 
      RRRRRR         

`

func renderASCIIRoundel(discColour *color.Color, barColour *color.Color) {

	for _, char := range tinyRoundel {
		switch char {
		case 'R':
			discColour.Print("O")
		case 'B':
			barColour.Print("=")
		case '\n':
			fmt.Println("")
		default:
			fmt.Print(" ")
		}

	}
}

func renderLine(line model.Line) {
	// fmt.Println("─────────────────────────────────────────────")
	color.New(color.Bold).Print(line.Name)

	fmt.Print(": \n")

	for _, ls := range line.LineStatuses {
		fmt.Print("\t")

		var disruptionColor color.Color
		switch ls.StatusSeverityDescription {
		case "Good Service":
			disruptionColor = *color.New(color.FgGreen)
		case "Minor Delays":
			disruptionColor = *color.New(color.FgYellow)
		case "Severe Delays":
			disruptionColor = *color.New(color.FgRed)
		case "Part Suspended":
			disruptionColor = *color.New(color.FgMagenta)
		default:
			disruptionColor = *color.New(color.FgWhite)
		}

		disruptionColor.Print(ls.StatusSeverityDescription)

		if ls.Reason != "" {
			fmt.Printf(" - %s", ls.Reason)
		}

		fmt.Printf("\n")
	}
	fmt.Println("─────────────────────────────────────────────")
}

/// STATION

func (s *Service) GetStationArrivals(ctx context.Context) ([]tfl.Prediction, error) {
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

func (s *Service) RenderArrivals(ctx context.Context, arrivals []tfl.Prediction, station string, width int) error {
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

		// Cap at 4 departures
		if len(currentPlatform.Departures) >= 4 {
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
	err := s.RenderDepartureBoard(ctx, board, width)
	if err != nil {
		return err
	}

	return nil
}

func getRoundelStrings(colour model.RoundelColour) []string {
	return []string{
		colour.Disc.Sprint(" ╭───╮"),
		colour.Bar.Sprint("───────"),
		colour.Disc.Sprint(" ╰───╯"),
	}
}

func padRight(str string, length int) string {
	return fmt.Sprintf("%-*s", length, str)
}

func centerText(width int, text string) string {
	padding := (width - len(text)) / 2
	return fmt.Sprintf("%*s%s%*s", padding, "", text, width-padding-len(text), "")
}

func renderPlatform(p model.Platform) []string {
	bold := color.New(color.Bold)

	roundel := getRoundelStrings(p.Color)
	header := fmt.Sprintf("%s │", centerText(28, bold.Sprint(fmt.Sprintf("Platform %s (%s)", p.Name, p.LineName))))

	lines := []string{}
	lines = append(lines, fmt.Sprintf("│ %s%s", roundel[0], "                               │"))
	lines = append(lines, fmt.Sprintf("│ %s %s", roundel[1], header))
	lines = append(lines, fmt.Sprintf("│ %s%s", roundel[2], "                               │"))
	lines = append(lines, "├──────────────────────────────────────┤")

	yellowBold := color.New(color.FgYellow, color.Bold)
	for i, dep := range p.Departures {
		line := fmt.Sprintf("│ %s %s - %dmins", yellowBold.Sprint(strconv.FormatInt(int64(i+1), 10)), dep.Destination, dep.MinutesUntilArrival)
		lines = append(lines, padRight(line, 38)+"│")
	}

	// Padding if fewer than 4 departures
	for i := len(p.Departures); i < 4; i++ {
		lines = append(lines, "│                                      │")
	}

	lines = append(lines, "├──────────────────────────────────────┤")
	return lines
}

func (s *Service) RenderDepartureBoard(ctx context.Context, b model.Board, width int) error {
	// Deal with no data
	if len(b.Platforms) == 0 {
		return ErrNoStationFound
	}

	bold := color.New(color.Bold)
	output := []string{
		"   ╭────────────────────────────────╮",
		fmt.Sprintf("┌──┤ %-30s ├──┐", bold.Sprint(b.StationName)),
		"│  └────────────────────────────────┘  │",
	}

	for _, p := range b.Platforms {
		lines := renderPlatform(p)
		output = append(output, lines...)
	}

	output = append(output, "└──────────────────────────────────────┘\n")

	for _, line := range output {
		fmt.Println(line)
	}

	return nil
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
