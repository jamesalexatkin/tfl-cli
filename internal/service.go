package internal

import (
	"context"
	"fmt"
	"jamesalexatkin/tfl-cli/internal/model"
	"log/slog"
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
	renderRoundel(color.New(color.FgRed), color.New(color.FgBlue))
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
	renderRoundel(color.RGB(239, 123, 16), color.New(color.FgBlue))

	fmt.Println("┌───────────────────────────")
	bold.Println("ELIZABETH LINE")
	renderRoundel(color.New(color.FgMagenta), color.New(color.FgBlue))
	renderLine(status.ElizabethLine)

	fmt.Println("┌───────────────────────────")
	bold.Println("DLR")
	renderRoundel(color.New(color.FgCyan), color.New(color.FgBlue))
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

func renderRoundel(discColour *color.Color, barColour *color.Color) {

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
