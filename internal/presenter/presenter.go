package presenter

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/jamesalexatkin/tfl-cli/internal/config"
	"github.com/jamesalexatkin/tfl-cli/internal/model"
)

// Presenter is used to render data to the terminal.
type Presenter struct{}

// RenderConfig renders the current app configuration.
func (p *Presenter) RenderConfig(ctx context.Context, cfg *config.Config) {
	bold := color.New(color.Bold)
	yellow := color.New(color.FgYellow)

	fmt.Println("Current configuration:")

	fmt.Printf("%s: %s\n", yellow.Sprint(config.AppIDConfigKey), bold.Sprint(cfg.AppID))
	fmt.Printf("%s: %s\n", yellow.Sprint(config.AppKeyConfigKey), bold.Sprint(cfg.AppKey))
	fmt.Printf("%s: %s\n", yellow.Sprint(config.DepartureBoardWidthConfigKey), bold.Sprint(cfg.DepartureBoardWidth))
	fmt.Printf("%s: %s\n", yellow.Sprint(config.HomeStationConfigKey), bold.Sprint(cfg.HomeStation))
	fmt.Printf("%s: %s\n", yellow.Sprint(config.WorkStationConfigKey), bold.Sprint(cfg.WorkStation))
}

// RenderStatus renders the current TfL status for all lines.
//
//nolint:funlen
func (p *Presenter) RenderStatus(ctx context.Context, status *model.TfLStatus, verbose bool) {
	bold := color.New(color.Bold)
	italic := color.New(color.Italic)

	fmt.Println("╭───────────────────────────╮             ╭───────────────────────────╮             ╭───────────────────────────╮             ╭───────────────────────────╮")
	fmt.Printf("│ %s        │             │ %s         │             │ %s            │             │ %s                       │\n",
		bold.Sprint("London Underground"),
		bold.Sprint("London Overground"),
		bold.Sprint("Elizabeth Line"),
		bold.Sprint("DLR"))
	fmt.Println("├───────────────────────────┴─────────────┼───────────────────────────┴─────────────┼───────────────────────────┴─────────────┼───────────────────────────┴─────────────┐")

	undergroundRoundel := model.CreateRoundelColourFromLineName("tube")
	overgroundRoundel := model.CreateRoundelColourFromLineName("overground")
	elizabethLineRoundel := model.CreateRoundelColourFromLineName("elizabeth-line")
	dlrRoundel := model.CreateRoundelColourFromLineName("dlr")

	roundelTop := "╭───╮"
	roundelMiddle := "───────"
	roundelBottom := "╰───╯"

	fmt.Printf("│  %s                                  │  %s                                  │  %s                                  │  %s                                  │\n",
		undergroundRoundel.Disc.Sprint(roundelTop),
		overgroundRoundel.Disc.Sprint(roundelTop),
		elizabethLineRoundel.Disc.Sprint(roundelTop),
		dlrRoundel.Disc.Sprint(roundelTop),
	)
	fmt.Printf("│ %s                                 │ %s                                 │ %s                                 │ %s                                 │\n",
		undergroundRoundel.Bar.Sprint(roundelMiddle),
		overgroundRoundel.Bar.Sprint(roundelMiddle),
		elizabethLineRoundel.Bar.Sprint(roundelMiddle),
		dlrRoundel.Bar.Sprint(roundelMiddle),
	)
	fmt.Printf("│  %s                                  │  %s                                  │  %s                                  │  %s                                  │\n",
		undergroundRoundel.Disc.Sprint(roundelBottom),
		overgroundRoundel.Disc.Sprint(roundelBottom),
		elizabethLineRoundel.Disc.Sprint(roundelBottom),
		dlrRoundel.Disc.Sprint(roundelBottom),
	)

	emptyPadding := "                                       "
	boxWidth := 39

	fmt.Printf("│ %s │ %s │ %s │ %s │\n",
		padRight(getLine(status.Underground.Bakerloo, verbose), boxWidth),
		padRight(getLine(status.Overground.Liberty, verbose), boxWidth),
		padRight(getLine(status.ElizabethLine, verbose), boxWidth),
		padRight(getLine(status.DLR, verbose), boxWidth),
	)
	fmt.Printf("│ %s │ %s │ %s │ %s │\n",
		padRight(getLine(status.Underground.Central, verbose), boxWidth),
		padRight(getLine(status.Overground.Lioness, verbose), boxWidth),
		emptyPadding,
		emptyPadding,
	)
	fmt.Printf("│ %s │ %s │ %s │ %s │\n",
		padRight(getLine(status.Underground.Circle, verbose), boxWidth),
		padRight(getLine(status.Overground.Mildmay, verbose), boxWidth),
		emptyPadding,
		emptyPadding,
	)
	fmt.Printf("│ %s │ %s │ %s │ %s │\n",
		padRight(getLine(status.Underground.District, verbose), boxWidth),
		padRight(getLine(status.Overground.Suffragette, verbose), boxWidth),
		emptyPadding,
		emptyPadding,
	)
	fmt.Printf("│ %s │ %s │ %s │ %s │\n",
		padRight(getLine(status.Underground.HammersmithAndCity, verbose), boxWidth),
		padRight(getLine(status.Overground.Weaver, verbose), boxWidth),
		emptyPadding,
		emptyPadding,
	)
	fmt.Printf("│ %s │ %s │ %s │ %s │\n",
		padRight(getLine(status.Underground.Jubilee, verbose), boxWidth),
		padRight(getLine(status.Overground.Windrush, verbose), boxWidth),
		emptyPadding,
		emptyPadding,
	)
	fmt.Printf("│ %s │ %s │ %s │ %s │\n",
		padRight(getLine(status.Underground.Metropolitan, verbose), boxWidth),
		emptyPadding,
		emptyPadding,
		emptyPadding,
	)
	fmt.Printf("│ %s │ %s │ %s │ %s │\n",
		padRight(getLine(status.Underground.Northern, verbose), boxWidth),
		emptyPadding,
		emptyPadding,
		emptyPadding,
	)
	fmt.Printf("│ %s │ %s │ %s │ %s │\n",
		padRight(getLine(status.Underground.Piccadilly, verbose), boxWidth),
		emptyPadding,
		emptyPadding,
		emptyPadding,
	)
	fmt.Printf("│ %s │ %s │ %s │ %s │\n",
		padRight(getLine(status.Underground.Victoria, verbose), boxWidth),
		emptyPadding,
		emptyPadding,
		emptyPadding,
	)
	fmt.Printf("│ %s │ %s │ %s │ %s │\n",
		padRight(getLine(status.Underground.WaterlooAndCity, verbose), boxWidth),
		emptyPadding,
		emptyPadding,
		emptyPadding,
	)

	fmt.Println("└─────────────────────────────────────────┴─────────────────────────────────────────┴─────────────────────────────────────────┴─────────────────────────────────────────┘")

	italic.Printf("(Correct as of %s)\n", status.Time.Format(time.DateTime))
	fmt.Println("")
}

func getLine(line model.Line, verbose bool) string {
	roundelColour := model.CreateRoundelColourFromLineName(line.Name)

	disruption := ""
	for _, ls := range line.LineStatuses {
		var disruptionColor color.Color
		switch ls.StatusSeverityDescription {
		case "Good Service":
			disruptionColor = *color.New(color.FgGreen)
		case "Minor Delays":
			disruptionColor = *color.New(color.FgYellow)
		case "Severe Delays":
			disruptionColor = *color.New(color.FgRed)
		case "Reduced Service", "Part Suspended":
			disruptionColor = *color.New(color.FgMagenta)
		case "Part Closure", "Planned Closure", "Service Closed":
			disruptionColor = *color.New(color.FgBlue)
		default:
			disruptionColor = *color.New(color.FgWhite)
		}

		disruption = disruptionColor.Sprint(ls.StatusSeverityDescription)

		// if ls.Reason != "" && verbose {
		// 	fmt.Printf(" - %s", ls.Reason)
		// }
	}

	lineContent := fmt.Sprintf("%s %s: %s",
		roundelColour.Disc.Sprint("█"),
		color.New(color.Bold).Sprint(line.Name),
		disruption,
	)

	return lineContent
}

// RenderDepartureBoard renders a departure board.
func (p *Presenter) RenderDepartureBoard(ctx context.Context, b model.Board, width int) error {
	// Deal with no data
	if len(b.Platforms) == 0 {
		return ErrNoStationFound
	}

	bold := color.New(color.Bold)

	output := []string{
		drawFrameLine("   ╭", "╮   ", '─', width),
		"┌──┤" + centerText(width-realLen("┌──┤"+"├──┐"), bold.Sprint(b.StationName)) + "├──┐",
		drawFrameLine("│  └", "┘  │", '─', width),
	}

	for _, p := range b.Platforms {
		lines := getPlatformStrings(p, width)
		output = append(output, lines...)
	}

	output = append(output,
		drawFrameLine("└", "┘", '─', width))

	for _, line := range output {
		// fmt.Println(line + fmt.Sprintf(" %d", realLen(line)))
		fmt.Println(line)
	}

	return nil
}

func drawFrameLine(leftPiece string, rightPiece string, centreChar rune, width int) string {
	return leftPiece + strings.Repeat(string(centreChar), width-realLen(leftPiece+rightPiece)) + rightPiece
}

// 2. to properly count special Unicode chars which use multiple bytes (e.g. box-drawing chars) as only 1 character.
func realLen(s string) int {
	ansiRegexp := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	ansiEscapedString := ansiRegexp.ReplaceAllString(s, "")

	return len([]rune(ansiEscapedString))
}

func getRoundelStrings(colour model.RoundelColour) []string {
	return []string{
		colour.Disc.Sprint(" ╭───╮"),
		colour.Bar.Sprint("───────"),
		colour.Disc.Sprint(" ╰───╯"),
	}
}

func generatePadding(paddingChar rune, length int) string {
	if length < 0 {
		return ""
	}

	return strings.Repeat(string(paddingChar), length)
}

func padRight(str string, length int) string {
	paddingAmount := length - realLen(str)

	return str + generatePadding(' ', paddingAmount)
}

func centerText(width int, text string) string {
	leftPadding := (width - realLen(text)) / 2
	rightPadding := width - leftPadding - realLen(text)

	return fmt.Sprintf("%*s%s%*s", leftPadding, "", text, rightPadding, "")
}

func getPlatformStrings(p model.Platform, width int) []string {
	bold := color.New(color.Bold)

	roundel := getRoundelStrings(p.Color)

	// Default of 'Platform X'
	platformName := "Platform " + p.Name
	if strings.Contains(p.Name, "Platform") {
		// Some platforms are formatted differently so we just take the raw name if so
		// e.g. 'Eastbound - Platform 3' at North Acton
		platformName = p.Name
	}
	header := bold.Sprint(fmt.Sprintf("%s (%s)", platformName, p.LineName)) + "  "

	lines := []string{}
	roundelTopContent := "│ " + roundel[0]
	lines = append(lines, fmt.Sprintf("%s%s│", roundelTopContent, generatePadding(' ', width-realLen(roundelTopContent)-1)))
	roundelMiddleContent := fmt.Sprintf("│ %s %s", roundel[1], header)
	lines = append(lines, fmt.Sprintf("%s%s│", roundelMiddleContent, generatePadding(' ', width-realLen(roundelMiddleContent)-1)))
	roundelBottomContent := "│ " + roundel[2]
	lines = append(lines, fmt.Sprintf("%s%s│", roundelBottomContent, generatePadding(' ', width-realLen(roundelBottomContent)-1)))
	lines = append(lines, drawFrameLine("├", "┤", '─', width))

	yellowBold := color.New(color.FgYellow, color.Bold)
	for i, dep := range p.Departures {
		content := fmt.Sprintf("│ %s %s - %dmins", yellowBold.Sprint(strconv.FormatInt(int64(i+1), 10)), dep.Destination, dep.MinutesUntilArrival)
		line := fmt.Sprintf("%s%s│", content, generatePadding(' ', width-realLen(content)-1))
		lines = append(lines, line)
	}

	// Padding if fewer than 4 departures
	for i := len(p.Departures); i < 4; i++ {
		lines = append(lines, drawFrameLine("│", "│", ' ', width))
	}

	lines = append(lines, drawFrameLine("├", "┤", '─', width))

	return lines
}
