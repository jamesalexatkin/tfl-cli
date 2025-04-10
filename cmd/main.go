package main

import (
	"context"
	"flag"
	"fmt"
	"jamesalexatkin/tfl-cli/internal"
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

	service := internal.Service{
		TFLClient: tflClient,
	}

	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:  "status",
				Usage: "Show status of all lines",
				Action: func(ctx context.Context, cmd *cli.Command) error {

					status, err := service.GetStatus(ctx)
					if err != nil {
						return err
					}

					service.RenderStatus(ctx, status)

					return nil
				},
			},
			{
				Name:  "station",
				Usage: "Show departures from a given station",
				Action: func(ctx context.Context, cmd *cli.Command) error {

					if cmd.Args().First() == "" {
						fmt.Printf("") // TODO: print something here when there's no arg

						// TODO: create NoStationError or MissingArgError
						return nil
					}

					arrivals, err := service.GetStationArrivals(ctx)
					if err != nil {
						return err
					}

					err = service.RenderArrivals(ctx, arrivals, cmd.Args().First(), DefaultWidth)
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
