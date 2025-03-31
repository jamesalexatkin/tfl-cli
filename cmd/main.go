package main

import (
	"context"
	"flag"
	"fmt"
	"jamesalexatkin/tfl-cli/internal"
	"log"
	"os"

	"github.com/jamesalexatkin/tfl-golang"
	"github.com/schachmat/ingo"
	"github.com/urfave/cli/v3"
)

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
				Action: func(context.Context, *cli.Command) error {
					fmt.Println("boom! I say!")
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
