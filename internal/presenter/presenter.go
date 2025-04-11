package presenter

import (
	"context"
	"fmt"
	"jamesalexatkin/tfl-cli/internal/model"
	"regexp"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type Presenter struct {
}

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

// realLen calculates the true length of a string as judged per character.
// This is needed for two reasons:
// 1. to escape ANSI strings (used for colouring);
// 2. to properly count special Unicode chars which use multiple bytes (e.g. box-drawing chars) as only 1 character
func realLen(s string) int {
	var ansiRegexp = regexp.MustCompile(`\x1b\[[0-9;]*m`)
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
	return fmt.Sprintf("%-*s", length, str)
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
	header := fmt.Sprintf("%s  ", bold.Sprint(fmt.Sprintf("%s (%s)", platformName, p.LineName)))

	lines := []string{}
	roundelTopContent := fmt.Sprintf("│ %s", roundel[0])
	lines = append(lines, fmt.Sprintf("%s%s│", roundelTopContent, generatePadding(' ', width-realLen(roundelTopContent)-1)))
	// lines = append(lines, fmt.Sprintf("│ %s%s", roundel[0], "                               │"))
	// lines = append(lines, fmt.Sprintf("│ %s %s", roundel[1], header))
	roundelMiddleContent := fmt.Sprintf("│ %s %s", roundel[1], header)
	lines = append(lines, fmt.Sprintf("%s%s│", roundelMiddleContent, generatePadding(' ', width-realLen(roundelMiddleContent)-1)))
	// lines = append(lines, fmt.Sprintf("│ %s%s", roundel[2], "                               │"))
	roundelBottomContent := fmt.Sprintf("│ %s", roundel[2])
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
