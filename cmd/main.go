package main

import (
	"context"
	"fmt"
	"jamesalexatkin/tfl-cli/internal/config"
	"jamesalexatkin/tfl-cli/internal/presenter"
	"jamesalexatkin/tfl-cli/internal/service"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/jamesalexatkin/tfl-golang"
	"github.com/mattn/go-isatty"
	"github.com/urfave/cli/v3"
)

func init() {
	// Turn off color if not running in a proper terminal
	color.NoColor = !isatty.IsTerminal(os.Stdout.Fd())
}

const DefaultWidth = 70

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	tflClient := tfl.New(cfg.AppID, cfg.AppKey)

	service := service.Service{
		TFLClient: tflClient,
	}

	presenter := presenter.Presenter{}

	var statusVerbose bool

	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:  "show-config",
				Usage: "Show current config",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					presenter.RenderConfig(ctx, cfg)

					return nil
				},
			},
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
					err := cfg.Validate()
					if err != nil {
						return err
					}

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
					err := cfg.Validate()
					if err != nil {
						return err
					}

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
