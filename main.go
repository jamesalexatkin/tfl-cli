package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/jamesalexatkin/tfl-cli/internal/config"
	"github.com/jamesalexatkin/tfl-cli/internal/presenter"
	"github.com/jamesalexatkin/tfl-cli/internal/service"
	tfl "github.com/jamesalexatkin/tfl-golang"
	isatty "github.com/mattn/go-isatty"
	cli "github.com/urfave/cli/v3"
)

func init() {
	// Turn off color if not running in a proper terminal
	color.NoColor = !isatty.IsTerminal(os.Stdout.Fd())
}

//nolint:cyclop,funlen
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
		Name: "tfl",
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

					presenter.RenderStatus(ctx, status, statusVerbose)

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

					switch stationName {
					case "home":
						stationName = cfg.HomeStation
					case "work":
						stationName = cfg.WorkStation
					case "":
						fmt.Printf("") // TODO: print something here when there's no arg

						// TODO: create NoStationError or MissingArgError
						return nil
					}

					arrivals, err := service.FetchStationArrivalsBoard(ctx, stationName, cfg.NumDepartures)
					if err != nil {
						return err
					}

					err = presenter.RenderDepartureBoard(ctx, *arrivals, cfg.DepartureBoardWidth)
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
