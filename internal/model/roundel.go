package model

import "github.com/fatih/color"

type RoundelColour struct {
	Disc *color.Color
	Bar  *color.Color
}

func CreateRoundelColourFromLineName(lineName string) RoundelColour {
	switch lineName {
	// Tube
	case "Bakerloo":
		return RoundelColour{
			Disc: color.RGB(178, 99, 0),
			Bar:  color.RGB(178, 99, 0),
		}
	case "Central":
		return RoundelColour{
			Disc: color.New(color.FgRed),
			Bar:  color.New(color.FgRed),
		}
	case "Circle":
		return RoundelColour{
			Disc: color.New(color.FgYellow),
			Bar:  color.New(color.FgYellow),
		}
	case "District":
		return RoundelColour{
			Disc: color.New(color.FgGreen),
			Bar:  color.New(color.FgGreen),
		}
	case "Hammersmith & City":
		return RoundelColour{
			Disc: color.RGB(244, 169, 190),
			Bar:  color.RGB(244, 169, 190),
		}
	case "Jubilee":
		return RoundelColour{
			Disc: color.New(color.FgWhite),
			Bar:  color.New(color.FgWhite),
		}
	case "Metropolitan":
		return RoundelColour{
			Disc: color.New(color.FgMagenta),
			Bar:  color.New(color.FgMagenta),
		}
	case "Northern":
		return RoundelColour{
			Disc: color.New(color.FgBlack),
			Bar:  color.New(color.FgBlack),
		}
	case "Piccadilly":
		return RoundelColour{
			Disc: color.New(color.FgBlue),
			Bar:  color.New(color.FgBlue),
		}
	case "Victoria":
		return RoundelColour{
			Disc: color.RGB(0, 152, 216),
			Bar:  color.RGB(0, 152, 216),
		}
	case "Waterloo & City":
		return RoundelColour{
			Disc: color.New(color.FgCyan),
			Bar:  color.New(color.FgCyan),
		}
	// Overground
	case "Liberty":
		return RoundelColour{
			Disc: color.New(color.FgWhite),
			Bar:  color.New(color.FgWhite),
		}
	case "Lioness":
		return RoundelColour{
			Disc: color.New(color.FgYellow),
			Bar:  color.New(color.FgYellow),
		}
	case "Mildmay":
		return RoundelColour{
			Disc: color.New(color.FgBlue),
			Bar:  color.New(color.FgBlue),
		}
	case "Suffragette":
		return RoundelColour{
			Disc: color.New(color.FgGreen),
			Bar:  color.New(color.FgGreen),
		}
	case "Weaver":
		return RoundelColour{
			Disc: color.New(color.FgMagenta),
			Bar:  color.New(color.FgMagenta),
		}
	case "Windrush":
		return RoundelColour{
			Disc: color.New(color.FgRed),
			Bar:  color.New(color.FgRed),
		}
	// Elizabeth Line
	case "Elizabeth Line":
		return RoundelColour{
			Disc: color.New(color.FgMagenta),
			Bar:  color.New(color.FgBlue),
		}
	// DLR
	case "DLR":
		return RoundelColour{
			Disc: color.New(color.FgCyan),
			Bar:  color.New(color.FgBlue),
		}
	default:
		return RoundelColour{
			Disc: color.New(color.FgWhite),
			Bar:  color.New(color.FgWhite),
		}
	}
}
