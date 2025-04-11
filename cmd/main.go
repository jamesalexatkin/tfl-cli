package main

import (
	"context"
	"flag"
	"fmt"
	"jamesalexatkin/tfl-cli/internal/presenter"
	"jamesalexatkin/tfl-cli/internal/service"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/jamesalexatkin/tfl-golang"
	"github.com/mattn/go-isatty"
	"github.com/schachmat/ingo"
	"github.com/urfave/cli/v3"
)

func init() {
	// Turn off color if not running in a proper terminal
	color.NoColor = !isatty.IsTerminal(os.Stdout.Fd())
}

const DefaultWidth = 70

func main() {
	appID := flag.String("app_id", "", "App ID in TfL's portal")
	appKey := flag.String("app_key", "", "App key for TfL's Unified API")

	// read/write config and parse flags
	if err := ingo.Parse("tfl"); err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}

	tflClient := tfl.New(*appID, *appKey)

	service := service.Service{
		TFLClient: tflClient,
	}

	presenter := presenter.Presenter{}

	var statusVerbose bool

	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:  "status",
				Usage: "Show status of all lines",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "verbose",
						Aliases:     []string{"v"},
						Usage:       "Include verbose disruption detail for statuses.",
						Destination: &statusVerbose,
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					status, err := service.GetStatus(ctx)
					if err != nil {
						return err
					}

					err = presenter.RenderStatus(ctx, status, statusVerbose)
					if err != nil {
						return err
					}

					return nil
				},
			},
			{
				Name:  "station",
				Usage: "Show departures from a given station",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					stationName := cmd.Args().First()

					if stationName == "" {
						fmt.Printf("") // TODO: print something here when there's no arg

						// TODO: create NoStationError or MissingArgError
						return nil
					}

					arrivals, err := service.FetchStationArrivalsBoard(ctx, stationName)
					if err != nil {
						return err
					}

					err = presenter.RenderDepartureBoard(ctx, *arrivals, DefaultWidth)
					if err != nil {
						return err
					}

					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
